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

func (self *Executor) SetInstructionFiliter(instruction *pb.Instruction) (err error) {
	var enode EPlan.EPlanFiliterNode
	if err = msgpack.Unmarshal(instruction.EncodedEPlanNodeBytes, &enode); err != nil {
		return err
	}
	self.Instruction = instruction
	self.EPlanNode = &enode
	self.InputLocations = []*pb.Location{&enode.Input}
	self.OutputLocations = []*pb.Location{&enode.Output}
	return nil
}

func (self *Executor) RunFiliter() (err error) {
	defer self.Clear()

	if self.Instruction == nil {
		return fmt.Errorf("No Instruction")
	}
	enode := self.EPlanNode.(*EPlan.EPlanFiliterNode)

	md := &Util.Metadata{}
	reader := self.Readers[0]
	writer := self.Writers[0]
	if err = Util.ReadObject(reader, md); err != nil {
		return err
	}

	//write metadata
	if err = Util.WriteObject(writer, enode.Metadata); err != nil {
		return err
	}

	//write rows
	var row *Util.Row
	var rg *Util.RowsGroup
	for {
		row, err = Util.ReadRow(reader)
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			return err
		}
		rg = Util.NewRowsGroup(md)
		rg.Write(row)
		flag := true
		for _, booleanExpression := range enode.BooleanExpressions {
			rg.Reset()
			if ok, err := booleanExpression.Result(rg); !ok.(bool) && err == nil {
				flag = false
				break
			} else if err != nil {
				return err
			}
		}

		if flag {
			if err = Util.WriteRow(writer, row); err != nil {
				return err
			}
		}
	}

	Util.WriteEOFMessage(writer)

	Logger.Infof("RunFiliter finished")
	return err
}
