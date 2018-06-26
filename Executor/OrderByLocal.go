package Executor

import (
	"io"

	"github.com/vmihailenco/msgpack"
	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Split"
	"github.com/xitongsys/guery/Type"
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
	md := &Metadata.Metadata{}

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

	rbReader, rbWriter := Split.NewSplitBuffer(md, reader, nil), Split.NewSplitBuffer(enode.Metadata, nil, writer)

	defer func() {
		rbWriter.Flush()
	}()

	//write
	var sp *Split.Split
	spOrder := Split.NewSplit(enode.Metadata)
	spOrder.OrderTypes = self.GetOrderLocal(enode)

	for {
		sp, err = rbReader.ReadSplit()
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			return err
		}

		spOrder.Append(sp)

		for i := 0; i < sp.GetRowsNumber(); i++ {
			keys, err = self.CalSortKey(enode, sp, i)
			if err != nil {
				return err
			}
			spOrder.AppendKeyValues(keys)
		}
	}
	spOrder.Sort()

	for _, row := range rows.Data {
		if err = rbWriter.WriteRow(row); err != nil {
			return err
		}
	}

	Logger.Infof("RunOrderByLocal finished")
	return nil
}

func (self *Executor) GetOrderLocal(enode *EPlan.EPlanOrderByLocalNode) []Type.OrderType {
	res := []Type.OrderType{}
	for _, item := range enode.SortItems {
		res = append(res, item.OrderType)
	}
	return res
}

func (self *Executor) CalSortKey(enode *EPlan.EPlanOrderByLocalNode, sp *Split.Split, index int) ([]interface{}, error) {
	var err error
	res := []interface{}{}
	for _, item := range enode.SortItems {
		key, err := item.Result(sp, index)
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
