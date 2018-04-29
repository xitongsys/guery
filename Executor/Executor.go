package Executor

import (
	"log"
	"net"
	"os"
	"os/exec"

	"github.com/kardianos/osext"
	"github.com/xitongsys/guery/pb"
	"google.golang.org/grpc"
)

type Executor struct {
	MasterAddress string

	DataCenter string
	Rack       string
	Address    string

	Name            string
	Inputs          []*Channel
	Outputs         []*Channel
	Status          int32
	IsStatusChanged bool
}

var executorServer *Executor

func NewExecutor(masterAddress string, dc, rack, address, name string) *Executor {
	res := &Executor{
		MasterAddress: masterAddress,
		DataCenter:    dc,
		Rack:          rack,
		Address:       address,
		Name:          name,
	}
	return res
}

func (self *Executor) Duplicate() {
	exeFullName, _ := osext.Executable()
	command := exec.Command(exeFullName,
		fmt.Sprintf("executor"),
		fmt.Sprintf("--master %v:%v", self.MasterAddress, self.MasterPort),
		fmt.Sprintf("--datacenter %v", self.DataCenter),
		fmt.Sprintf("--rack %v", self.Rack),
		fmt.Sprintf("--name %v", self.Name),
		fmt.Sprintf("--address %v", strings.Split(self.Address, ":")[0]+":0"),
	)
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Start()
}

func (self *Executor) Quit() {
	self.Duplicate()
	time.Sleep(10 * time.Second)
	os.Exit(0)
}

func (self *Executor) SendInstruction(instruction pb.Instruction) {
}

///////////////////////////////
func RunExecutor(masterAddress string, dc, rack string, name string) {
	executorServer = NewExecutor(masterAddress, dc, rack, name)
	listener, err := net.Listen("tcp", executorServer.Address)
	if err != nil {
		log.Fatalf("Executor failed to run: %v", err)
	}
	defer listener.Close()
	executorServer.Address = listener.Addr().String()

	grpcS := grpc.NewServer()
	pb.RegisterGueryExecutorServer(grpcS, executorServer)
	grpcS.Serve(listener)
}
