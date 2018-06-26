package Executor

import (
	"fmt"
	"io"

	"github.com/vmihailenco/msgpack"
	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Split"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/pb"
)

func (self *Executor) SetInstructionGroupByLocal(instruction *pb.Instruction) (err error) {
	var enode EPlan.EPlanGroupByLocalNode
	if err = msgpack.Unmarshal(instruction.EncodedEPlanNodeBytes, &enode); err != nil {
		return err
	}
	self.Instruction = instruction
	self.EPlanNode = &enode
	self.InputLocations = []*pb.Location{&enode.Input}
	self.OutputLocations = []*pb.Location{&enode.Output}
	return nil
}

func (self *Executor) RunGroupByLocal() (err error) {
	Logger.Infof("RunGroupByLocal")
	defer self.Clear()

	if self.Instruction == nil {
		return fmt.Errorf("no instruction")
	}
	enode := self.EPlanNode.(*EPlan.EPlanGroupByLocalNode)

	reader := self.Readers[0]
	writer := self.Writers[0]
	md := &Metadata.Metadata{}

	if err = Util.ReadObject(reader, md); err != nil {
		return err
	}

	//write metadata
	if err = Util.WriteObject(writer, md); err != nil {
		return err
	}

	rbReader := Split.NewSplitBuffer(md, reader, nil)
	rbWriter := Split.NewSplitBuffer(md, nil, writer)

	//group by
	var sp *Split.Split
	var rgs = make(map[string]*Split.Split)
	for {
		sp, err = rbReader.ReadSplit()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}

		for i := 0; i < sp.GetRowsNumber(); i++ {
			key := sp.GetKeyString(i)
			if _, ok := rgs[key]; !ok {
				rgs[key] = Split.NewSplit(enode.Metadata)
			}
			rgs[key].Append(sp, i)
		}
	}

	defer func() {
		rbWriter.Flush()
	}()

	//write
	for _, rg := range rgs {
		var (
			ok  interface{} = true
			err error       = nil
		)
		if enode.GroupBy.Having != nil {
			ok, err = enode.GroupBy.Having.Result(rg, 0)
		}

		if err == nil && ok.(bool) {
			if err = rbWriter.FlushSplit(rg); err != nil {
				return err
			}

		} else if err != nil {
			return err
		}
	}

	return err
}
