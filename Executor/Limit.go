package Executor

import (
	"io"

	"github.com/vmihailenco/msgpack"
	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Split"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/pb"
)

func (self *Executor) SetInstructionLimit(instruction *pb.Instruction) (err error) {
	var enode EPlan.EPlanLimitNode
	if err = msgpack.Unmarshal(instruction.EncodedEPlanNodeBytes, &enode); err != nil {
		return err
	}
	self.Instruction = instruction
	self.EPlanNode = &enode
	self.InputLocations = []*pb.Location{&enode.Input}
	self.OutputLocations = []*pb.Location{&enode.Output}
	return nil
}

func (self *Executor) RunLimit() (err error) {
	defer self.Clear()

	enode := self.EPlanNode.(*EPlan.EPlanLimitNode)
	writer := self.Writers[0]
	md := &Metadata.Metadata{}
	//read md
	for _, reader := range self.Readers {
		if err = Util.ReadObject(reader, md); err != nil {
			return err
		}
	}

	//write md
	if err = Util.WriteObject(writer, md); err != nil {
		return err
	}

	rbReaders := make([]*Split.SplitBuffer, len(self.Readers))
	for i, reader := range self.Readers {
		rbReaders[i] = Split.NewSplitBuffer(md, reader, nil)
	}
	rbWriter := Split.NewSplitBuffer(md, nil, writer)

	defer func() {
		rbWriter.Flush()
	}()

	//write rows
	var sp *Split.Split
	readRowCnt := 0
	for _, rbReader := range rbReaders {
		for readRowCnt < *(enode.LimitNumber) {
			sp, err = rbReader.ReadSplit()
			if err == io.EOF || readRowCnt >= int(*(enode.LimitNumber)) {
				err = nil
				break
			}
			if err != nil {
				return err
			}
			readRowCnt += sp.GetRowsNumber()
			if err = rbWriter.FlushSplit(sp); err != nil {
				return err
			}
		}
	}

	Logger.Infof("RunAggregate finished")
	return nil
}
