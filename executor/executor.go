package executor

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
	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/eplan"
	"github.com/xitongsys/guery/logger"
	"github.com/xitongsys/guery/pb"
	"google.golang.org/grpc"
)

type Executor struct {
	sync.Mutex
	AgentAddress string

	Address string
	Name    string

	Instruction                                   *pb.Instruction
	EPlanNode                                     eplan.ENode
	InputLocations, OutputLocations               []*pb.Location
	InputChannelLocations, OutputChannelLocations []*pb.Location
	Readers                                       []io.Reader
	Writers                                       []io.Writer

	Status          pb.TaskStatus
	IsStatusChanged bool
	Infos           []*pb.LogInfo

	DoneChan chan int
}

var executorServer *Executor

func NewExecutor(agentAddress string, address, name string) *Executor {
	res := &Executor{
		AgentAddress: agentAddress,
		Address:      address,
		Name:         name,
		DoneChan:     make(chan int),
		Infos:        []*pb.LogInfo{},
		Status:       pb.TaskStatus_TODO,
	}
	return res
}

func (self *Executor) AddLogInfo(info interface{}, level pb.LogLevel) {
	if info == nil {
		return
	}
	logInfo := &pb.LogInfo{
		Level: level,
		Info:  []byte(fmt.Sprintf("%v", info)),
	}
	self.Lock()
	defer self.Unlock()
	self.Infos = append(self.Infos, logInfo)
	if level == pb.LogLevel_ERR {
		self.Status = pb.TaskStatus_ERROR
		self.IsStatusChanged = true
	}
}

func (self *Executor) Clear() {
	for _, writer := range self.Writers {
		writer.(io.WriteCloser).Close()
	}
	self.IsStatusChanged = true
	if self.Status != pb.TaskStatus_ERROR {
		self.Status = pb.TaskStatus_SUCCEED
	}

	select {
	case <-self.DoneChan:
	default:
		close(self.DoneChan)
	}
}

func (self *Executor) Duplicate(ctx context.Context, em *pb.Empty) (*pb.Empty, error) {
	res := &pb.Empty{}
	exeFullName, _ := osext.Executable()

	command := exec.Command(exeFullName,
		fmt.Sprintf("executor"),
		"--agent",
		fmt.Sprintf("%v", self.AgentAddress),
		"--address",
		fmt.Sprintf("%v", strings.Split(self.Address, ":")[0]+":0"),
		"--config",
		fmt.Sprintf("%v", config.Conf.File),
	)

	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr
	err := command.Start()
	return res, err
}

func (self *Executor) Quit(ctx context.Context, em *pb.Empty) (*pb.Empty, error) {
	res := &pb.Empty{}
	os.Exit(0)
	return res, nil
}

func (self *Executor) Restart(ctx context.Context, em *pb.Empty) (*pb.Empty, error) {
	res := &pb.Empty{}
	self.Duplicate(context.Background(), em)
	time.Sleep(time.Second)
	self.Quit(ctx, em)
	return res, nil
}

func (self *Executor) SendInstruction(ctx context.Context, instruction *pb.Instruction) (*pb.Empty, error) {
	res := &pb.Empty{}

	if err := self.SetRuntime(instruction); err != nil {
		return res, err
	}

	nodeType := eplan.EPlanNodeType(instruction.TaskType)
	logger.Infof("Instruction: %v", instruction.TaskType)
	self.Status = pb.TaskStatus_RUNNING
	self.IsStatusChanged = true

	self.DoneChan = make(chan int)
	switch nodeType {
	case eplan.ESCANNODE:
		return res, self.SetInstructionScan(instruction)
	case eplan.ESELECTNODE:
		return res, self.SetInstructionSelect(instruction)
	case eplan.EGROUPBYNODE:
		return res, self.SetInstructionGroupBy(instruction)
	case eplan.EJOINNODE:
		return res, self.SetInstructionJoin(instruction)
	case eplan.EHASHJOINNODE:
		return res, self.SetInstructionHashJoin(instruction)
	case eplan.ESHUFFLENODE:
		return res, self.SetInstructionShuffle(instruction)
	case eplan.EDUPLICATENODE:
		return res, self.SetInstructionDuplicate(instruction)
	case eplan.EAGGREGATENODE:
		return res, self.SetInstructionAggregate(instruction)
	case eplan.EAGGREGATEFUNCGLOBALNODE:
		return res, self.SetInstructionAggregateFuncGlobal(instruction)
	case eplan.EAGGREGATEFUNCLOCALNODE:
		return res, self.SetInstructionAggregateFuncLocal(instruction)
	case eplan.ELIMITNODE:
		return res, self.SetInstructionLimit(instruction)
	case eplan.EFILTERNODE:
		return res, self.SetInstructionFilter(instruction)
	case eplan.EUNIONNODE:
		return res, self.SetInstructionUnion(instruction)
	case eplan.EORDERBYLOCALNODE:
		return res, self.SetInstructionOrderByLocal(instruction)
	case eplan.EORDERBYNODE:
		return res, self.SetInstructionOrderBy(instruction)
	case eplan.ESHOWNODE:
		return res, self.SetInstructionShow(instruction)
	case eplan.EBALANCENODE:
		return res, self.SetInstructionBalance(instruction)
	default:
		self.Status = pb.TaskStatus_TODO
		return res, fmt.Errorf("Unknown node type")
	}
	return res, nil
}

func (self *Executor) Run(ctx context.Context, empty *pb.Empty) (*pb.Empty, error) {
	res := &pb.Empty{}
	nodeType := eplan.EPlanNodeType(self.Instruction.TaskType)

	switch nodeType {
	case eplan.ESCANNODE:
		go self.RunScan()
	case eplan.ESELECTNODE:
		go self.RunSelect()
	case eplan.EGROUPBYNODE:
		go self.RunGroupBy()
	case eplan.EJOINNODE:
		go self.RunJoin()
	case eplan.EHASHJOINNODE:
		go self.RunHashJoin()
	case eplan.ESHUFFLENODE:
		go self.RunShuffle()
	case eplan.EDUPLICATENODE:
		go self.RunDuplicate()
	case eplan.EAGGREGATENODE:
		go self.RunAggregate()
	case eplan.EAGGREGATEFUNCGLOBALNODE:
		go self.RunAggregateFuncGlobal()
	case eplan.EAGGREGATEFUNCLOCALNODE:
		go self.RunAggregateFuncLocal()
	case eplan.ELIMITNODE:
		go self.RunLimit()
	case eplan.EFILTERNODE:
		go self.RunFilter()
	case eplan.EORDERBYLOCALNODE:
		go self.RunOrderByLocal()
	case eplan.EORDERBYNODE:
		go self.RunOrderBy()
	case eplan.EUNIONNODE:
		go self.RunUnion()
	case eplan.ESHOWNODE:
		go self.RunShow()
	case eplan.EBALANCENODE:
		go self.RunBalance()
	default:
		return res, fmt.Errorf("Unknown node type")
	}
	return res, nil
}

func (self *Executor) GetOutputChannelLocation(ctx context.Context, location *pb.Location) (*pb.Location, error) {
	if int(location.ChannelIndex) >= len(self.OutputChannelLocations) {
		return nil, fmt.Errorf("ChannelLocation %v not found: %v", location.ChannelIndex, location)
	}
	return self.OutputChannelLocations[location.ChannelIndex], nil
}

///////////////////////////////
func RunExecutor(masterAddress string, address, name string) {
	executorServer = NewExecutor(masterAddress, address, name)
	listener, err := net.Listen("tcp", executorServer.Address)
	if err != nil {
		log.Fatalf("Executor failed to run: %v", err)
	}
	defer listener.Close()
	executorServer.Address = listener.Addr().String()
	logger.Infof("Executor: %v", executorServer.Address)

	go executorServer.Heartbeat()

	grpcS := grpc.NewServer()
	pb.RegisterGueryExecutorServer(grpcS, executorServer)
	grpcS.Serve(listener)
}
