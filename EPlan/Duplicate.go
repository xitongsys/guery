package EPlan

import (
	"github.com/xitongsys/guery/Logger"
	. "github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/pb"
)

type EPlanDuplicateNode struct {
	Location        *pb.Location
	Inputs, Outputs []*pb.Location
}

func (self *EPlanDuplicateNode) GetNodeType() EPlanNodeType {
	return EDUPLICATE
}

func NewEPlanDuplicateNode(inputs, outputs []*pb.Location) *EPlanDuplicateNode {
	return &EPlanDuplicateNode{
		Location: outputs[0],
		Inputs:   inputs,
		Outputs:  outputs,
	}
}
