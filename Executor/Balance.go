package Executor

import (
	"io"

	"github.com/vmihailenco/msgpack"
	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Row"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/pb"
)

func (self *Executor) SetInstructionBalance(instruction *pb.Instruction) (err error) {
	var enode EPlan.EPlanBalanceNode
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

func (self *Executor) RunBalance() (err error) {
	defer self.Clear()
	//read md
	md := &Metadata.Metadata{}
	for _, reader := range self.Readers {
		if err = Util.ReadObject(reader, md); err != nil {
			return err
		}
	}

	mdOutput := md.Copy()

	//write md
	for _, writer := range self.Writers {
		if err = Util.WriteObject(writer, mdOutput); err != nil {
			return err
		}
	}

	rbWriters := make([]*Row.RowsBuffer, len(self.Writers))
	for i, writer := range self.Writers {
		rbWriters[i] = Row.NewRowsBuffer(mdOutput, nil, writer)
	}

	defer func() {
		for _, rbWriter := range rbWriters {
			rbWriter.Flush()
		}
	}()

	//write rows
	var row *Row.Row
	i, wn := 0, len(rbWriters)
	for _, reader := range self.Readers {
		rbReader := Row.NewRowsBuffer(md, reader, nil)
		for {
			row, err = rbReader.ReadRow()
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}

			rbWriter := rbWriters[i]
			i++
			i = i % wn

			if err = rbWriter.WriteRow(row); err != nil {
				return err
			}
		}
	}

	return nil
}