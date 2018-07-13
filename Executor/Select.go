package Executor

import (
	"fmt"
	"io"
	"os"
	"runtime/pprof"
	"time"

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
	fname := fmt.Sprintf("executor_%v_select_%v_cpu.pprof", self.Name, time.Now().Format("20060102150405"))
	f, _ := os.Create(fname)
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

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

	//init
	for _, item := range enode.SelectItems {
		if err := item.Init(md); err != nil {
			return err
		}
	}

	//write rows
	var row *Row.Row
	keys := map[string]*Row.RowsGroup{}
	resMap := map[string]*Row.Row{}
	if enode.IsAggregate {
		for {
			row, err = rbReader.ReadRow()
			if err == io.EOF {
				err = nil
				for key, rg := range keys {
					if rg.GetRowsNum() > 0 {
						if resMap[key], err = self.CalSelectItems(enode, rg); err != nil {
							break
						}
					}
				}

				for _, row := range resMap {
					if row != nil {
						if err = rbWriter.WriteRow(row); err != nil {
							break
						}
					}
				}

				break
			}
			if err != nil {
				break
			}

			key := row.GetKeyString()
			if _, ok := keys[key]; !ok {
				keys[key] = Row.NewRowsGroup(md)
			}
			rg := keys[key]
			rg.Write(row)

			if rg.GetRowsNum() > Row.ROWS_BUFFER_SIZE {
				if resMap[key], err = self.CalSelectItems(enode, rg); err != nil {
					break
				}
				rg.ClearRows()
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
			rg := Row.NewRowsGroup(md)
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
	key := rg.GetKeyString()
	for _, item := range enode.SelectItems {
		rg.Reset()
		res, err = item.Result(rg)
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}
		if item.IsAggregate() {
			r := res.([]interface{})[0].(map[string]interface{})
			if val, ok := r[key]; !ok {
				return nil, fmt.Errorf("CalSelectItems Error")

			} else {
				row.AppendVals(val)
			}
		} else {
			row.AppendVals(res.([]interface{})...)
		}
	}
	return row, err
}
