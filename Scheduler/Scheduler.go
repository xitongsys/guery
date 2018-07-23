package Scheduler

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
	"github.com/xitongsys/guery/Config"
	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Optimizer"
	"github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/Row"
	"github.com/xitongsys/guery/Topology"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/pb"
	"google.golang.org/grpc"
)

type Scheduler struct {
	sync.Mutex
	Topology *Topology.Topology

	Todos, Doings, Dones, Fails TaskList

	TotalTaskNumber int64
}

func NewScheduler(topology *Topology.Topology) *Scheduler {
	res := &Scheduler{
		Topology: topology,
	}
	return res
}

func (self *Scheduler) AutoFresh() {
	go func() {
		for {
			time.Sleep(time.Millisecond * 5)
			self.RunTask()
		}
	}()
}

func (self *Scheduler) CancelTask(taskid int64) error {
	self.Lock()
	defer self.Unlock()

	var task *Task
	for i := 0; i < len(self.Todos) && task == nil; i++ {
		t := self.Todos[i]
		if t.TaskId == taskid {
			task = t
			break
		}
	}

	for i := 0; i < len(self.Doings) && task == nil; i++ {
		t := self.Doings[i]
		if t.TaskId == taskid {
			task = t
			break
		}
	}
	if task == nil {
		return fmt.Errorf("task not found")
	}
	self.FinishTask(task, pb.TaskStatus_ERROR, []*pb.LogInfo{pb.NewErrLogInfo("canceled by user")})
	return nil
}

func (self *Scheduler) AddTask(runtime *Config.ConfigRuntime, query string, output io.Writer) (*Task, error) {
	var err error
	self.Lock()
	defer self.Unlock()

	self.TotalTaskNumber++
	taskId := self.TotalTaskNumber
	task := &Task{
		TaskId:     taskId,
		Status:     pb.TaskStatus_TODO,
		Query:      query,
		Runtime:    runtime,
		CommitTime: time.Now(),
		Output:     output,
		DoneChan:   make(chan int),
	}

	self.Todos.Add(task)

	var logicalPlanTree Plan.PlanNode
	logicalPlanTree, err = Optimizer.CreateLogicalTree(runtime, query)
	if err == nil {
		task.LogicalPlanTree = logicalPlanTree
	}

	if err != nil {
		self.FinishTask(task, pb.TaskStatus_ERROR, []*pb.LogInfo{pb.NewErrLogInfo(fmt.Sprintf("%v", err))})
	}
	return task, err
}

func (self *Scheduler) RunTask() {
	self.Lock()
	defer self.Unlock()

	task := self.Todos.Top()
	if task == nil {
		return
	}

	freeExecutorsNumber := self.Topology.GetFreeExecutorNumber()

	l, r := int32(1), int32(task.Runtime.MaxConcurrentNumber)
	for l <= r {
		m := l + (r-l)/2
		men, _ := EPlan.GetEPlanExecutorNumber(task.LogicalPlanTree, m)
		if men > freeExecutorsNumber {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	pn := r
	if pn <= 0 {
		Logger.Infof("no enough executors")
		return
	}

	executorNumber, _ := EPlan.GetEPlanExecutorNumber(task.LogicalPlanTree, pn)
	self.Todos.Pop()

	task.SetStatus(pb.TaskStatus_RUNNING)
	self.Doings.Add(task)

	//start send to agents
	ePlanNodes := []EPlan.ENode{}
	freeAgents, freeExecutors := self.Topology.GetFreeExecutors(executorNumber)
	task.Agents = freeAgents

	var aggNode EPlan.ENode
	var err error

	if aggNode, err = EPlan.CreateEPlan(task.LogicalPlanTree, &ePlanNodes, &freeExecutors, int(pn)); err == nil {
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
				Logger.Errorf("failed to dial: %v", err)
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
				Logger.Errorf("failed to dial: %v", err)
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
		Logger.Errorf("task failed: %v", err)
		self.FinishTask(task, pb.TaskStatus_ERROR, []*pb.LogInfo{pb.NewErrLogInfo(fmt.Sprintf("%v", err))})
		return
	}

	task.EPlanNodes = ePlanNodes

	Logger.Infof("begin to collect results")
	go self.CollectResults(task)

}

func (self *Scheduler) FinishTask(task *Task, status pb.TaskStatus, errs []*pb.LogInfo) {
	self.KillTask(task)
	switch task.Status {
	case pb.TaskStatus_SUCCEED, pb.TaskStatus_ERROR:
		return
	case pb.TaskStatus_RUNNING:
		if self.Doings.Delete(task) != nil {
			return
		}
	case pb.TaskStatus_TODO:
		if self.Todos.Delete(task) != nil {
			return
		}
	default:
		return
	}

	task.Status = status
	switch task.Status {
	case pb.TaskStatus_SUCCEED:
		self.Dones = append(self.Dones, task)
	default:
		task.Status = pb.TaskStatus_ERROR
		self.Fails = append(self.Fails, task)
	}

	task.EndTime = time.Now()

	for _, err := range errs {
		if err != nil {
			task.Infos = append(task.Infos, err)
		}
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
		if taskInfo.Status == pb.TaskStatus_ERROR {
			task := &Task{
				TaskId: taskInfo.TaskId,
			}
			self.FinishTask(task, pb.TaskStatus_ERROR, taskInfo.Infos)
		}

		task := self.Doings.GetTask(taskInfo.TaskId)
		if task != nil {
			task.Progress = taskInfo.Progress
		}
	}
}

func (self *Scheduler) CollectResults(task *Task) {
	var errs []*pb.LogInfo
	defer func() {
		self.Lock()
		if len(errs) > 0 {
			self.FinishTask(task, pb.TaskStatus_ERROR, errs)
		} else {
			self.FinishTask(task, pb.TaskStatus_SUCCEED, errs)
		}
		self.Unlock()
	}()

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
		Logger.Errorf("failed to connect to input channel %v: %v", inputChannelLocation, err)
		errs = append(errs, pb.NewErrLogInfo(fmt.Sprintf("%v", err)))
		return
	}

	//send results
	var (
		msg []byte
		n   int
		row *Row.Row
	)

	md := &Metadata.Metadata{}
	if err = Util.ReadObject(cconn, md); err != nil {
		errs = append(errs, pb.NewErrLogInfo(fmt.Sprintf("%v", err)))
		return
	}

	if msg, err = json.MarshalIndent(md, "", "    "); err != nil {
		Logger.Errorf("json marshal: %v", err)
		errs = append(errs, pb.NewErrLogInfo(fmt.Sprintf("%v", err)))
		return
	}
	msg = append(msg, []byte("\n")...)

	if n, err = response.Write(msg); n != len(msg) || err != nil {
		errs = append(errs, pb.NewErrLogInfo(fmt.Sprintf("%v", err)))
		return
	}

	rbReader := Row.NewRowsBuffer(md, cconn, nil)

	for {
		row, err = rbReader.ReadRow()

		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			errs = append(errs, pb.NewErrLogInfo(fmt.Sprintf("%v", err)))
			return
		}

		res := []string{}
		for i := 0; i < len(row.Vals); i++ {
			res = append(res, fmt.Sprintf("%v", row.Vals[i]))
		}
		msg = []byte(strings.Join(res, ","))
		msg = append(msg, []byte("\n")...)

		if n, err = response.Write(msg); n != len(msg) || err != nil {
			errs = append(errs, pb.NewErrLogInfo(fmt.Sprintf("%v", err)))
			return
		}
	}
}
