package Executor

import (
	"github.com/vmihailenco/msgpack"
	"github.com/xitongsys/guery/Config"
	"github.com/xitongsys/guery/pb"
)

func (self *Executor) SetRuntime(instruction *pb.Instruction) (err error) {
	var runtime Config.ConfigRuntime
	if err = msgpack.Unmarshal(instruction.RuntimeBytes, &runtime); err != nil {
		return err
	}
	Config.Conf.Runtime = &runtime
	return nil
}
