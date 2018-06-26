package Executor

import (
	"fmt"
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

func (self *Executor) SetInstructionGroupBy(instruction *pb.Instruction) (err error) {
	var enode EPlan.EPlanGroupByNode
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
	self.OutputLocations = []*pb.Location{}
	for i := 0; i < len(enode.Outputs); i++ {
		loc := enode.Outputs[i]
		self.OutputLocations = append(self.OutputLocations, &loc)
	}
	return nil
}

func (self *Executor) RunGroupBy() (err error) {
	Logger.Infof("RunGroupBy")
	defer self.Clear()

	if self.Instruction == nil {
		return fmt.Errorf("no instruction")
	}
	enode := self.EPlanNode.(*EPlan.EPlanGroupByNode)

	mds := make([]*Metadata.Metadata, len(self.Readers))
	for i, reader := range self.Readers {
		mds[i] = &Metadata.Metadata{}
		if err = Util.ReadObject(reader, mds[i]); err != nil {
			return err
		}
	}

	//write metadata
	enode.Metadata.ClearKeys()
	enode.Metadata.AppendKeyByType(Type.STRING)
	for _, writer := range self.Writers {
		if err = Util.WriteObject(writer, enode.Metadata); err != nil {
			return err
		}
	}
	rbWriters := make([]*Split.SplitBuffer, len(self.Writers))
	for i, writer := range self.Writers {
		rbWriters[i] = Split.NewSplitBuffer(enode.Metadata, nil, writer)
	}

	defer func() {
		for _, rbWriters := range rbWriters {
			rbWriters.Flush()
		}
	}()

	//group by
	var sp *Split.Split
	var distMap = make(map[string]int)
	j, ln := 0, len(self.Writers)
	for i, reader := range self.Readers {
		rbReader := Split.NewSplitBuffer(mds[i], reader, nil)
		for {
			sp, err = rbReader.ReadSplit()
			if err != nil {
				if err == io.EOF {
					err = nil
				}
				break
			}

			keys := make([]interface{}, sp.GetRowsNumber())
			for i := 0; i < sp.GetRowsNumber(); i++ {
				keys[i], err = self.CalGroupByKey(enode, mds[i], sp, i)
				if err != nil {
					return err
				}
			}
			sp.AppendKeyColumns(keys)

			for i := 0; i < sp.GetRowsNumber(); i++ {
				key := keys[i]
				if k, ok = distMap[key]; !ok {
					distMap[key] = j
					k = j
					j = (j + 1) % ln
				}

				if err := rbWriters[k].Write(sp, i); err != nil {
					return err
				}
			}
		}
	}
	return err
}

func (self *Executor) CalGroupByKey(enode *EPlan.EPlanGroupByNode, md *Metadata.Metadata, sp *Split.Split, index int) (string, error) {
	res, err := enode.GroupBy.Result(sp, index)
	if err != nil {
		return res, err
	}
	return res, nil
}
