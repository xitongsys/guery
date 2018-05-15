package Executor

import (
	"fmt"
	"io"

	"github.com/vmihailenco/msgpack"
	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/pb"
)

func (self *Executor) SetInstructionUnion(instruction *pb.Instruction) (err error) {
	var enode EPlan.EPlanUnionNode
	if err = msgpack.Unmarshal(instruction.EncodedEPlanNodeBytes, &enode); err != nil {
		return err
	}
	self.Instruction = instruction
	self.EPlanNode = &enode
	self.InputLocations = []*pb.Location{&enode.LeftInput, &enode.RightInput}
	self.OutputLocations = []*pb.Location{&enode.Output}
	return nil
}

func (self *Executor) RunUnion() (err error) {
	defer self.Clear()
	writer := self.Writers[0]
	enode := self.EPlanNode.(*EPlan.EPlanUnionNode)

	//read md
	if len(self.Readers) != 2 {
		return fmt.Errorf("union readers number %v <> 2", len(self.Readers))
	}

	md := &Util.Metadata{}
	if len(self.Readers) != 2 {
		return fmt.Errorf("union input number error")
	}
	for _, reader := range self.Readers {
		if err = Util.ReadObject(reader, md); err != nil {
			return err
		}
	}

	//write md
	if err = Util.WriteObject(writer, enode.Metadata); err != nil {
		return err
	}

	//write rows
	var row *Util.Row
	for _, reader := range self.Readers {
		for {
			row, err = Util.ReadRow(reader)
			if err == io.EOF {
				err = nil
				break
			}
			if err != nil {
				return err
			}
			rb := Util.NewRowsBuffer(md)
			rb.Write(row)
			if err = Util.WriteRow(writer, row); err != nil {
				return err
			}
		}
	}

	Util.WriteEOFMessage(writer)

	Logger.Infof("RunUnion finished")
	return err
}
