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
		return err
	}

	//write rows
	var row *Row.Row
	rg := Row.NewRowsGroup(md)
	var res []map[string]interface{}
	keys := map[string]*Row.Row{}
	for {
		row, err = rbReader.ReadRow()

		if err == io.EOF {
			err = nil
			if res, err = self.CalAggregateFuncLocal(enode, rg); err != nil {
				break
			}
			if len(res) <= 0 {
				break
			}
			for key, row := range keys {
				rowg := Row.NewRow()
				for i := 0; i < len(res); i++ {
					rowg.AppendVals(res[i][key])
				}
				row.AppendRow(rowg)
				rbWriter.WriteRow(row)
			}

			break
		}
		if err != nil {
			break
		}

		key := row.GetKeyString()
		keys[key] = row

		rg.Write(row)
		if rg.GetRowsNum() > Row.ROWS_BUFFER_SIZE {
			if _, err = self.CalAggregateFuncLocal(enode, rg); err != nil {
				break
			}
			rg.ClearRows()
		}
	}

	Logger.Infof("RunAggregateFuncLocal finished")
	return err
}

func (self *Executor) CalAggregateFuncLocal(enode *EPlan.EPlanAggregateFuncLocalNode, rg *Row.RowsGroup) ([]map[string]interface{}, error) {
	var err error
	var res []map[string]interface{}
	var resc map[string]interface{}
	var resci interface{}
	for _, item := range enode.FuncNodes {
		rg.Reset()
		resci, err = item.Result(rg)
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}
		resc = resci.(map[string]interface{})
		res = append(res, resc)
	}
	return res, err
}
