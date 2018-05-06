package Executor

import (
	"bytes"
	"encoding/gob"

	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/pb"
)

func (self *Executor) SetInstructionGroupBy(instruction *pb.Instruction) (err error) {
	var enode EPlan.EPlanGroupByNode
	if err = gob.NewDecoder(bytes.NewBufferString(instruction.EncodedEPlanNodeBytes)).Decode(&enode); err != nil {
		return err
	}
	return nil
}

func (self *Executor) RunGroupBy() (err error) {
	return nil
}
