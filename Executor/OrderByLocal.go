package Executor

import (
	"io"

	"github.com/vmihailenco/msgpack"
	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/pb"
)

func (self *Executor) SetInstructionOrderByLocal(instruction *pb.Instruction) (err error) {
	var enode EPlan.EPlanOrderByLocalNode
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
	defer self.Clear()

	reader, writer := self.Readers[0], self.Writers[0]
	enode := self.EPlanNode.(*EPlan.EPlanOrderByLocalNode)
	md := &Util.Metadata{}

	//read md
	if err = Util.ReadObject(reader, md); err != nil {
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
	if err = Util.WriteObject(writer, enode.Metadata); err != nil {
		return err
	}

	rbReader, rbWriter := Util.NewRowsBuffer(md, reader, nil), Util.NewRowsBuffer(enode.Metadata, nil, writer)

	defer func() {
		rbWriter.Flush()
	}()

	//write rows
	var row *Util.Row
	rows := Util.NewRows(self.GetOrderLocal(enode))

	for {
		row, err = rbReader.ReadRow()
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			return err
		}
		rg := Util.NewRowsGroup(md)
		rg.Write(row)
		row.Keys, err = self.CalSortKey(enode, rg)
		if err != nil {
			return err
		}
		rows.Append(row)
	}
	rows.Sort()
	for _, row := range rows.Data {
		if err = rbWriter.WriteRow(row); err != nil {
			return err
		}
	}

	Logger.Infof("RunOrderByLocal finished")
	return nil
}

func (self *Executor) GetOrderLocal(enode *EPlan.EPlanOrderByLocalNode) []Util.OrderType {
	res := []Util.OrderType{}
	for _, item := range enode.SortItems {
		res = append(res, item.OrderType)
	}
	return res
}

func (self *Executor) CalSortKey(enode *EPlan.EPlanOrderByLocalNode, rg *Util.RowsGroup) ([]interface{}, error) {
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
