package Executor

import (
	"fmt"
	"io"
	"log"
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

func (self *Executor) SetInstructionAggregateFuncLocal(instruction *pb.Instruction) (err error) {
	var enode EPlan.EPlanAggregateFuncLocalNode
	if err = msgpack.Unmarshal(instruction.EncodedEPlanNodeBytes, &enode); err != nil {
		return err
	}
	self.Instruction = instruction
	self.EPlanNode = &enode
	self.InputLocations = []*pb.Location{&enode.Input}
	self.OutputLocations = []*pb.Location{&enode.Output}
	return nil
}

func (self *Executor) RunAggregateFuncLocal() (err error) {
	fname := fmt.Sprintf("executor_%v_aggregatefunclocal_%v_cpu.pprof", self.Name, time.Now().Format("20060102150405"))
	f, _ := os.Create(fname)
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	defer self.Clear()

	reader, writer := self.Readers[0], self.Writers[0]
	enode := self.EPlanNode.(*EPlan.EPlanAggregateFuncLocalNode)
	md := &Metadata.Metadata{}

	//read md
	if err = Util.ReadObject(reader, md); err != nil {
		return err
	}

	//write md
	if err = Util.WriteObject(writer, enode.Metadata); err != nil {
		return err
	}

	rbReader, rbWriter := Row.NewRowsBuffer(md, reader, nil), Row.NewRowsBuffer(enode.Metadata, nil, writer)

	defer func() {
		rbWriter.Flush()
	}()

	//init
	if err := enode.Init(enode.Metadata); err != nil {
		log.Println("=========agglocal", enode.Metadata, err)
		return err
	}

	//write rows
	var row *Row.Row
	var resRow, rowTmp *Row.Row
	var rg *Row.RowsGroup
	for {
		row, err = rbReader.ReadRow()
		log.Println("======", row, err)

		if err == io.EOF {
			err = nil
			if rg != nil && rg.GetRowsNum() > 0 {
				if rowTmp, err = self.CalAggregateFuncLocal(enode, rg); err != nil {
					break
				}
				resRow.AppendRow(rowTmp)
				rbWriter.WriteRow(resRow)
			}
			break
		}
		if err != nil {
			break
		}

		if rg == nil {
			rg = Row.NewRowsGroup(md)
			rg.Write(row)
			resRow = row

		} else {
			if rg.GetKeyString() == row.GetKeyString() {
				rg.Write(row)

			} else {
				if rowTmp, err = self.CalAggregateFuncLocal(enode, rg); err != nil {
					break
				}
				resRow.AppendRow(rowTmp)
				rbWriter.WriteRow(resRow)

				rg = Row.NewRowsGroup(md)
				rg.Write(row)
				resRow = row
				if err = enode.Init(md); err != nil {
					break
				}
			}

			if rg.GetRowsNum() > Row.ROWS_BUFFER_SIZE {
				if _, err = self.CalAggregateFuncLocal(enode, rg); err != nil {
					break
				}
				rg.ClearRows()
			}
		}
	}

	Logger.Infof("RunAggregateFuncLocal finished")
	return err
}

func (self *Executor) CalAggregateFuncLocal(enode *EPlan.EPlanAggregateFuncLocalNode, rg *Row.RowsGroup) (*Row.Row, error) {
	var err error
	var res interface{}
	row := Row.NewRow()
	for _, item := range enode.FuncNodes {
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
