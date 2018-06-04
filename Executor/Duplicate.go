package Executor

import (
	"io"

	"github.com/vmihailenco/msgpack"
	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/pb"
)

func (self *Executor) SetInstructionDuplicate(instruction *pb.Instruction) (err error) {
	var enode EPlan.EPlanDuplicateNode
	if err = msgpack.Unmarshal(instruction.EncodedEPlanNodeBytes, &enode); err != nil {
		return err
	}
	self.Instruction = instruction
	self.EPlanNode = &enode
	self.InputLocations = []*pb.Location{}
	for i := 0; i < len(enode.Inputs); i++ {
		loc := enode.Inputs[i]
		self.InputLocations = append(self.InputLocations, &loc)
	}
	self.OutputLocations = []*pb.Location{}
	for i := 0; i < len(enode.Outputs); i++ {
		loc := enode.Outputs[i]
		self.OutputLocations = append(self.OutputLocations, &loc)
	}
	return nil
}

func (self *Executor) RunDuplicate() (err error) {
	defer self.Clear()
	enode := self.EPlanNode.(*EPlan.EPlanDuplicateNode)
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

			if enode.Keys != nil && len(enode.Keys) > 0 {
				rb := Util.NewRowsBuffer(md)
				rb.Write(row)
				key, err := CalHashKey(enode.Keys, rb)
				if err != nil {
					return err
				}
				row.AppendKeys(key)
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
