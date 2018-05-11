package Executor

import (
	"bytes"
	"encoding/gob"
	"io"

	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/pb"
)

func (self *Executor) SetInstructionDuplicate(instruction *pb.Instruction) (err error) {
	var enode EPlan.EPlanDuplicateNode
	if err = gob.NewDecoder(bytes.NewBuffer(instruction.EncodedEPlanNodeBytes)).Decode(&enode); err != nil {
		return err
	}
	self.Instruction = instruction
	self.EPlanNode = &enode
	self.InputLocations = []*pb.Location{}
	for _, loc := range enode.Inputs {
		self.InputLocations = append(self.InputLocations, &loc)
	}
	self.OutputLocations = []*pb.Location{}
	for _, loc := range enode.Outputs {
		self.OutputLocations = append(self.OutputLocations, &loc)
	}
	return nil
}

func (self *Executor) RunDuplicate() (err error) {
	defer self.Clear()

	//read md
	md := &Util.Metadata{}
	for _, reader := range self.Readers {
		if err = Util.ReadObject(reader, md); err != nil {
			return err
		}
	}

	//write md
	for _, writer := range self.Writers {
		if err = Util.WriteObject(writer, md); err != nil {
			return err
		}
	}

	//write rows
	var row *Util.Row
	for _, reader := range self.Readers {
		for {
			row, err = Util.ReadRow(reader)
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}

			for _, writer := range self.Writers {
				if err = Util.WriteRow(writer, row); err != nil {
					return err
				}
			}
		}
	}

	for _, writer := range self.Writers {
		if err = Util.WriteEOFMessage(writer); err != nil {
			return err
		}
	}
	return nil
}
