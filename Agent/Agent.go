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
	"time"

	"github.com/kardianos/osext"
	"github.com/xitongsys/guery/Config"
	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/pb"
	"google.golang.org/grpc"
)

type Agent struct {
	MasterAddress string
	Address       string
	Name          string
	Topology      *Topology
	StartTime     time.Time
}

var agentServer *Executor

func NewAgent(masterAddress string, address, name string) *Agent {
	res := &Agent{
		MasterAddress: masterAddress,
		Address:       address,
		Name:          name,
		Topology:      NewTopology(),
	}
	return res
}

func (self *Agent) Clear() {
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

func (self *Agent) Duplicate(ctx context.Context, em *pb.Empty) (*pb.Empty, error) {
	res := &pb.Empty{}
	exeFullName, _ := osext.Executable()

	command := exec.Command(exeFullName,
		fmt.Sprintf("executor"),
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
