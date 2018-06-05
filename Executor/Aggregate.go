package Executor

import (
	"io"

	"github.com/vmihailenco/msgpack"
	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/pb"
)

func (self *Executor) SetInstructionAggregate(instruction *pb.Instruction) (err error) {
	var enode EPlan.EPlanAggregateNode
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
	self.OutputLocations = []*pb.Location{&enode.Output}
	return nil
}

func (self *Executor) RunAggregate() (err error) {
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

	rbWriter := Util.NewRowsBuffer(md, nil, writer)

	//write rows
	var row *Util.Row
	for _, reader := range self.Readers {
		rbReader := Util.NewRowsBuffer(md, reader, nil)
		for {
			row, err = rbReader.ReadRow()
			if err == io.EOF {
				err = nil
				break
			}
			if err != nil {
				return err
			}
			rbWriter.WriteRow(row)
		}
	}
	rbWriter.Flush()
	Logger.Infof("RunAggregate finished")
	return nil
}
