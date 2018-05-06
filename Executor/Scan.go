package Executor

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"strings"

	"github.com/xitongsys/guery/Catalog"
	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/pb"
)

func (self *Executor) SetInstructionScan(instruction *pb.Instruction) error {
	var enode EPlan.EPlanScanNode
	var err error
	if err = gob.NewDecoder(bytes.NewBufferString(instruction.EncodedEPlanNodeBytes)).Decode(&enode); err != nil {
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
	if self.Instruction == nil {
		return fmt.Errorf("No Instruction")
	}

	enode := self.EPlanNode.(*EPlan.EPlanScanNode)

	var catalog Catalog.Catalog
	switch strings.ToUpper(self.Instruction.Catalog) {
	case "TEST":
		catalog = Catalog.NewTestCatalog(enode.SourceName)
	case "HIVE":
	}

	ln := len(self.Writers)
	i := 0

	//send metadata
	md := catalog.GetMetadata()

	var buf bytes.Buffer
	if err = gob.NewEncoder(&buf).Encode(md); err != nil {
		return err
	}
	for i := 0; i < ln; i++ {
		Util.WriteMessage(self.Writers[i], buf.Bytes())
	}

	//send rows
	var row *Util.Row
	for {
		row, err = catalog.ReadRow()
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
	return err
}
