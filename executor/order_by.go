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

func (self *Executor) SetInstructionOrderBy(instruction *pb.Instruction) (err error) {
	var enode eplan.EPlanOrderByNode
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

	enode := self.EPlanNode.(*eplan.EPlanOrderByNode)
	md := &metadata.Metadata{}
	//read md
	for _, reader := range self.Readers {
		if err = util.ReadObject(reader, md); err != nil {
			return err
		}
	}

	//write md
	writer := self.Writers[0]
	if err = util.WriteObject(writer, enode.Metadata); err != nil {
		return err
	}

	rbReaders := make([]*row.RowsBuffer, len(self.Readers))
	for i, reader := range self.Readers {
		rbReaders[i] = row.NewRowsBuffer(md, reader, nil)
	}
	rbWriter := row.NewRowsBuffer(enode.Metadata, nil, writer)

	defer func() {
		rbWriter.Flush()
	}()

	//write rows
	var r *row.Row
	rs := row.NewRows(self.GetOrder(enode))
	rs.Data = make([]*row.Row, len(self.Readers))

	isEnd := make([]bool, len(self.Readers))
	for {
		for i := 0; i < len(isEnd); i++ {
			if !isEnd[i] && rs.Data[i] == nil {
				r, err = rbReaders[i].ReadRow()
				if err == io.EOF {
					err = nil
					isEnd[i] = true
					continue
				}
				if err != nil {
					return err
				}
				rs.Data[i] = r
			}
		}

		if minIndex := rs.Min(); minIndex < 0 {
			break

		} else {
			rs.Data[minIndex].ClearKeys()
			if err = rbWriter.WriteRow(rs.Data[minIndex]); err != nil {
				return err
			}
			rs.Data[minIndex] = nil
		}
	}

	logger.Infof("RunOrderBy finished")
	return nil
}

func (self *Executor) GetOrder(enode *eplan.EPlanOrderByNode) []gtype.OrderType {
	res := []gtype.OrderType{}
	for _, item := range enode.SortItems {
		res = append(res, item.OrderType)
	}
	return res
}
