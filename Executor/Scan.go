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
	defer self.Clear()

	if self.Instruction == nil {
		return fmt.Errorf("No Instruction")
	}

	enode := self.EPlanNode.(*EPlan.EPlanScanNode)

	connector, err := Connector.NewConnector(enode.Catalog, enode.Schema, enode.Table)
	if err != nil {
		return err
	}

	partitionInfo := connector.GetPartitionInfo()

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

	//send rows
	var row *Util.Row
	for {
		row, err = connector.ReadByColumns(colIndexes)
		//log.Printf("===%v, %v\n", row, err)
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		if err = Util.WriteRow(self.Writers[i], row); err != nil {
			break
		}

		i++
		i = i % ln
	}

	for i := 0; i < ln; i++ {
		Util.WriteEOFMessage(self.Writers[i])
		self.Writers[i].(io.WriteCloser).Close()
	}
	Logger.Infof("RunScan finished")
	return err
}
