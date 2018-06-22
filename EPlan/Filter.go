package EPlan

import (
	"github.com/xitongsys/guery/Metadata"
	. "github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/pb"
)

type EPlanFilterNode struct {
	Location           pb.Location
	Input, Output      pb.Location
	BooleanExpressions []*BooleanExpressionNode
	Metadata           *Metadata.Metadata
}

func (self *EPlanFilterNode) GetNodeType() EPlanNodeType {
	return EFILTERNODE
}

func (self *EPlanFilterNode) GetInputs() []pb.Location {
	return []pb.Location{self.Input}
}

func (self *EPlanFilterNode) SetInputs(inputs []pb.Location) {
	self.Input = inputs[0]
}

func (self *EPlanFilterNode) GetOutputs() []pb.Location {
	return []pb.Location{self.Output}
}

func (self *EPlanFilterNode) GetLocation() pb.Location {
	return self.Location
}

func NewEPlanFilterNode(node *PlanFilterNode, input, output pb.Location) *EPlanFilterNode {
	return &EPlanFilterNode{
		Location:           output,
		Input:              input,
		Output:             output,
		BooleanExpressions: node.BooleanExpressions,
		Metadata:           node.GetMetadata(),
	}
}
