package Agent

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/kardianos/osext"
	"github.com/xitongsys/guery/Config"
	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/pb"
	"google.golang.org/grpc"
)

const (
	TIMEOUT = 10000 //ms
)

type Agent struct {
	MasterAddress string
	Address       string
	Name          string
	Topology      *Topology
	StartTime     time.Time
	Tasks         *TaskMap
}

var agentServer *Executor

func NewAgent(masterAddress string, address, name string) *Agent {
	res := &Agent{
		MasterAddress: masterAddress,
		Address:       address,
		Name:          name,
		Topology:      NewTopology(),
		Tasks:         NewTaskMap(),
	}
	return res
}

func (self *Agent) SendTask(ctx context.Context, task *pb.Task) (*pb.Empty, error) {
	res := &pb.Empty{}
	var err error
	if task == nil {
		return res, nil
	}
	if task.GetInstruction() == nil {
		return res, fmt.Errorf("[Agent] instruction is nil")
	}

	if err = self.Tasks.AddTask(task); err != nil {
		return err
	}
	for _, inst := range task.GetInstruction() {
		if err = self.LanchExecutor(inst.Location.Name); err != nil {
			return res, err
		}
	}

	flag := false
	timeout := time.After(TIMEOUT * time.Millisecond)
	tick := time.Tick(50 * time.Millisecond)
	for !flag {
		flag = true
		select {
		case <-timeout:
			err = fmt.Errorf("timeout")
			break
		case <-tick:
			for _, inst := range task.GetInstruction() {
				name := inst.Location.Name
				if !self.Topology.HasExecutor(name) {
					falg = false
					break
				}
			}
		}
	}
	if err != nil {
		self.KillTask(context.Background(), task)
	}
	return res, err
}

func (self *Agent) KillTask(ctx context.Context, task *pb.Task) (*pb.Empty, error) {
	res := &pb.Empty{}
	var err error
	if task == nil {
		return nil
	}
	task := self.Tasks.PopTask(task.TaskId)
	if task == nil {
		return nil
	}
	for _, inst := range task.Instruction {
		ename := inst.Location.Name
		self.KillExecutor(ename) //err handle?
	}
	return res, nil

}

func (self *Agent) KillExecutor(ename string) error {
	return self.Topology.KillExecutor(ename)
}

func (self *Agent) LanchExecutor(ename string) error {
	exeFullName, _ := osext.Executable()
	command := exec.Command(exeFullName,
		fmt.Sprintf("executor"),
		"--agent",
		fmt.Sprintf("%v", self.Address),
		"--address",
		fmt.Sprintf("%v", strings.Split(self.Address, ":")[0]+":0"),
		"--name",
		ename,
		"--config",
		fmt.Sprintf("%v", Config.Conf.File),
	)
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr
	return command.Start()
}

func (self *Agent) Duplicate(ctx context.Context, em *pb.Empty) (*pb.Empty, error) {
	res := &pb.Empty{}
	exeFullName, _ := osext.Executable()

	command := exec.Command(exeFullName,
		fmt.Sprintf("agent"),
		"--master",
		fmt.Sprintf("%v", self.MasterAddress),
		"--address",
		fmt.Sprintf("%v", strings.Split(self.Address, ":")[0]+":0"),
		"--config",
		fmt.Sprintf("%v", Config.Conf.File),
	)

	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr
	err := command.Start()
	return res, err
}

func (self *Agent) Quit(ctx context.Context, em *pb.Empty) (*pb.Empty, error) {
	res := &pb.Empty{}
	os.Exit(0)
	return res, nil
}

func (self *Agent) Restart(ctx context.Context, em *pb.Empty) (*pb.Empty, error) {
	res := &pb.Empty{}
	self.Duplicate(context.Background(), em)
	time.Sleep(time.Second)
	self.Quit(ctx, em)
	return res, nil
}

func (self *Agent) SendInstruction(ctx context.Context, instruction *pb.Instruction) (*pb.Empty, error) {
	res := &pb.Empty{}

	if err := self.SetRuntime(instruction); err != nil {
		return res, err
	}

	nodeType := EPlan.EPlanNodeType(instruction.TaskType)
	Logger.Infof("Instruction: %v", instruction.TaskType)
	self.Status = 1
	self.IsStatusChanged = true

	self.DoneChan = make(chan int)
	switch nodeType {
	case EPlan.ESCANNODE:
		return res, self.SetInstructionScan(instruction)
	case EPlan.ESELECTNODE:
		return res, self.SetInstructionSelect(instruction)
	case EPlan.EGROUPBYNODE:
		return res, self.SetInstructionGroupBy(instruction)
	case EPlan.EJOINNODE:
		return res, self.SetInstructionJoin(instruction)
	case EPlan.EHASHJOINNODE:
		return res, self.SetInstructionHashJoin(instruction)
	case EPlan.EHASHJOINSHUFFLENODE:
		return res, self.SetInstructionHashJoinShuffle(instruction)
	case EPlan.EDUPLICATENODE:
		return res, self.SetInstructionDuplicate(instruction)
	case EPlan.EAGGREGATENODE:
		return res, self.SetInstructionAggregate(instruction)
	case EPlan.EAGGREGATEFUNCGLOBALNODE:
		return res, self.SetInstructionAggregateFuncGlobal(instruction)
	case EPlan.EAGGREGATEFUNCLOCALNODE:
		return res, self.SetInstructionAggregateFuncLocal(instruction)
	case EPlan.ELIMITNODE:
		return res, self.SetInstructionLimit(instruction)
	case EPlan.EFILTERNODE:
		return res, self.SetInstructionFilter(instruction)
	case EPlan.EUNIONNODE:
		return res, self.SetInstructionUnion(instruction)
	case EPlan.EORDERBYLOCALNODE:
		return res, self.SetInstructionOrderByLocal(instruction)
	case EPlan.EORDERBYNODE:
		return res, self.SetInstructionOrderBy(instruction)
	case EPlan.ESHOWNODE:
		return res, self.SetInstructionShow(instruction)
	case EPlan.EBALANCENODE:
		return res, self.SetInstructionBalance(instruction)
	default:
		self.Status = 0
		return res, fmt.Errorf("Unknown node type")
	}
	return res, nil
}

func (self *Agent) Run(ctx context.Context, empty *pb.Empty) (*pb.Empty, error) {
	res := &pb.Empty{}
	nodeType := EPlan.EPlanNodeType(self.Instruction.TaskType)

	switch nodeType {
	case EPlan.ESCANNODE:
		go self.RunScan()
	case EPlan.ESELECTNODE:
		go self.RunSelect()
	case EPlan.EGROUPBYNODE:
		go self.RunGroupBy()
	case EPlan.EJOINNODE:
		go self.RunJoin()
	case EPlan.EHASHJOINNODE:
		go self.RunHashJoin()
	case EPlan.EHASHJOINSHUFFLENODE:
		go self.RunHashJoinShuffle()
	case EPlan.EDUPLICATENODE:
		go self.RunDuplicate()
	case EPlan.EAGGREGATENODE:
		go self.RunAggregate()
	case EPlan.EAGGREGATEFUNCGLOBALNODE:
		go self.RunAggregateFuncGlobal()
	case EPlan.EAGGREGATEFUNCLOCALNODE:
		go self.RunAggregateFuncLocal()
	case EPlan.ELIMITNODE:
		go self.RunLimit()
	case EPlan.EFILTERNODE:
		go self.RunFilter()
	case EPlan.EORDERBYLOCALNODE:
		go self.RunOrderByLocal()
	case EPlan.EORDERBYNODE:
		go self.RunOrderBy()
	case EPlan.EUNIONNODE:
		go self.RunUnion()
	case EPlan.ESHOWNODE:
		go self.RunShow()
	case EPlan.EBALANCENODE:
		go self.RunBalance()
	default:
		return res, fmt.Errorf("Unknown node type")
	}
	return res, nil
}

func (self *Agent) GetOutputChannelLocation(ctx context.Context, location *pb.Location) (*pb.Location, error) {
	if int(location.ChannelIndex) >= len(self.OutputChannelLocations) {
		return nil, fmt.Errorf("ChannelLocation %v not found: %v", location.ChannelIndex, location)
	}
	return self.OutputChannelLocations[location.ChannelIndex], nil
}

///////////////////////////////
func RunAgent(masterAddress string, address, name string) {
	executorServer = NewAgent(masterAddress, address, name)
	listener, err := net.Listen("tcp", executorServer.Address)
	if err != nil {
		log.Fatalf("Agent failed to run: %v", err)
	}
	defer listener.Close()
	executorServer.Address = listener.Addr().String()
	Logger.Infof("Agent: %v", executorServer.Address)

	go executorServer.Heartbeat()

	grpcS := grpc.NewServer()
	pb.RegisterGueryAgentServer(grpcS, executorServer)
	grpcS.Serve(listener)
}
