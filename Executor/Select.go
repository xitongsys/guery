package Executor

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"strings"

	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/pb"
)

func (self *Executor) SetInstructionSelect(instruction *pb.Instruction) (err error) {
	var enode EPlan.EPlanSelectNode
	if err = gob.NewDecoder(bytes.NewBuffer(instruction.EncodedEPlanNodeBytes)).Decode(&enode); err != nil {
		return err
	}
	self.Instruction = instruction
	self.EPlanNode = &enode
	self.InputLocations = []*pb.Location{&enode.Input}
	self.OutputLocations = []*pb.Location{&enode.Output}
	return nil
}

func (self *Executor) RunSelect() (err error) {

	if self.Instruction == nil {
		return fmt.Errorf("No Instruction")
	}
	enode := self.EPlanNode.(*EPlan.EPlanSelectNode)

	md := &Util.Metadata{}
	reader := self.Readers[0]
	writer := self.Writers[0]
	if err = Util.ReadObject(reader, md); err != nil {
		return err
	}

	colNames := []string{}
	isAgg := false
	for _, item := range enode.SelectItems {
		if item.IsAggregate() {
			isAgg = true
		}
		for _, cname := range item.GetNames() {
			names := strings.Split(cname, ".")
			name := names[len(names)-1]
			colNames = append(colNames, name)
		}
	}

	//write metadata--------------------TODO
	smd := Util.NewMetadata(md.Name, md.ColumnNames, md.ColumnTypes)
	if err = Util.WriteObject(writer, smd); err != nil {
		return err
	}

	//write rows
	var row *Util.Row
	var rowsBuf *Util.RowsBuffer
	if isAgg {
		for {
			row, err = Util.ReadRow(reader)
			if err == io.EOF {
				err = nil
				if rowsBuf != nil {
					if row, err = self.CalSelectItems(enode, rowsBuf); err == nil {
						Util.WriteRow(writer, row)
					}
				}
				break
			}
			if err != nil {
				break
			}

			if rowsBuf == nil {
				rowsBuf = Util.NewRowsBuffer(md)
				rowsBuf.Write(row)

			} else {
				if rowsBuf.Key == row.Key {
					rowsBuf.Write(row)

				} else {
					if row, err = self.CalSelectItems(enode, rowsBuf); err != nil {
						break
					}
					rowsBuf = Util.NewRowsBuffer(md)
					rowsBuf.Write(row)
				}
			}
		}

	} else {
		for {
			row, err = Util.ReadRow(reader)

			if err == io.EOF {
				err = nil
				break
			}
			if err != nil {
				break
			}
			rowsBuf = Util.NewRowsBuffer(md)
			rowsBuf.Write(row)
			if row, err = self.CalSelectItems(enode, rowsBuf); err != nil {
				break
			}

			//Logger.Infof("===%v, %v", row, err)
			if err = Util.WriteRow(writer, row); err != nil {
				Logger.Errorf("failed to WriteRow %v", err)
				break
			}
		}
	}
	Util.WriteEOFMessage(writer)
	writer.(io.WriteCloser).Close()

	Logger.Infof("RunSelect finished")
	return err
}

func (self *Executor) CalSelectItems(enode *EPlan.EPlanSelectNode, rowsBuf *Util.RowsBuffer) (*Util.Row, error) {
	var err error
	var res interface{}
	row := Util.NewRow()
	for _, item := range enode.SelectItems {
		rowsBuf.Reset()
		res, err = item.Result(rowsBuf)
		if err != nil {
			break
		}
		row.AppendVals(res.([]interface{})...)
	}
	return row, err
}
