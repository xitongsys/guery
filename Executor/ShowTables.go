package Executor

import (
	"fmt"
	"io"

	"github.com/vmihailenco/msgpack"
	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Row"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/pb"
)

func (self *Executor) SetInstructionShowTables(instruction *pb.Instruction) error {
	Logger.Infof("set instruction show tables")
	var enode EPlan.EPlanShowTablesNode
	var err error
	if err = msgpack.Unmarshal(instruction.EncodedEPlanNodeBytes, &enode); err != nil {
		return err
	}

	self.EPlanNode = &enode
	self.Instruction = instruction
	self.InputLocations = []*pb.Location{}
	self.OutputLocations = append(self.OutputLocations, &enode.Output)
	return nil
}

func (self *Executor) RunShowTables() (err error) {
	defer self.Clear()

	if self.Instruction == nil {
		return fmt.Errorf("No Instruction")
	}

	enode := self.EPlanNode.(*EPlan.EPlanShowTablesNode)
	md := &Metadata.Metadata{}
	reader := self.Readers[0]
	writer := self.Writers[0]
	if err = Util.ReadObject(reader, md); err != nil {
		return err
	}

	//write metadata
	if err = Util.WriteObject(writer, md); err != nil {
		return err
	}

	rbReader := Row.NewRowsBuffer(md, reader, nil)
	rbWriter := Row.NewRowsBuffer(md, nil, writer)

	//writer rows
	var row *Row.Row
	for {

		if err = rbWriter.WriteRow(row); err != nil {
			return err
		}
	}

	if err = rbWriter.Flush(); err != nil {
		return err
	}

	Logger.Infof("RunShowTables finished")
	return err

}
