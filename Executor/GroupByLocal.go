package Executor

import (
	"fmt"
	"io"

	"github.com/vmihailenco/msgpack"
	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Logger"
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
	md := &Util.Metadata{}

	if err = Util.ReadObject(reader, md); err != nil {
		return err
	}

	//write metadata
	if err = Util.WriteObject(writer, enode.Metadata); err != nil {
		return err
	}

	//group by
	var row *Util.Row
	var rgs = make(map[string]*Util.RowsGroup)
	for {
		row, err = Util.ReadRow(reader)
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}
		key := row.GetKeyString()
		if _, ok := rgs[key]; !ok {
			rgs[key] = Util.NewRowsGroup(enode.Metadata)
		}
		rgs[key].Write(row)
	}

	defer func() {
		Util.WriteEOFMessage(writer)
	}()

	//write rows
	for _, rb := range rgs {
		var (
			ok  interface{} = true
			err error       = nil
		)
		if enode.GroupBy.Having != nil {
			ok, err = enode.GroupBy.Having.Result(rb)
		}

		if err == nil && ok.(bool) {
			rb.Reset()
			for {
				row, err := rb.Read()
				if err != nil {
					return err
				}
				if err = Util.WriteRow(writer, row); err != nil {
					return err
				}
			}
		} else if err != nil {
			return err
		}
	}

	return err
}
