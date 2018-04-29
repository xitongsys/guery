package Executor

import (
	"os"
	"os/exec"

	"github.com/kardianos/osext"
	"github.com/xitongsys/guery/pb"
)

type Executor struct {
	MasterAddress string
	MasterPort    int32

	DataCenter string
	Rack       string
	Address    string
	Port       int32

	Name            string
	Inputs          []*Channel
	Outputs         []*Channel
	Status          int32
	IsStatusChanged bool
}

func (self *Executor) Duplicate() {
	exeFullName, _ := osext.Executable()
	command := exec.Command(exeFullName,
		fmt.Sprintf("executor"),
		fmt.Sprintf("--master %v:%v", self.MasterAddress, self.MasterPort),
		fmt.Sprintf("--datacenter %v", self.DataCenter),
		fmt.Sprintf("--rack %v", self.Rack),
		fmt.Sprintf("--address %v:%v", self.Address, self.Port),
		fmt.Sprintf("--name %v", self.Name),
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
