package Executor

import (
	"bytes"
	"encoding/gob"

	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/pb"
)

func (self *Executor) SetInstructionDuplicate(instruction *pb.Instruction) (err error) {
	var enode EPlan.EPlanDuplicateNode
	if err = gob.NewDecoder(bytes.NewBuffer(instruction.EncodedEPlanNodeBytes)).Decode(&enode); err != nil {
		return err
	}
	return nil
}

func (self *Executor) RunDuplicate() (err error) {
	return nil
}
