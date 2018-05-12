package Scheduler

import (
	"context"
	"encoding/json"
	"io"
	"net"
	"sort"
	"sync"
	"time"

	"github.com/vmihailenco/msgpack"
	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/Topology"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/pb"
	"google.golang.org/grpc"
)

const (
	MAXPN int32 = 1000
	MINPN int32 = 1
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

	var logicalPlanTree Plan.PlanNode
	logicalPlanTree, err = Plan.CreateLogicalTree(query)
	if err == nil {
		task.LogicalPlanTree = logicalPlanTree
		task.ExecutorNumber, err = EPlan.GetEPlanExecutorNumber(task.LogicalPlanTree, 1)
		if err == nil {
			self.Todos = append(self.Todos, task)
			sort.Sort(self.Todos)
		}
	}

	if err != nil {
		self.Fails = append(self.Fails, task)
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

	l, r := MINPN, MAXPN
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

	if aggNode, err = EPlan.CreateEPlan(task.LogicalPlanTree, &ePlanNodes, &freeExecutors, 1); err == nil {
		task.AggNode = aggNode

		var grpcConn *grpc.ClientConn

		for _, enode := range ePlanNodes {
			buf, err := msgpack.Marshal(enode)

			instruction := pb.Instruction{
				TaskId:                task.TaskId,
				TaskType:              int32(enode.GetNodeType()),
				Catalog:               task.Catalog,
				Schema:                task.Schema,
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
		task.Status = FAILED
		self.Fails = append(self.Fails, task)
		return
	}

	self.Doings = append(self.Doings, task)
	task.BeginTime = time.Now()

	task.Executors = []string{}
	for _, executor := range allFreeExecutors[:task.ExecutorNumber] {
		self.AllocatedMap[executor.Name] = task.TaskId
		task.Executors = append(task.Executors, executor.Name)
	}

	Logger.Infof("begin to collect results")
	go self.CollectResults(task)

}

func (self *Scheduler) FinishTask(task *Task) {
	self.Lock()
	defer self.Unlock()

	i, ln := 0, len(self.Doings)
	for i = 0; i < ln; i++ {
		if self.Doings[i].TaskId == task.TaskId {
			break
		}
	}

	if task.Status == DONE {
		self.Dones = append(self.Dones, task)
	} else {
		task.Status = FAILED
		self.Fails = append(self.Fails, task)
	}
	task.EndTime = time.Now()

	if i < ln {
		for j := i; j < ln-1; j++ {
			self.Doings[j] = self.Doings[j+1]
		}
		self.Doings = self.Doings[:ln-1]
	}

	for _, name := range task.Executors {
		delete(self.AllocatedMap, name)
	}

	close(task.DoneChan)
}

func (self *Scheduler) CollectResults(task *Task) {
	defer self.FinishTask(task)

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
		return
	}
	conn.Close()

	cconn, err := net.Dial("tcp", inputChannelLocation.GetURL())
	if err != nil {
		Logger.Errorf("failed to connect to input channel %v: %v", inputChannelLocation, err)
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
		return
	}

	if msg, err = json.Marshal(md); err != nil {
		Logger.Errorf("json marshal: %v", err)
		return
	}

	if n, err = response.Write(msg); n != len(msg) || err != nil {
		return
	}

	for {
		row, err = Util.ReadRow(cconn)
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			break
		}

		if msg, err = json.Marshal(row); err != nil {
			break
		}

		if n, err = response.Write(msg); n != len(msg) || err != nil {
			break
		}
	}

	if err != nil {
		task.Status = FAILED
		task.Err = err
	} else {
		task.Status = DONE
	}
}
