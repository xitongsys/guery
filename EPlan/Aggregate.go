package EPlan

import (
	"github.com/xitongsys/guery/pb"
)

type EPlanAggregateNode struct {
	Location pb.Location
	Inputs   []pb.Location
	Output   pb.Location
}

func (self *EPlanAggregateNode) GetNodeType() EPlanNodeType {
	return EAGGREGATENODE
}

func (self *EPlanAggregateNode) GetInputs() []pb.Location {
	return self.Inputs
}

func (self *EPlanAggregateNode) GetOutputs() []pb.Location {
	return []pb.Location{self.Output}
}

func (self *EPlanAggregateNode) GetLocation() pb.Location {
	return self.Location
}

func NewEPlanAggregateNode(inputs []pb.Location, output pb.Location) *EPlanAggregateNode {
	return &EPlanAggregateNode{
		Location: output,
		Inputs:   inputs,
		Output:   output,
	}
}
