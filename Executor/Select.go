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

func (self *Executor) SetInstructionSelect(instruction *pb.Instruction) (err error) {
	var enode EPlan.EPlanSelectNode
	if err = msgpack.Unmarshal(instruction.EncodedEPlanNodeBytes, &enode); err != nil {
		return err
	}
	self.Instruction = instruction
	self.EPlanNode = &enode
	self.InputLocations = []*pb.Location{&enode.Input}
	self.OutputLocations = []*pb.Location{&enode.Output}
	return nil
}

func (self *Executor) RunSelect() (err error) {
	defer self.Clear()

	if self.Instruction == nil {
		return fmt.Errorf("No Instruction")
	}
	enode := self.EPlanNode.(*EPlan.EPlanSelectNode)

	md := &Metadata.Metadata{}
	reader := self.Readers[0]
	writer := self.Writers[0]
	if err = Util.ReadObject(reader, md); err != nil {
		return err
	}

	//write metadata
	if err = Util.WriteObject(writer, enode.Metadata); err != nil {
		return err
	}

	rbReader, rbWriter := Row.NewRowsBuffer(md, reader, nil), Row.NewRowsBuffer(enode.Metadata, nil, writer)
	defer func() {
		rbWriter.Flush()
	}()

	//write rows
	var row *Row.Row
	var rg *Row.RowsGroup
	if enode.IsAggregate {
		for {
			row, err = rbReader.ReadRow()

			if err == io.EOF {
				err = nil
				if rg != nil {
					if row, err = self.CalSelectItems(enode, rg); err == nil {
						rbWriter.WriteRow(row)
					}
				}
				break
			}
			if err != nil {
				break
			}

			if rg == nil {
				rg = Row.NewRowsGroup(md)
				rg.Write(row)

			} else {

				if rg.GetKeyString() == row.GetKeyString() {
					rg.Write(row)

				} else {
					var row2 *Row.Row
					if row2, err = self.CalSelectItems(enode, rg); err != nil {
						break
					}
					rbWriter.WriteRow(row2)

					rg = Row.NewRowsGroup(md)
					rg.Write(row)
				}
			}
		}

	} else {
		for {
			row, err = rbReader.ReadRow()
			if err == io.EOF {
				err = nil
				break
			}
			if err != nil {
				break
			}
			rg = Row.NewRowsGroup(md)
			rg.Write(row)

			if row, err = self.CalSelectItems(enode, rg); err != nil {
				break
			}

			if err = rbWriter.WriteRow(row); err != nil {
				Logger.Errorf("failed to WriteRow %v", err)
				break
			}
		}
	}

	Logger.Infof("RunSelect finished")
	return err
}

func (self *Executor) CalSelectItems(enode *EPlan.EPlanSelectNode, rg *Row.RowsGroup) (*Row.Row, error) {
	var err error
	var res interface{}
	row := Row.NewRow()
	for _, item := range enode.SelectItems {
		rg.Reset()
		res, err = item.Result(rg)
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}
		row.AppendVals(res.([]interface{})...)
	}
	return row, err
}
