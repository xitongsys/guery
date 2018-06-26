package Executor

import (
	"fmt"
	"io"

	"github.com/vmihailenco/msgpack"
	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Split"
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

	rbReader, rbWriter := Split.NewSplitBuffer(md, reader, nil), Split.NewSplitBuffer(enode.Metadata, nil, writer)
	defer func() {
		rbWriter.Flush()
	}()

	//write

	if enode.IsAggregate {
		var sp, spAgg *Split.Split
		for err == nil {
			sp, err = rbReader.ReadSplit()
			if err == io.EOF {
				err = nil
				break
			}
			if err != nil {
				break
			}

			for i := 0; i < sp.GetRowsNumber(); i++ {
				if spAgg == nil {
					spAgg = Split.NewSplit(md)
					spAgg.SplitKeys = sp.GetKeys()
					continue
				}
				if spAgg.GetKeyString() == sp.GetKeyString(i) {
					spAgg.Append(sp, i)

				} else {
					var row []interface{}
					if row, err = self.CalSelectItems(enode, rg); err != nil {
						break
					}
					if err = rbWriter.WriteValues(row); err != nil {
						break
					}

					spAgg == nil
					i--
				}
			}
		}

		if err == nil && spAgg != nil {
			var row []interface{}
			if row, err = self.CalSelectItems(enode, rg); err != nil {
				break
			}
			err = rbWriter.WriteValues(row)
		}

	} else {
		for {
			var sp *Split.Split
			sp, err = rbReader.ReadSplit()
			if err == io.EOF {
				err = nil
				break
			}
			if err != nil {
				break
			}

			if row, err = self.CalSelectItems(enode, rg); err != nil {
				break
			}

			if err = rbWriter.WriteValues(row); err != nil {
				Logger.Errorf("failed to WriteRow %v", err)
				break
			}
		}
	}

	Logger.Infof("RunSelect finished")
	return err
}

func (self *Executor) CalSelectItems(enode *EPlan.EPlanSelectNode, sp *Split.Split, index int) ([]interface{}, error) {
	var err error
	var res interface{}
	var row []interface{}
	for _, item := range enode.SelectItems {
		res, err = item.Result(sp, index)
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}
		row = append(row, res.([]interface{})...)
	}
	return row, err
}
