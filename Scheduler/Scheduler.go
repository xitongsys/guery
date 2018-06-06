package Scheduler

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net"
	"sync"
	"time"

	"github.com/vmihailenco/msgpack"
	"github.com/xitongsys/guery/Config"
	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/Optimizer"
	"github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/Topology"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/pb"
	"google.golang.org/grpc"
)

type Scheduler struct {
	sync.Mutex
	Topology *Topology.Topology

	Todos, Doings, Dones, Fails TaskList
	AllocatedMap                map[string]int64 //executorName:taskId

	TotalTaskNumber int64
}

func NewScheduler(topology *Topology.Topology) *Scheduler {
	res := &Scheduler{
		Topology:     topology,
		AllocatedMap: map[string]int64{},
	}
	return res
}

func (self *Scheduler) AutoFresh() {
	go func() {
		for {
			self.RunTask()
		}
	}()
}

func (self *Scheduler) AddTask(query, catalog, schema string, priority int32, output io.Writer) (*Task, error) {
	var err error
	self.Lock()
	defer self.Unlock()

	self.TotalTaskNumber++
	taskId := self.TotalTaskNumber
	task := &Task{
		TaskId: taskId,
		Status: TODO,

		Executors: []string{},
		Query:     query,
		Catalog:   catalog,
		Schema:    schema,
		Priority:  priority,

		CommitTime: time.Now(),

		Output: output,

		DoneChan: make(chan int),
	}

	self.Todos.Add(task)

	var logicalPlanTree Plan.PlanNode
	logicalPlanTree, err = Optimizer.CreateLogicalTree(query)
	if err == nil {
		task.LogicalPlanTree = logicalPlanTree
		task.ExecutorNumber, err = EPlan.GetEPlanExecutorNumber(task.LogicalPlanTree, 1)
	}

	if err != nil {
		self.FinishTask(task, FAILED, []error{err})
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

	allFreeExecutors := []pb.Location{}
	for _, loc := range self.Topology.GetExecutors() {
		name := loc.Name
		if _, ok := self.AllocatedMap[name]; !ok {
			allFreeExecutors = append(allFreeExecutors, loc)
		}
	}

	freeExecutorsNumber := int32(len(allFreeExecutors))

	l, r := int32(1), int32(Config.Conf.Runtime.MaxConcurrentNumber)
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

	task.ExecutorNumber, _ = EPlan.GetEPlanExecutorNumber(task.LogicalPlanTree, pn)
	self.Todos.Pop()

	task.SetStatus(DOING)
	self.Doings.Add(task)

	//start send to executor
	ePlanNodes := []EPlan.ENode{}
	freeExecutors := allFreeExecutors[:task.ExecutorNumber]

	var aggNode EPlan.ENode
	var err error

	///debug info
	Logger.Infof("================")
	for _, loc := range freeExecutors {
		Logger.Infof("%v", loc.Name, loc.GetURL())
	}
	Logger.Infof("================")

	log.Println("========", pn, task.ExecutorNumber)

	if aggNode, err = EPlan.CreateEPlan(task.LogicalPlanTree, &ePlanNodes, &freeExecutors, int(pn)); err == nil {
		task.AggNode = aggNode

		var grpcConn *grpc.ClientConn
		var buf []byte
		for _, enode := range ePlanNodes {
			buf, err = msgpack.Marshal(enode)

			instruction := pb.Instruction{
				TaskId:                task.TaskId,
				TaskType:              int32(enode.GetNodeType()),
				EncodedEPlanNodeBytes: buf,
			}

			loc := enode.GetLocation()

			Logger.Infof("dial %v", loc.GetURL())
			grpcConn, err = grpc.Dial(loc.GetURL(), grpc.WithInsecure())
			if err != nil {
				Logger.Errorf("failed to dial: %v", err)
				break
			}
			client := pb.NewGueryExecutorClient(grpcConn)

			Logger.Infof("send instruction to %v", loc.GetURL())
			if _, err = client.SendInstruction(context.Background(), &instruction); err != nil {
				grpcConn.Close()
				Logger.Errorf("failed to send instruction: %v", err)
				break
			}

			Logger.Infof("setup writers of %v", loc.GetURL())
			empty := pb.Empty{}
			if _, err = client.SetupWriters(context.Background(), &empty); err != nil {
				grpcConn.Close()
				break
			}
			grpcConn.Close()
		}

		Logger.Infof("finished to setup writers")

		if err == nil {
			for _, enode := range ePlanNodes {
				loc := enode.GetLocation()
				var grpcConn *grpc.ClientConn
				grpcConn, err = grpc.Dial(loc.GetURL(), grpc.WithInsecure())
				if err != nil {
					break
				}
				client := pb.NewGueryExecutorClient(grpcConn)
				empty := pb.Empty{}

				Logger.Infof("setup readers of %v", loc.GetURL())
				if _, err = client.SetupReaders(context.Background(), &empty); err != nil {
					Logger.Errorf("failed setup readers %v: %v", loc, err)
					grpcConn.Close()
					break
				}

				if _, err = client.Run(context.Background(), &empty); err != nil {
					Logger.Errorf("failed run %v: %v", loc, err)
					grpcConn.Close()
					break
				}
				grpcConn.Close()
			}
			Logger.Infof("finished to setup readers & run")
		}

	}

	if err != nil {
		self.FinishTask(task, FAILED, []error{err})
		return
	}

	task.Executors = []string{}
	for _, executor := range allFreeExecutors[:task.ExecutorNumber] {
		self.AllocatedMap[executor.Name] = task.TaskId
		task.Executors = append(task.Executors, executor.Name)
	}

	Logger.Infof("begin to collect results")
	go self.CollectResults(task)

}

func (self *Scheduler) FinishTask(task *Task, status TaskStatusType, errs []error) {
	switch task.Status {
	case DONE, FAILED:
		return
	case DOING:
		if self.Doings.Delete(task) != nil {
			return
		}
	case TODO:
		if self.Todos.Delete(task) != nil {
			return
		}
	default:
		return
	}

	task.Status = status
	switch task.Status {
	case DONE:
		self.Dones = append(self.Dones, task)
	case DOING, TODO, FAILED:
		task.Status = FAILED
		fallthrough
	default:
		self.Fails = append(self.Fails, task)
	}

	task.EndTime = time.Now()

	for _, name := range task.Executors {
		if status == FAILED {
			self.Topology.RestartExecutor(name)
		}
		delete(self.AllocatedMap, name)
	}
	for _, err := range errs {
		if err != nil {
			task.Errs = append(task.Errs, err)
		}
	}
	close(task.DoneChan)
}

func (self *Scheduler) CollectResults(task *Task) {
	var errs []error
	defer func() {
		self.Lock()
		if len(errs) > 0 {
			self.FinishTask(task, FAILED, errs)
		} else {
			self.FinishTask(task, DONE, errs)
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
	client := pb.NewGueryExecutorClient(conn)
	inputChannelLocation, err := client.GetOutputChannelLocation(context.Background(), &output)
	if err != nil {
		errs = append(errs, err)
		return
	}
	conn.Close()

	cconn, err := net.Dial("tcp", inputChannelLocation.GetURL())
	if err != nil {
		Logger.Errorf("failed to connect to input channel %v: %v", inputChannelLocation, err)
		errs = append(errs, err)
		return
	}

	//send results
	var (
		msg []byte
		n   int
		row *Util.Row
	)

	md := &Util.Metadata{}
	if err = Util.ReadObject(cconn, md); err != nil {
		errs = append(errs, err)
		return
	}

	if msg, err = json.MarshalIndent(md, "", "    "); err != nil {
		Logger.Errorf("json marshal: %v", err)
		errs = append(errs, err)
		return
	}

	if n, err = response.Write(msg); n != len(msg) || err != nil {
		errs = append(errs, err)
		return
	}

	rbReader := Util.NewRowsBuffer(md, cconn, nil)

	for {
		row, err = rbReader.ReadRow()
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			errs = append(errs, err)
			return
		}

		if msg, err = json.MarshalIndent(row, "", "    "); err != nil {
			errs = append(errs, err)
			return
		}

		if n, err = response.Write(msg); n != len(msg) || err != nil {
			errs = append(errs, err)
			return
		}
	}
}
