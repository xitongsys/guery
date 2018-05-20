package Executor

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/kardianos/osext"
	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/pb"
	"google.golang.org/grpc"
)

type Executor struct {
	MasterAddress string

	Address string
	Name    string

	Instruction                                   *pb.Instruction
	EPlanNode                                     EPlan.ENode
	InputLocations, OutputLocations               []*pb.Location
	InputChannelLocations, OutputChannelLocations []*pb.Location
	Readers                                       []io.Reader
	Writers                                       []io.Writer

	Status          int32
	IsStatusChanged bool

	DoneChan chan int
}

var executorServer *Executor

func NewExecutor(masterAddress string, address, name string) *Executor {
	res := &Executor{
		MasterAddress: masterAddress,
		Address:       address,
		Name:          name,
		DoneChan:      make(chan int),
	}
	return res
}

func (self *Executor) Clear() {
	self.Instruction = nil
	self.EPlanNode = nil
	self.InputLocations, self.OutputLocations = []*pb.Location{}, []*pb.Location{}
	self.InputChannelLocations, self.OutputChannelLocations = []*pb.Location{}, []*pb.Location{}
	for _, writer := range self.Writers {
		writer.(io.WriteCloser).Close()
	}
	self.Readers, self.Writers = []io.Reader{}, []io.Writer{}
	self.Status = 0
	self.IsStatusChanged = true

	select {
	case <-self.DoneChan:
	default:
		close(self.DoneChan)
	}
}

func (self *Executor) Duplicate(ctx context.Context, em *pb.Empty) (*pb.Empty, error) {
	exeFullName, _ := osext.Executable()
	command := exec.Command(exeFullName,
		fmt.Sprintf("executor"),
		fmt.Sprintf("--master %v:%v", self.MasterAddress),
		fmt.Sprintf("--address %v", strings.Split(self.Address, ":")[0]+":0"),
		fmt.Sprintf("--name %v", self.Name),
	)
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	err := command.Start()
	return nil, err
}

func (self *Executor) Quit(ctx context.Context, em *pb.Empty) (*pb.Empty, error) {
	os.Exit(0)
	return nil, nil
}

func (self *Executor) Restart(ctx context.Context, em *pb.Empty) (*pb.Empty, error) {
	self.Duplicate(context.Background(), nil)
	time.Sleep(time.Second)
	self.Quit(ctx, em)
	return nil, nil
}

func (self *Executor) SendInstruction(ctx context.Context, instruction *pb.Instruction) (*pb.Empty, error) {
	res := &pb.Empty{}

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
	case EPlan.EDUPLICATENODE:
		return res, self.SetInstructionDuplicate(instruction)
	case EPlan.EAGGREGATENODE:
		return res, self.SetInstructionAggregate(instruction)
	case EPlan.ELIMITNODE:
		return res, self.SetInstructionLimit(instruction)
	case EPlan.EFILITERNODE:
		return res, self.SetInstructionFiliter(instruction)
	case EPlan.EUNIONNODE:
		return res, self.SetInstructionUnion(instruction)
	case EPlan.EORDERBYLOCALNODE:
		return res, self.SetInstructionOrderByLocal(instruction)
	case EPlan.EORDERBYNODE:
		return res, self.SetInstructionOrderBy(instruction)
	default:
		self.Status = 0
		return res, fmt.Errorf("Unknown node type")
	}
	return res, nil
}

func (self *Executor) Run(ctx context.Context, empty *pb.Empty) (*pb.Empty, error) {
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
	case EPlan.EDUPLICATENODE:
		go self.RunDuplicate()
	case EPlan.EAGGREGATENODE:
		go self.RunAggregate()
	case EPlan.ELIMITNODE:
		go self.RunLimit()
	case EPlan.EFILITERNODE:
		go self.RunFiliter()
	case EPlan.EORDERBYLOCALNODE:
		go self.RunOrderByLocal()
	case EPlan.EORDERBYNODE:
		go self.RunOrderBy()
	case EPlan.EUNIONNODE:
		go self.RunUnion()
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
	Logger.Infof("Executor: %v", executorServer.Address)

	go executorServer.Heartbeat()

	grpcS := grpc.NewServer()
	pb.RegisterGueryExecutorServer(grpcS, executorServer)
	grpcS.Serve(listener)
}
