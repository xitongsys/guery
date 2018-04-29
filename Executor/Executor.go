package Executor

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

func (self *Executor) Duplicate(ctx context.Context, em *pb.Empty) (*pb.Empty, error) {
	exeFullName, _ := osext.Executable()
	command := exec.Command(exeFullName,
		fmt.Sprintf("executor"),
		fmt.Sprintf("--master %v:%v", self.MasterAddress),
		fmt.Sprintf("--datacenter %v", self.DataCenter),
		fmt.Sprintf("--rack %v", self.Rack),
		fmt.Sprintf("--address %v", strings.Split(self.Address, ":")[0]+":0"),
		fmt.Sprintf("--name %v", self.Name),
	)
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	err := command.Start()
	return nil, err
}

func (self *Executor) Quit(ctx context.Context, em *pb.Empty) (*pb.Empty, error) {
	self.Duplicate(context.Background(), nil)
	time.Sleep(10 * time.Second)
	os.Exit(0)
	return nil, nil
}

func (self *Executor) SendInstruction(ctx context.Context, instruction *pb.Instruction) (*pb.Empty, error) {
	return nil, nil
}

///////////////////////////////
func RunExecutor(masterAddress string, dc, rack, address, name string) {
	executorServer = NewExecutor(masterAddress, dc, rack, address, name)
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
