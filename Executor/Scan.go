package Executor

import (
	"fmt"
	"io"

	"github.com/vmihailenco/msgpack"
	"github.com/xitongsys/guery/Connector"
	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/FileSystem/FileReader"
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
		if len(enode.Partitons) > 0 && index >= inputMetadata.GetColumnNumber()-len(enode.Partitons[0].Vals) {
			continue
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
	var i int = 0
	for fi, file := range enode.Files {
		reader, err := FileReader.NewReader(file, inputMetadata)
		//log.Println("[executor.scan]=====file", file)
		if err != nil {
			return err
		}
		for {
			row, err = reader.ReadByColumns(colIndexes)
			//log.Println("[executor.scan]====row", enode.Partitons[0], colIndexes, row, err)
			row.AppendVals(enode.Partitons[fi].Vals...)

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
