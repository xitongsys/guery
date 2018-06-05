package Executor

import (
	"fmt"
	"io"

	"github.com/vmihailenco/msgpack"
	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/pb"
)

func (self *Executor) SetInstructionHashJoin(instruction *pb.Instruction) (err error) {
	var enode EPlan.EPlanHashJoinNode
	if err = msgpack.Unmarshal(instruction.EncodedEPlanNodeBytes, &enode); err != nil {
		return err
	}
	self.Instruction = instruction
	self.EPlanNode = &enode
	self.InputLocations = []*pb.Location{&enode.LeftInput, &enode.RightInput}
	self.OutputLocations = []*pb.Location{&enode.Output}
	return nil
}

func CalHashKey(es []*Plan.ValueExpressionNode, rg *Util.RowsGroup) (string, error) {
	res := ""
	for _, e := range es {
		r, err := e.Result(rg)
		if err != nil {
			return res, err
		}
		res += fmt.Sprintf("_%v", r)
	}
	return res, nil
}

func (self *Executor) RunHashJoin() (err error) {
	defer self.Clear()
	writer := self.Writers[0]
	enode := self.EPlanNode.(*EPlan.EPlanHashJoinNode)

	//read md
	if len(self.Readers) != 2 {
		return fmt.Errorf("join readers number %v <> 2", len(self.Readers))
	}

	mds := make([]*Util.Metadata, 2)
	if len(self.Readers) != 2 {
		return fmt.Errorf("join input number error")
	}
	for i, reader := range self.Readers {
		mds[i] = &Util.Metadata{}
		if err = Util.ReadObject(reader, mds[i]); err != nil {
			return err
		}
	}
	leftReader, rightReader := self.Readers[0], self.Readers[1]
	leftMd, _ := mds[0], mds[1]

	//write md
	if err = Util.WriteObject(writer, enode.Metadata); err != nil {
		return err
	}

	//write rows
	var row *Util.Row
	rows := make([]*Util.Row, 0)
	rowsMap := make(map[string][]int)

	switch enode.JoinType {
	case Plan.INNERJOIN:
		fallthrough
	case Plan.LEFTJOIN:
		for {
			row, err = Util.ReadRow(rightReader)
			if err == io.EOF {
				err = nil
				break
			}
			if err != nil {
				return err
			}
			rows = append(rows, row)
			key := row.GetKeyString()

			if v, ok := rowsMap[key]; ok {
				v = append(v, len(rows)-1)
			} else {
				rowsMap[key] = []int{len(rows) - 1}
			}
		}

		for {
			row, err = Util.ReadRow(leftReader)
			if err == io.EOF {
				err = nil
				break
			}
			if err != nil {
				return err
			}
			rg := Util.NewRowsGroup(leftMd)
			rg.Write(row)
			leftKey, err := CalHashKey(enode.LeftKeys, rg)
			if err != nil {
				return err
			}

			joinNum := 0
			if _, ok := rowsMap[leftKey]; ok {
				for _, i := range rowsMap[leftKey] {
					rightRow := rows[i]
					joinRow := Util.NewRow(row.Vals...)
					joinRow.AppendVals(rightRow.Vals...)
					rg := Util.NewRowsGroup(enode.Metadata)
					rg.Write(joinRow)

					if ok, err := enode.JoinCriteria.Result(rg); ok && err == nil {
						if err = Util.WriteRow(writer, joinRow); err != nil {
							return err
						}
						joinNum++
					} else if err != nil {
						return err
					}
				}
			}

			if enode.JoinType == Plan.LEFTJOIN && joinNum == 0 {
				joinRow := Util.NewRow(row.Vals...)
				joinRow.AppendVals(make([]interface{}, len(mds[1].GetColumnNames()))...)
				if err = Util.WriteRow(writer, joinRow); err != nil {
					return err
				}
			}

		}

	case Plan.RIGHTJOIN:

	}
	Util.WriteEOFMessage(writer)

	Logger.Infof("RunJoin finished")
	return err
}
