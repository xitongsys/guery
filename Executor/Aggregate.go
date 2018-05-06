package Executor

import (
	"bytes"
	"encoding/gob"

	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/pb"
)

func (self *Executor) SetInstructionAggregate(instruction *pb.Instruction) (err error) {
	var enode EPlan.EPlanAggregateNode
	if err = gob.NewDecoder(bytes.NewBufferString(instruction.EncodedEPlanNodeBytes)).Decode(&enode); err != nil {
		return err
	}
	self.Instruction = instruction
	self.EPlanNode = &enode
	self.InputLocations = []*pb.Location{}
	for _, loc := range enode.Inputs {
		self.InputLocations = append(self.InputLocations, &loc)
	}
	self.OutputLocations = []*pb.Location{&enode.Output}
	return nil
}

func (self *Executor) RunAggregate() (err error) {
	return nil
}
