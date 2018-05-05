package EPlan

import (
	"github.com/xitongsys/guery/pb"
)

type EPlanDuplicateNode struct {
	Location        pb.Location
	Inputs, Outputs []pb.Location
}

func (self *EPlanDuplicateNode) GetNodeType() EPlanNodeType {
	return EDUPLICATENODE
}

func (self *EPlanDuplicateNode) GetOutputs() []pb.Location {
	return self.Outputs
}

func NewEPlanDuplicateNode(inputs, outputs []pb.Location) *EPlanDuplicateNode {
	return &EPlanDuplicateNode{
		Location: outputs[0],
		Inputs:   inputs,
		Outputs:  outputs,
	}
}
