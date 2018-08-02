package scheduler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/vmihailenco/msgpack"
	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/eplan"
	"github.com/xitongsys/guery/logger"
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/pb"
	"github.com/xitongsys/guery/row"
	"github.com/xitongsys/guery/topology"
	"github.com/xitongsys/guery/util"
	"google.golang.org/grpc"
)

type Scheduler struct {
	sync.Mutex
	Topology *topology.Topology

	TodoQueue, RunningQueue  *Queue
	SucceedQueue, ErrorQueue *Queue
}

func NewScheduler(topology *topology.Topology) *Scheduler {
	res := &Scheduler{
		Topology:     topology,
		TodoQueue:    NewQueue("TODO"),
		RunningQueue: NewQueue("RUNNING"),
		SucceedQueue: NewQueue("SUCCEED"),
		ErrorQueue:   NewQueue("ERROR"),
	}
	return res
}

func (self *Scheduler) AutoFresh() {
	go func() {
		for {
			time.Sleep(time.Millisecond * 5)
			if self.RunningQueue.Tasks.Len() < int(config.Conf.Runtime.MaxConcurrentTaskNumber) {
				self.RunTask()
			}
		}
	}()
}

func (self *Scheduler) CancelTask(taskid string) error {
	self.Lock()
	defer self.Unlock()

	task := self.TodoQueue.GetTask(taskid)
	if task == nil {
		task = self.RunningQueue.GetTask(taskid)
	}

	if task == nil {
		return fmt.Errorf("task not found")
	}

	self.FinishTask(task, pb.TaskStatus_ERROR, []*pb.LogInfo{pb.NewErrLogInfo("canceled by user")})
	return nil
}

func (self *Scheduler) AddTask(task *Task) (err error) {
	self.Lock()
	defer self.Unlock()

	if task.Status == pb.TaskStatus_TODO {
		return self.TodoQueue.Add(task)
	} else if task.Status == pb.TaskStatus_ERROR {
		return self.ErrorQueue.AddForce(task)
	} else if task.Status == pb.TaskStatus_SUCCEED {
		return self.SucceedQueue.AddForce(task)
	} else if task.Status == pb.TaskStatus_RUNNING {
		return self.RunningQueue.Add(task)
	} else {
		return fmt.Errorf("unknown task status")
	}
	return nil
}

func (self *Scheduler) RunTask() {
	self.Lock()
	defer self.Unlock()

	task := self.TodoQueue.Pop()
	if task == nil {
		return
	}
	task.SetStatus(pb.TaskStatus_RUNNING)
	self.RunningQueue.Add(task)

	//start send to agents
	ePlanNodes := []eplan.ENode{}
	executorNumber, _ := eplan.GetEPlanExecutorNumber(task.LogicalPlanTree, task.Runtime.ParallelNumber)
	freeAgents, freeExecutors := self.Topology.GetExecutors(int(executorNumber + 1))
	task.Agents = freeAgents

	var aggNode eplan.ENode
	var err error

	if aggNode, err = eplan.CreateEPlan(task.LogicalPlanTree, &ePlanNodes, &freeExecutors, int(task.Runtime.ParallelNumber)); err == nil {
		task.AggNode = aggNode

		var grpcConn *grpc.ClientConn
		var (
			buf        []byte
			runtimeBuf []byte
		)

		agentTasks := map[string]*pb.Task{}
		for _, enode := range ePlanNodes {
			if buf, err = msgpack.Marshal(enode); err != nil {
				break
			}

			if runtimeBuf, err = msgpack.Marshal(task.Runtime); err != nil {
				break
			}
			loc := enode.GetLocation()

			instruction := pb.Instruction{
				TaskId:                task.TaskId,
				TaskType:              int32(enode.GetNodeType()),
				EncodedEPlanNodeBytes: buf,
				RuntimeBytes:          runtimeBuf,
				Location:              &loc,
			}
			agentURL := loc.GetURL()
			if _, ok := agentTasks[agentURL]; !ok {
				agentTasks[agentURL] = &pb.Task{
					TaskId:       task.TaskId,
					Instructions: []*pb.Instruction{},
				}
			}
			agentTasks[agentURL].Instructions = append(agentTasks[agentURL].Instructions, &instruction)
		}

		//send task & setup writer
		for agentURL, agentTask := range agentTasks {
			grpcConn, err = grpc.Dial(agentURL, grpc.WithInsecure())
			if err != nil {
				logger.Errorf("failed to dial: %v", err)
				break
			}
			client := pb.NewGueryAgentClient(grpcConn)
			if _, err = client.SendTask(context.Background(), agentTask); err != nil {
				grpcConn.Close()
				break
			}
			grpcConn.Close()
		}

		//send task & setup reader
		for agentURL, agentTask := range agentTasks {
			grpcConn, err = grpc.Dial(agentURL, grpc.WithInsecure())
			if err != nil {
				logger.Errorf("failed to dial: %v", err)
				break
			}
			client := pb.NewGueryAgentClient(grpcConn)
			if _, err = client.Run(context.Background(), agentTask); err != nil {
				grpcConn.Close()
				break
			}
			grpcConn.Close()
		}
	}

	if err != nil {
		logger.Errorf("task failed: %v", err)
		self.FinishTask(task, pb.TaskStatus_ERROR, []*pb.LogInfo{pb.NewErrLogInfo(fmt.Sprintf("%v", err))})
		return
	}

	task.EPlanNodes = ePlanNodes

	logger.Infof("begin to collect results")
	go self.CollectResults(task)

}

func (self *Scheduler) FinishTask(task *Task, status pb.TaskStatus, errs []*pb.LogInfo) {
	self.KillTask(task)
	switch task.Status {
	case pb.TaskStatus_SUCCEED, pb.TaskStatus_ERROR:
		return
	case pb.TaskStatus_RUNNING:
		if self.RunningQueue.Delete(task) != nil {
			return
		}
	case pb.TaskStatus_TODO:
		if self.TodoQueue.Delete(task) != nil {
			return
		}
	default:
		return
	}

	task.EndTime = time.Now()
	for _, err := range errs {
		if err != nil {
			task.Infos = append(task.Infos, err)
		}
	}
	task.Status = status

	switch task.Status {
	case pb.TaskStatus_SUCCEED:
		self.SucceedQueue.Add(task)
	default:
		task.Status = pb.TaskStatus_ERROR
		self.ErrorQueue.Add(task)
	}
	close(task.DoneChan)
}

func (self *Scheduler) KillTask(task *Task) error {
	for _, loc := range task.Agents {
		grpcConn, err := grpc.Dial(loc.GetURL(), grpc.WithInsecure())
		if err != nil {
			return err
		}
		client := pb.NewGueryAgentClient(grpcConn)
		if _, err = client.KillTask(context.Background(), &pb.Task{TaskId: task.TaskId}); err != nil {
			grpcConn.Close()
			return err
		}
		grpcConn.Close()
	}
	return nil
}

func (self *Scheduler) UpdateTasks(agentHeartbeat *pb.AgentHeartbeat) {
	self.Lock()
	defer self.Unlock()
	//kill error task
	for _, taskInfo := range agentHeartbeat.TaskInfos {
		task := self.RunningQueue.GetTask(taskInfo.TaskId)
		if task != nil {
			if taskInfo.Status == pb.TaskStatus_ERROR {
				self.FinishTask(task, taskInfo.Status, taskInfo.Infos)
			}
			task.Progress = taskInfo.Progress
			task.AgentStatus[agentHeartbeat.Location.Name] = taskInfo.Status
		}
	}
}

func (self *Scheduler) CollectResults(task *Task) {
	defer func() {
		task.DoneChan <- 0
		self.FinishTask(task, pb.TaskStatus_SUCCEED, []*pb.LogInfo{})
	}()
	var errs []*pb.LogInfo
	enode := task.AggNode
	response := task.Output

	output := enode.GetLocation()
	conn, err := grpc.Dial(output.GetURL(), grpc.WithInsecure())
	if err != nil {
		return
	}
	client := pb.NewGueryAgentClient(conn)
	inputChannelLocation, err := client.GetOutputChannelLocation(context.Background(), &output)
	if err != nil {
		errs = append(errs, pb.NewErrLogInfo(fmt.Sprintf("%v", err)))
		return
	}
	conn.Close()

	cconn, err := net.Dial("tcp", inputChannelLocation.GetURL())
	if err != nil {
		logger.Errorf("failed to connect to input channel %v: %v", inputChannelLocation, err)
		errs = append(errs, pb.NewErrLogInfo(fmt.Sprintf("%v", err)))
		return
	}

	//send results
	var (
		msg []byte
		n   int
		r   *row.Row
	)

	md := &metadata.Metadata{}
	if err = util.ReadObject(cconn, md); err != nil {
		errs = append(errs, pb.NewErrLogInfo(fmt.Sprintf("%v", err)))
		return
	}

	if msg, err = json.MarshalIndent(md, "", "    "); err != nil {
		logger.Errorf("json marshal: %v", err)
		errs = append(errs, pb.NewErrLogInfo(fmt.Sprintf("%v", err)))
		return
	}
	msg = append(msg, []byte("\n")...)

	if n, err = response.Write(msg); n != len(msg) || err != nil {
		errs = append(errs, pb.NewErrLogInfo(fmt.Sprintf("%v", err)))
		return
	}

	rbReader := row.NewRowsBuffer(md, cconn, nil)

	for {
		r, err = rbReader.ReadRow()

		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			errs = append(errs, pb.NewErrLogInfo(fmt.Sprintf("%v", err)))
			return
		}

		res := []string{}
		for i := 0; i < len(r.Vals); i++ {
			res = append(res, fmt.Sprintf("%v", r.Vals[i]))
		}
		msg = []byte(strings.Join(res, ","))
		msg = append(msg, []byte("\n")...)

		if n, err = response.Write(msg); n != len(msg) || err != nil {
			errs = append(errs, pb.NewErrLogInfo(fmt.Sprintf("%v", err)))
			return
		}
	}
}
