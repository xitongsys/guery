package Executor

import (
	"bytes"
	"context"
	"encoding/gob"

	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/pb"
)

func (self *Executor) RunScan(instruction *pb.Instruction) error {
	var enode EPlan.EPlanScanNode
	var err error
	if err = gob.NewDecoder(bytes.NewBufferString(instruction.EncodedEPlanNodeBytes)).Decode(&enode); err != nil {
		return err
	}

	self.OutputLocations = []*pb.Location{}
	for _, loc := range enode.Outputs {
		self.OutputLocations = append(self.OutputLocations, &loc)
	}

	if _, err = self.SetupWriters(context.Background(), &pb.Empty{}); err != nil {
		return err
	}

	return nil
}
