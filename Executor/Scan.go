package Executor

import (
	"fmt"
	"io"

	"github.com/vmihailenco/msgpack"
	"github.com/xitongsys/guery/Connector"
	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/pb"
)

func (self *Executor) SetInstructionScan(instruction *pb.Instruction) error {
	Logger.Infof("set instruction scan")
	var enode EPlan.EPlanScanNode
	var err error
	if err = msgpack.Unmarshal(instruction.EncodedEPlanNodeBytes, &enode); err != nil {
		return err
	}

	self.EPlanNode = &enode
	self.Instruction = instruction
	for i := 0; i < len(enode.Outputs); i++ {
		loc := enode.Outputs[i]
		self.OutputLocations = append(self.OutputLocations, &loc)
	}
	return nil
}

func (self *Executor) RunScan() (err error) {
	defer func() {
		for i := 0; i < len(self.Writers); i++ {
			Util.WriteEOFMessage(self.Writers[i])
			self.Writers[i].(io.WriteCloser).Close()
		}
		self.Clear()
	}()

	if self.Instruction == nil {
		return fmt.Errorf("No Instruction")
	}

	enode := self.EPlanNode.(*EPlan.EPlanScanNode)

	connector, err := Connector.NewConnector(enode.Catalog, enode.Schema, enode.Table)
	if err != nil {
		return err
	}

	ln := len(self.Writers)
	i := 0

	//send metadata
	for i := 0; i < ln; i++ {
		if err = Util.WriteObject(self.Writers[i], enode.Metadata); err != nil {
			return err
		}
	}

	colIndexes := []int{}
	inputMetadata := connector.GetMetadata()
	for _, c := range enode.Metadata.Columns {
		cn := c.ColumnName
		index, err := inputMetadata.GetIndexByName(cn)
		if err != nil {
			return err
		}
		colIndexes = append(colIndexes, index)
	}

	rbWriters := make([]*Util.RowsBuffer, len(self.Writers))
	for i, writer := range self.Writers {
		rbWriters[i] = Util.NewRowsBuffer(enode.Metadata, nil, writer)
	}

	defer func() {
		for _, rbWriter := range rbWriters {
			rbWriter.Flush()
		}
	}()

	//send rows

	var row *Util.Row
	for _, file := range enode.Files {

		for {
			row, err = connector.ReadByColumns(colIndexes)
			if err == io.EOF {
				err = nil
				break
			}
			if err != nil {
				return err
			}

			if err = rbWriters[i].WriteRow(row); err != nil {
				return err
			}

			i++
			i = i % ln
		}
	}

	Logger.Infof("RunScan finished")
	return err
}
