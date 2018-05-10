package Executor

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/pb"
)

func (self *Executor) SetInstructionGroupBy(instruction *pb.Instruction) (err error) {
	var enode EPlan.EPlanGroupByNode
	if err = gob.NewDecoder(bytes.NewBuffer(instruction.EncodedEPlanNodeBytes)).Decode(&enode); err != nil {
		return err
	}
	return nil
}

func (self *Executor) RunGroupBy() (err error) {
	defer self.Clear()

	if self.Instruction == nil {
		return fmt.Errorf("no instruction")
	}
	enode := self.EPlanNode.(*EPlan.EPlanGroupByNode)

	mds := make([]*Util.Metadata, len(self.Readers))
	for i, reader := range self.Readers {
		mds[i] = &Util.Metadata{}
		if err = Util.ReadObject(reader, mds[i]); err != nil {
			return err
		}
	}

	//write metadata
	gmd := enode.Metadata
	for _, writer := range self.Writers {
		if err = Util.WriteObject(writer, gmd); err != nil {
			return err
		}
	}

	//group by
	var row *Util.Row
	var rowsBufs = make(map[string]*Util.RowsBuffer)
	for i, reader := range self.Readers {
		for {
			row, err = Util.ReadRow(reader)
			if err != nil {
				if err == io.EOF {
					err = nil
				}
				break
			}

			row.Key, err = self.CalGroupByKey(enode, mds[i], row)
			if err != nil {
				return err
			}

			if _, ok := rowsBufs[key]; !ok {
				rowsBufs[key] = Util.NewRowsBuffer(gmd)
			}
			rowsBufs[key].Write(row)
		}
	}

	//write rows

	return nil
}

func (self *Executor) CalGroupByKey(enode *EPlan.EPlanGroupByNode, md *Util.Metadata, row *Util.Row) (string, error) {
	rowsBuf := Util.NewRowsBuffer(md)
	var res string
	for _, item := range enode.GroupBy {
		rowsBuf.Reset()
		r, err := item.Result(rowsBuf)
		if err != nil {
			return res, err
		}
		res += r
	}
	return res

}
