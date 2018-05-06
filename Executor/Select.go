package Executor

import (
	"bytes"
	"encoding/gob"

	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/pb"
)

func (self *Executor) SetInstructionSelect(instruction *pb.Instruction) (err error) {
	var enode EPlan.EPlanSelectNode
	if err = gob.NewDecoder(bytes.NewBufferString(instruction.EncodedEPlanNodeBytes)).Decode(&enode); err != nil {
		return err
	}
	self.Instruction = instruction
	self.EPlanNode = &enode
	return nil
}

func (self *Executor) RunSelect() (err error) {
	return nil
}
