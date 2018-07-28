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

func (self *Executor) SetInstructionOrderByLocal(instruction *pb.Instruction) (err error) {
	var enode eplan.EPlanOrderByLocalNode
	if err = msgpack.Unmarshal(instruction.EncodedEPlanNodeBytes, &enode); err != nil {
		return err
	}
	self.Instruction = instruction
	self.EPlanNode = &enode
	self.InputLocations = []*pb.Location{&enode.Input}
	self.OutputLocations = []*pb.Location{&enode.Output}
	return nil
}

func (self *Executor) RunOrderByLocal() (err error) {
	fname := fmt.Sprintf("executor_%v_orderbylocal_%v_cpu.pprof", self.Name, time.Now().Format("20060102150405"))
	f, _ := os.Create(fname)
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	defer func() {
		if err != nil {
			self.AddLogInfo(err, pb.LogLevel_ERR)
		}
		self.Clear()
	}()

	reader, writer := self.Readers[0], self.Writers[0]
	enode := self.EPlanNode.(*EPlan.EPlanOrderByLocalNode)
	md := &metadata.Metadata{}

	//read md
	if err = util.ReadObject(reader, md); err != nil {
		return err
	}

	//write md
	enode.Metadata.ClearKeys()
	for _, item := range enode.SortItems {
		t, err := item.GetType(md)
		if err != nil {
			return err
		}
		enode.Metadata.AppendKeyByType(t)
	}
	if err = util.WriteObject(writer, enode.Metadata); err != nil {
		return err
	}

	rbReader, rbWriter := row.NewRowsBuffer(md, reader, nil), row.NewRowsBuffer(enode.Metadata, nil, writer)

	defer func() {
		rbWriter.Flush()
	}()

	//init
	for _, item := range enode.SortItems {
		if err := item.Init(md); err != nil {
			return err
		}
	}

	//write rows
	var r *row.Row
	rs := row.NewRows(self.GetOrderLocal(enode))

	for {
		r, err = rbReader.ReadRow()
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			return err
		}
		rg := row.NewRowsGroup(md)
		rg.Write(r)
		r.Keys, err = self.CalSortKey(enode, rg)
		if err != nil {
			return err
		}
		rs.Append(r)
	}
	rs.Sort()
	for _, r := range rs.Data {
		if err = rbWriter.WriteRow(r); err != nil {
			return err
		}
	}

	logger.Infof("RunOrderByLocal finished")
	return nil
}

func (self *Executor) GetOrderLocal(enode *EPlan.EPlanOrderByLocalNode) []gtype.OrderType {
	res := []gtype.OrderType{}
	for _, item := range enode.SortItems {
		res = append(res, item.OrderType)
	}
	return res
}

func (self *Executor) CalSortKey(enode *eplan.EPlanOrderByLocalNode, rg *row.RowsGroup) ([]interface{}, error) {
	var err error
	res := []interface{}{}
	for _, item := range enode.SortItems {
		key, err := item.Result(rg)
		if err == io.EOF {
			return res, nil
		}
		if err != nil {
			return res, err
		}
		res = append(res, key)
	}

	return res, err

}
