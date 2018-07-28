package eplan

import (
	"github.com/xitongsys/guery/pb"
)

type EPlanBalanceNode struct {
	Location        pb.Location
	Inputs, Outputs []pb.Location
}

func (self *EPlanBalanceNode) GetNodeType() EPlanNodeType {
	return EBALANCENODE
}

func (self *EPlanBalanceNode) GetInputs() []pb.Location {
	return self.Inputs
}

func (self *EPlanBalanceNode) GetOutputs() []pb.Location {
	return self.Outputs
}

func (self *EPlanBalanceNode) GetLocation() pb.Location {
	return self.Location
}

func NewEPlanBalanceNode(inputs, outputs []pb.Location) *EPlanBalanceNode {
	return &EPlanBalanceNode{
		Location: outputs[0],
		Inputs:   inputs,
		Outputs:  outputs,
	}
}
