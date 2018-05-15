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

func (self *Executor) SetInstructionJoin(instruction *pb.Instruction) (err error) {
	var enode EPlan.EPlanJoinNode
	if err = msgpack.Unmarshal(instruction.EncodedEPlanNodeBytes, &enode); err != nil {
		return err
	}
	self.Instruction = instruction
	self.EPlanNode = &enode
	self.InputLocations = []*pb.Location{&enode.LeftInput, &enode.RightInput}
	self.OutputLocations = []*pb.Location{&enode.Output}
	return nil
}

func (self *Executor) RunJoin() (err error) {
	defer self.Clear()
	writer := self.Writers[0]
	enode := self.EPlanNode.(*EPlan.EPlanJoinNode)

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

	md := enode.Metadata
	//write md
	if err = Util.WriteObject(writer, md); err != nil {
		return err
	}

	//write rows
	var row *Util.Row
	rows := make([]*Util.Row, 0)
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
			joinNum := 0
			for _, rightRow := range rows {
				joinRow := Util.NewRow(row.Vals...)
				joinRow.AppendRow(rightRow)
				rb := Util.NewRowsBuffer(md)
				rb.Write(joinRow)
				if ok, err := enode.JoinCriteria.Result(rb); ok && err == nil {
					if err = Util.WriteRow(writer, joinRow); err != nil {
						return err
					}
					joinNum++
				} else if err != nil {
					return err
				}
			}
			if enode.JoinType == Plan.LEFTJOIN && joinNum == 0 {
				joinRow := Util.NewRow(row.Vals...)
				joinRow.AppendVals(make([]interface{}, len(mds[1].ColumnNames))...)
				if err = Util.WriteRow(writer, joinRow); err != nil {
					return err
				}
			}

		}

	case Plan.RIGHTJOIN:
		for {
			row, err = Util.ReadRow(leftReader)
			if err == io.EOF {
				err = nil
				break
			}
			if err != nil {
				return err
			}
			rows = append(rows, row)
		}

		for {
			row, err = Util.ReadRow(rightReader)
			if err == io.EOF {
				err = nil
				break
			}
			if err != nil {
				return err
			}
			joinNum := 0
			for _, leftRow := range rows {
				joinRow := Util.NewRow(leftRow.Vals...)
				joinRow.AppendRow(row)
				rb := Util.NewRowsBuffer(md)
				rb.Write(joinRow)
				if ok, err := enode.JoinCriteria.Result(rb); ok && err == nil {
					if err = Util.WriteRow(writer, joinRow); err != nil {
						return err
					}
					joinNum++
				} else if err != nil {
					return err
				}
			}
			if joinNum == 0 {
				joinRow := Util.NewRow(make([]interface{}, len(mds[1].ColumnNames))...)
				joinRow.AppendVals(row.Vals...)
				if err = Util.WriteRow(writer, joinRow); err != nil {
					return err
				}
			}
		}
	}
	Util.WriteEOFMessage(writer)

	Logger.Infof("RunJoin finished")
	return err
}
