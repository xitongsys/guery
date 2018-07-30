package executor

import (
	"fmt"
	"io"
	"os"
	"runtime/pprof"
	"time"

	"github.com/vmihailenco/msgpack"
	"github.com/xitongsys/guery/eplan"
	"github.com/xitongsys/guery/gtype"
	"github.com/xitongsys/guery/logger"
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/pb"
	"github.com/xitongsys/guery/row"
	"github.com/xitongsys/guery/util"
)

func (self *Executor) SetInstructionSelect(instruction *pb.Instruction) (err error) {
	var enode eplan.EPlanSelectNode
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

	defer func() {
		if err != nil {
			self.AddLogInfo(err, pb.LogLevel_ERR)
		}
		self.Clear()
	}()

	if self.Instruction == nil {
		return fmt.Errorf("No Instruction")
	}
	enode := self.EPlanNode.(*eplan.EPlanSelectNode)

	md := &metadata.Metadata{}
	reader := self.Readers[0]
	writer := self.Writers[0]
	if err = util.ReadObject(reader, md); err != nil {
		return err
	}

	//write metadata
	if err = util.WriteObject(writer, enode.Metadata); err != nil {
		return err
	}

	rbReader, rbWriter := row.NewRowsBuffer(md, reader, nil), row.NewRowsBuffer(enode.Metadata, nil, writer)
	defer func() {
		rbWriter.Flush()
	}()

	//init
	for _, item := range enode.SelectItems {
		if err := item.Init(md); err != nil {
			return err
		}
	}
	distinctMap := make(map[string]bool)

	//write rows
	var rg, res *row.RowsGroup
	for {
		rg, err = rbReader.Read()
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			break
		}

		if res, err = self.CalSelectItems(enode, rg); err != nil {
			break
		}

		//for distinct
		if enode.SetQuantifier != nil && (*enode.SetQuantifier) == gtype.DISTINCT {
			for i := 0; i < res.GetRowsNumber(); i++ {
				row := res.GetRow(i)
				rowkey := fmt.Sprintf("%v", row)
				if _, ok := distinctMap[rowkey]; ok {
					continue
				}
				distinctMap[rowkey] = true
				if err = rbWriter.WriteRow(row); err != nil {
					break
				}
			}

		} else {
			if err = rbWriter.Write(res); err != nil {
				logger.Errorf("failed to Write %v", err)
				break
			}
		}
	}

	logger.Infof("RunSelect finished")
	return err
}

func (self *Executor) CalSelectItems(enode *eplan.EPlanSelectNode, rg *row.RowsGroup) (*row.RowsGroup, error) {
	var err error
	var vs []interface{}
	res := row.NewRowsGroup(enode.Metadata)
	ci := 0

	if enode.Having != nil {
		rgtmp := row.NewRowsGroup(rg.Metadata)
		flags, err := enode.Having.Result(rg)
		if err != nil {
			return nil, err
		}

		for i, flag := range flags.([]interface{}) {
			if flag.(bool) {
				rgtmp.AppendRowVals(rg.GetRowVals(i)...)
				rgtmp.AppendRowKeys(rg.GetRowKeys(i)...)
			}
		}
		rg = rgtmp
	}

	for _, item := range enode.SelectItems {
		vs, err = item.Result(rg)
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}

		if item.Expression == nil { //*
			for _, vi := range vs {
				res.Vals[ci] = append(res.Vals[ci], vi.([]interface{})...)
				ci++
			}

		} else {
			res.Vals[ci] = append(res.Vals[ci], vs[0].([]interface{})...)
			ci++
		}

		res.RowsNumber = len(vs)
	}

	return res, err
}
