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

	defer func() {
		if err != nil {
			self.AddLogInfo(err, pb.LogLevel_ERR)
		}
		self.Clear()
	}()

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
	var rg, res *Row.RowsGroup
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

		if err = rbWriter.Write(res); err != nil {
			Logger.Errorf("failed to Write %v", err)
			break
		}
	}

	Logger.Infof("RunSelect finished")
	return err
}

func (self *Executor) CalSelectItems(enode *EPlan.EPlanSelectNode, rg *Row.RowsGroup) (*Row.RowsGroup, error) {
	var err error
	var vs []interface{}
	res := Row.NewRowsGroup(enode.Metadata)
	ci := 0

	if enode.Having != nil {
		rgtmp := Row.NewRowsGroup(rg.Metadata)
		flags, err := enode.Having.Result(rg)
		if err != nil {
			return nil, err
		}

		for i, flag := range flags.([]interface{}) {
			if flag.(bool) {
				rgtmp.AppendValRow(rg.GetRowVals(i)...)
				rgtmp.AppendKeyRow(rg.GetRowKeys(i)...)
				rgtmp.RowsNumber++
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
