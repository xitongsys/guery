package eplan

import (
	"github.com/xitongsys/guery/pb"
	"github.com/xitongsys/guery/plan"
)

type EPlanHashJoinShuffleNode struct {
	Location        pb.Location
	Keys            []*Plan.ValueExpressionNode
	Inputs, Outputs []pb.Location
}

func (self *EPlanHashJoinShuffleNode) GetNodeType() EPlanNodeType {
	return EHASHJOINSHUFFLENODE
}

func (self *EPlanHashJoinShuffleNode) GetInputs() []pb.Location {
	return self.Inputs
}

func (self *EPlanHashJoinShuffleNode) GetOutputs() []pb.Location {
	return self.Outputs
}

func (self *EPlanHashJoinShuffleNode) GetLocation() pb.Location {
	return self.Location
}

func NewEPlanHashJoinShuffleNode(inputs, outputs []pb.Location, keys []*Plan.ValueExpressionNode) *EPlanHashJoinShuffleNode {
	return &EPlanHashJoinShuffleNode{
		Location: outputs[0],
		Keys:     keys,
		Inputs:   inputs,
		Outputs:  outputs,
	}
}
