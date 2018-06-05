package Executor

import (
	"io"

	"github.com/vmihailenco/msgpack"
	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Logger"
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
	defer self.Clear()

	enode := self.EPlanNode.(*EPlan.EPlanOrderByNode)
	md := &Util.Metadata{}
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

	//write rows
	var row *Util.Row
	rows := Util.NewRows(self.GetOrder(enode))
	rows.Data = make([]*Util.Row, len(self.Readers))

	isEnd := make([]bool, len(self.Readers))
	for {
		for i := 0; i < len(isEnd); i++ {
			if !isEnd[i] && rows.Data[i] == nil {
				row, err = Util.ReadRow(self.Readers[i])
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
			if err = Util.WriteRow(writer, rows.Data[minIndex]); err != nil {
				return err
			}
			rows.Data[minIndex] = nil
		}
	}

	Util.WriteEOFMessage(writer)
	Logger.Infof("RunOrderBy finished")
	return nil
}

func (self *Executor) GetOrder(enode *EPlan.EPlanOrderByNode) []Util.OrderType {
	res := []Util.OrderType{}
	for _, item := range enode.SortItems {
		res = append(res, item.OrderType)
	}
	return res
}
