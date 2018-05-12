package Executor

import (
	"fmt"

	"github.com/vmihailenco/msgpack"
	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/pb"
)

func (self *Executor) SetInstructionJoin(instruction *pb.Instruction) (err error) {
	var enode EPlan.EPlanJoinNode
	if err = msgpack.Unmarshal(instruction.EncodedEPlanNodeBytes, &enode); err != nil {
		return err
	}
	self.Instruction = instruction
	self.EPlanNode = &enode
	self.InputLocations = []*pb.Location{&enode.LeftInput, &enode.RightInput}
	self.OutputLocations = []*pb.Location{&enode.Output}
	return nil
}

func (self *Executor) RunJoin() (err error) {
	defer self.Clear()
	if len(self.Readers) != 2 {
		return fmt.Errorf("join readers number %v <> 2", len(self.Readers))
	}

	//read md
	mds := make([]*Util.Metadata, 2)
	for i, reader := range self.Readers {
		mds[i] = &Util.Metadata{}
		if err = Util.ReadObject(reader, mds[i]); err != nil {
			return err
		}
	}

	return nil
}
