package executor

import (
	"github.com/vmihailenco/msgpack"
	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/pb"
)

func (self *Executor) SetRuntime(instruction *pb.Instruction) (err error) {
	var runtime config.ConfigRuntime
	if err = msgpack.Unmarshal(instruction.RuntimeBytes, &runtime); err != nil {
		return err
	}
	config.Conf.Runtime = &runtime
	return nil
}
