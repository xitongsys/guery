package Executor

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"strings"

	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/pb"
)

func (self *Executor) SetInstructionSelect(instruction *pb.Instruction) (err error) {
	var enode EPlan.EPlanSelectNode
	if err = gob.NewDecoder(bytes.NewBufferString(instruction.EncodedEPlanNodeBytes)).Decode(&enode); err != nil {
		return err
	}
	self.Instruction = instruction
	self.EPlanNode = &enode
	return nil
}

func (self *Executor) RunSelect() (err error) {
	if self.Instruction == nil {
		return fmt.Errorf("No Instruction")
	}
	enode := self.EPlanNode.(*EPlan.EPlanSelectNode)

	md := &Util.Metadata{}
	reader := self.Readers[0]
	writer := self.Writers[0]
	if err = Util.ReadObject(reader, md); err != nil {
		return err
	}

	colNames := []string{}
	for _, item := range enode.SelectItems {
		for _, cname := range item.GetNames() {
			names := strings.Split(cname, ".")
			name := names[len(names)-1]
			colNames = append(colNames, name)
		}
	}

	smd := Util.NewMetadata(md.Name, colNames, nil)

	if err = Util.WriteObject(writer, smd); err != nil {
		return err
	}

	return nil
}
