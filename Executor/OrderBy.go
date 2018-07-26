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
	"github.com/xitongsys/guery/Type"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/pb"
)

func (self *Executor) SetInstructionOrderBy(instruction *pb.Instruction) (err error) {
	var enode EPlan.EPlanOrderByNode
	if err = msgpack.Unmarshal(instruction.EncodedEPlanNodeBytes, &enode); err != nil {
		return err
	}
	self.Instruction = instruction
	self.EPlanNode = &enode
	self.InputLocations = []*pb.Location{}
	for i := 0; i < len(enode.Inputs); i++ {
		loc := enode.Inputs[i]
		self.InputLocations = append(self.InputLocations, &loc)
	}
	self.OutputLocations = []*pb.Location{&enode.Output}
	return nil
}

func (self *Executor) RunOrderBy() (err error) {
	fname := fmt.Sprintf("executor_%v_orderby_%v_cpu.pprof", self.Name, time.Now().Format("20060102150405"))
	f, _ := os.Create(fname)
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	defer func() {
		if err != nil {
			self.AddLogInfo(err, pb.LogLevel_ERR)
		}
		self.Clear()
	}()

	enode := self.EPlanNode.(*EPlan.EPlanOrderByNode)
	md := &Metadata.Metadata{}
	//read md
	for _, reader := range self.Readers {
		if err = Util.ReadObject(reader, md); err != nil {
			return err
		}
	}

	//write md
	writer := self.Writers[0]
	if err = Util.WriteObject(writer, enode.Metadata); err != nil {
		return err
	}

	rbReaders := make([]*Row.RowsBuffer, len(self.Readers))
	for i, reader := range self.Readers {
		rbReaders[i] = Row.NewRowsBuffer(md, reader, nil)
	}
	rbWriter := Row.NewRowsBuffer(enode.Metadata, nil, writer)

	defer func() {
		rbWriter.Flush()
	}()

	//write rows
	var row *Row.Row
	rows := Row.NewRows(self.GetOrder(enode))
	rows.Data = make([]*Row.Row, len(self.Readers))

	isEnd := make([]bool, len(self.Readers))
	for {
		for i := 0; i < len(isEnd); i++ {
			if !isEnd[i] && rows.Data[i] == nil {
				row, err = rbReaders[i].ReadRow()
				if err == io.EOF {
					err = nil
					isEnd[i] = true
					continue
				}
				if err != nil {
					return err
				}
				rows.Data[i] = row
			}
		}

		if minIndex := rows.Min(); minIndex < 0 {
			break

		} else {
			rows.Data[minIndex].ClearKeys()
			if err = rbWriter.WriteRow(rows.Data[minIndex]); err != nil {
				return err
			}
			rows.Data[minIndex] = nil
		}
	}

	Logger.Infof("RunOrderBy finished")
	return nil
}

func (self *Executor) GetOrder(enode *EPlan.EPlanOrderByNode) []Type.OrderType {
	res := []Type.OrderType{}
	for _, item := range enode.SortItems {
		res = append(res, item.OrderType)
	}
	return res
}
