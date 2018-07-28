package executor

import (
	"fmt"
	"io"
	"os"
	"runtime/pprof"
	"time"

	"github.com/vmihailenco/msgpack"
	"github.com/xitongsys/guery/eplan"
	"github.com/xitongsys/guery/logger"
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/pb"
	"github.com/xitongsys/guery/plan"
	"github.com/xitongsys/guery/row"
	"github.com/xitongsys/guery/util"
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
	fname := fmt.Sprintf("executor_%v_join_%v_cpu.pprof", self.Name, time.Now().Format("20060102150405"))
	f, _ := os.Create(fname)
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	defer func() {
		if err != nil {
			self.AddLogInfo(err, pb.LogLevel_ERR)
		}
		self.Clear()
	}()
	writer := self.Writers[0]
	enode := self.EPlanNode.(*EPlan.EPlanJoinNode)

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

	leftRbReader, rightRbReader := Row.NewRowsBuffer(leftMd, leftReader, nil), Row.NewRowsBuffer(rightMd, rightReader, nil)
	rbWriter := Row.NewRowsBuffer(enode.Metadata, nil, writer)

	defer func() {
		rbWriter.Flush()
	}()

	//init
	if err := enode.JoinCriteria.Init(enode.Metadata); err != nil {
		return err
	}

	//write rows
	var row *Row.Row
	rows := make([]*Row.Row, 0)
	switch enode.JoinType {
	case Plan.INNERJOIN:
		fallthrough
	case Plan.LEFTJOIN:
		for {
			row, err = rightRbReader.ReadRow()
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
			row, err = leftRbReader.ReadRow()
			if err == io.EOF {
				err = nil
				break
			}
			if err != nil {
				return err
			}
			joinNum := 0
			for _, rightRow := range rows {
				joinRow := Row.NewRow(row.Vals...)
				joinRow.AppendVals(rightRow.Vals...)
				rg := Row.NewRowsGroup(enode.Metadata)
				rg.Write(joinRow)
				if ok, err := enode.JoinCriteria.Result(rg); ok && err == nil {
					if err = rbWriter.WriteRow(joinRow); err != nil {
						return err
					}
					joinNum++
				} else if err != nil {
					return err
				}
			}
			if enode.JoinType == Plan.LEFTJOIN && joinNum == 0 {
				joinRow := Row.NewRow(row.Vals...)
				joinRow.AppendVals(make([]interface{}, len(mds[1].GetColumnNames()))...)
				if err = rbWriter.WriteRow(joinRow); err != nil {
					return err
				}
			}
		}

	case Plan.RIGHTJOIN:
	}

	Logger.Infof("RunJoin finished")
	return err
}
