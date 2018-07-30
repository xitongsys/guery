package eplan

import (
	"github.com/xitongsys/guery/pb"
	"github.com/xitongsys/guery/plan"
)

type EPlanShuffleNode struct {
	Location        pb.Location
	Keys            []*plan.ExpressionNode
	Inputs, Outputs []pb.Location
}

func (self *EPlanShuffleNode) GetNodeType() EPlanNodeType {
	return ESHUFFLENODE
}

func (self *EPlanShuffleNode) GetInputs() []pb.Location {
	return self.Inputs
}

func (self *EPlanShuffleNode) GetOutputs() []pb.Location {
	return self.Outputs
}

func (self *EPlanShuffleNode) GetLocation() pb.Location {
	return self.Location
}

func NewEPlanShuffleNode(inputs, outputs []pb.Location, keys []*plan.ExpressionNode) *EPlanShuffleNode {
	return &EPlanShuffleNode{
		Location: outputs[0],
		Keys:     keys,
		Inputs:   inputs,
		Outputs:  outputs,
	}
}
