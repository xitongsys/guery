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

func (self *Executor) CancelTask() error {
	exeFullName, _ := osext.Executable()
	command := exec.Command(exeFullName,
		fmt.Sprintf("-MasterAddress=%v", self.MasterAddress),
		fmt.Sprintf("-MasterPort=%v", self.MasterPort),
		fmt.Sprintf("-DataCenter=%v", self.DataCenter),
		fmt.Sprintf("-Rack=%v", self.Rack),
		fmt.Sprintf("-Address=%v", self.Address),
		fmt.Sprintf("-Port=%v", self.Port),
		fmt.Sprintf("-Name=%v", self.Name),
	)
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	if err := command.Start(); err != nil {
		return err
	}
	os.Exit(0)
}

func (self *Executor) SendInstruction(instruction pb.Instruction) {

}
