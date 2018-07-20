package Agent

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/kardianos/osext"
	"github.com/xitongsys/guery/Config"
	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/pb"
	"google.golang.org/grpc"
)

const (
	TIMEOUT = 10000 //ms
)

type Agent struct {
	MasterAddress     string
	Address           string
	Name              string
	Topology          *Topology
	StartTime         time.Time
	Tasks             *TaskMap
	MaxExecutorNumber int32

	IsStatusChanged bool
}

var agentServer *Agent

func NewAgent(masterAddress string, address, name string) *Agent {
	res := &Agent{
		MasterAddress:     masterAddress,
		Address:           address,
		Name:              name,
		Topology:          NewTopology(),
		Tasks:             NewTaskMap(),
		MaxExecutorNumber: Config.Conf.Runtime.MaxExecutorNumber,
		IsStatusChanged:   true,
	}
	return res
}

func (self *Agent) KillTask(ctx context.Context, task *pb.Task) (*pb.Empty, error) {
	res := &pb.Empty{}
	if task == nil {
		return res, nil
	}
	task = self.Tasks.PopTask(task.TaskId)
	if task == nil {
		return res, nil
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

///////////////////////////////
func RunAgent(masterAddress string, address, name string) {
	agentServer = NewAgent(masterAddress, address, name)
	listener, err := net.Listen("tcp", agentServer.Address)
	if err != nil {
		log.Fatalf("Agent failed to run: %v", err)
	}
	defer listener.Close()
	agentServer.Address = listener.Addr().String()
	Logger.Infof("Agent: %v", agentServer.Address)

	go agentServer.Heartbeat()

	grpcS := grpc.NewServer()
	pb.RegisterGueryAgentServer(grpcS, agentServer)
	grpcS.Serve(listener)
}
