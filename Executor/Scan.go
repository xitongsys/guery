package Executor

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"

	"github.com/xitongsys/guery/Catalog"
	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/pb"
)

func (self *Executor) SetInstructionScan(instruction *pb.Instruction) error {
	Logger.Infof("set instruction scan")
	var enode EPlan.EPlanScanNode
	var err error
	if err = gob.NewDecoder(bytes.NewBuffer(instruction.EncodedEPlanNodeBytes)).Decode(&enode); err != nil {
		return err
	}

	self.EPlanNode = &enode
	self.Instruction = instruction
	self.OutputLocations = []*pb.Location{}
	for _, loc := range enode.Outputs {
		self.OutputLocations = append(self.OutputLocations, &loc)
	}

	return nil
}

func (self *Executor) RunScan() (err error) {
	defer self.Clear()

	if self.Instruction == nil {
		return fmt.Errorf("No Instruction")
	}

	enode := self.EPlanNode.(*EPlan.EPlanScanNode)

	catalog := Catalog.NewCatalog(self.Instruction.Catalog, self.Instruction.Schema, enode.SourceName)

	ln := len(self.Writers)
	i := 0

	//send metadata
	md := catalog.GetMetadata()
	for i := 0; i < ln; i++ {
		if err = Util.WriteObject(self.Writers[i], md); err != nil {
			return err
		}
	}

	//send rows
	var row *Util.Row
	for {
		row, err = catalog.ReadRow()
		//Logger.Infof("===%v, %v", row, err)
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		if err = Util.WriteRow(self.Writers[i], row); err != nil {
			break
		}

		i++
		i = i % ln
	}

	for i := 0; i < ln; i++ {
		Util.WriteEOFMessage(self.Writers[i])
		self.Writers[i].(io.WriteCloser).Close()
	}
	Logger.Infof("RunScan finished")
	return err
}
