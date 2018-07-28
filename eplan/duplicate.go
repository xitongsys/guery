package eplan

import (
	"github.com/xitongsys/guery/pb"
	"github.com/xitongsys/guery/plan"
)

type EPlanDuplicateNode struct {
	Location        pb.Location
	Keys            []*Plan.ValueExpressionNode
	Inputs, Outputs []pb.Location
}

func (self *EPlanDuplicateNode) GetNodeType() EPlanNodeType {
	return EDUPLICATENODE
}

func (self *EPlanDuplicateNode) GetInputs() []pb.Location {
	return self.Inputs
}

func (self *EPlanDuplicateNode) GetOutputs() []pb.Location {
	return self.Outputs
}

func (self *EPlanDuplicateNode) GetLocation() pb.Location {
	return self.Location
}

func NewEPlanDuplicateNode(inputs, outputs []pb.Location, keys []*Plan.ValueExpressionNode) *EPlanDuplicateNode {
	return &EPlanDuplicateNode{
		Location: outputs[0],
		Keys:     keys,
		Inputs:   inputs,
		Outputs:  outputs,
	}
}
