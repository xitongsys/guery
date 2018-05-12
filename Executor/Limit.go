package Executor

import (
	"io"

	"github.com/vmihailenco/msgpack"
	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/pb"
)

func (self *Executor) SetInstructionLimit(instruction *pb.Instruction) (err error) {
	var enode EPlan.EPlanLimitNode
	if err = msgpack.Unmarshal(instruction.EncodedEPlanNodeBytes, &enode); err != nil {
		return err
	}
	self.Instruction = instruction
	self.EPlanNode = &enode
	self.InputLocations = []*pb.Location{}
	self.InputLocations = append(self.InputLocations, &enode.Input)
	self.OutputLocations = []*pb.Location{&enode.Output}
	return nil
}

func (self *Executor) RunLimit() (err error) {
	defer self.Clear()

	writer := self.Writers[0]
	md := &Util.Metadata{}
	//read md
	for _, reader := range self.Readers {
		if err = Util.ReadObject(reader, md); err != nil {
			return err
		}
	}

	//write md
	if err = Util.WriteObject(writer, md); err != nil {
		return err
	}

	//write rows
	var row *Util.Row
	readRowCnt := int64(0)
	enode := self.EPlanNode.(*EPlan.EPlanLimitNode)
	for _, reader := range self.Readers {
		for readRowCnt < *(enode.LimitNumber) {
			row, err = Util.ReadRow(reader)
			//Logger.Infof("===%v, %v", row, err)
			if err == io.EOF || readRowCnt >= *(enode.LimitNumber) {
				err = nil
				break
			}
			if err != nil {
				return err
			}
			readRowCnt++
			Util.WriteRow(writer, row)
		}
	}

	Util.WriteEOFMessage(writer)
	Logger.Infof("RunAggregate finished")
	return nil
}
