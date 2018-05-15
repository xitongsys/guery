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
	if err = Util.WriteObject(writer, enode.Metadata); err != nil {
		return err
	}

	//write rows
	var row *Util.Row
	rows := Util.NewRows(self.GetOrderLocal(enode))

	for {
		row, err = Util.ReadRow(reader)
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			return err
		}
		rb := Util.NewRowsBuffer(md)
		rb.Write(row)
		row.Keys, err = self.CalSortKey(enode, rb)
		if err != nil {
			return err
		}
		rows.Append(row)
	}
	rows.Sort()
	for _, row := range rows.Data {
		if err = Util.WriteRow(writer, row); err != nil {
			return err
		}
	}

	Util.WriteEOFMessage(writer)
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

func (self *Executor) CalSortKey(enode *EPlan.EPlanOrderByLocalNode, rowsBuf *Util.RowsBuffer) ([]interface{}, error) {
	var err error
	res := []interface{}{}
	for _, item := range enode.SortItems {
		key, err := item.Result(rowsBuf)
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
