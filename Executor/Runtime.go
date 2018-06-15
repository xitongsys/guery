package Executor

import (
	"github.com/vmihailenco/msgpack"
	"github.com/xitongsys/guery/Config"
)

func (self *Executor) SetRuntime(instruction *pb.Instruction) error {
	var runtime Config.ConfigRuntime
	if err = msgpack.Unmarshal(instruction.RuntimeBytes, &runtime); err != nil {
		return err
	}
	Config.Conf.Runtime = &runtime
	return nil
}
