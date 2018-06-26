package Executor

import (
	"fmt"
	"io"

	"github.com/vmihailenco/msgpack"
	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/Split"
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

func CalHashKey(es []*Plan.ValueExpressionNode, sp *Split.Split, index int) (string, error) {
	res := ""
	for _, e := range es {
		r, err := e.Result(sp, index)
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

	mds := make([]*Metadata.Metadata, 2)
	if len(self.Readers) != 2 {
		return fmt.Errorf("join input number error")
	}
	for i, reader := range self.Readers {
		mds[i] = &Metadata.Metadata{}
		if err = Util.ReadObject(reader, mds[i]); err != nil {
			return err
		}
	}
	leftReader, rightReader := self.Readers[0], self.Readers[1]
	leftMd, rightMd := mds[0], mds[1]

	//write md
	if err = Util.WriteObject(writer, enode.Metadata); err != nil {
		return err
	}

	leftRbReader, rightRbReader := Split.NewSplitBuffer(leftMd, leftReader, nil), Split.NewSplitBuffer(rightMd, rightReader, nil)
	rbWriter := Split.NewSplitBuffer(enode.Metadata, nil, writer)

	defer func() {
		rbWriter.Flush()
	}()

	//write rows
	var sp *Split.Split
	rightSp := Split.NewSplit(rightMd)
	rowsMap := make(map[string][]int)

	switch enode.JoinType {
	case Plan.INNERJOIN:
		fallthrough
	case Plan.LEFTJOIN:
		for {
			sp, err = rightRbReader.ReadSplit()
			if err == io.EOF {
				err = nil
				break
			}
			if err != nil {
				return err
			}

			for i := 0; i < sp.GetRowsNumber(); i++ {
				key := sp.GetKeyString(i) //key calculated in duplicate

				if v, ok := rowsMap[key]; ok {
					v = append(v, i+rightSp.GetRowsNum())
				} else {
					rowsMap[key] = []int{i + rightSp.GetRowsNumber()}
				}
			}
			rightSp.Append(sp)
		}

		for {
			sp, err = leftRbReader.ReadSplit()
			if err == io.EOF {
				err = nil
				break
			}
			if err != nil {
				return err
			}

			joinNum := 0
			for i := 0; i < sp.GetRowsNumber(); i++ {
				leftKey = CalHashKey(enode.LeftKeys, sp, i)

				if _, ok := rowsMap[leftKey]; ok {
					for _, j := range rowsMap[leftKey] {
						joinSp := Split.NewSplit(enode.Metadata)
						vals := sp.GetValues(i)
						vals = append(vals, rightSp.GetValues(j)...)
						joinSp.AppendValues(vals)

						if ok, err := enode.JoinCriteria.Result(joinSp, 0); ok && err == nil {
							if err = rbWriter.Write(joinSp, 0); err != nil {
								return err
							}
							joinNum++
						} else if err != nil {
							return err
						}
					}
				}

				if enode.JoinType == Plan.LEFTJOIN && joinNum == 0 {
					joinSp := Split.NewSplit(enode.Metadata)
					vals := sp.GetValues(i)
					vals = append(vals, make([]interface{}, rightSp.GetColumnNumber())...)
					joinSp.AppendValues(vals)

					if err = rbWriter.Write(joinSp, 0); err != nil {
						return err
					}
				}
			}

		}

	case Plan.RIGHTJOIN:

	}

	Logger.Infof("RunJoin finished")
	return err
}
