package EPlan

import (
	"github.com/xitongsys/guery/Metadata"
	. "github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/pb"
)

type EPlanUnionNode struct {
	Location              pb.Location
	LeftInput, RightInput pb.Location
	Output                pb.Location
	Operator              UnionType
	Metadata              *Metadata.Metadata
}

func (self *EPlanUnionNode) GetNodeType() EPlanNodeType {
	return EUNIONNODE
}

func (self *EPlanUnionNode) GetInputs() []pb.Location {
	return []pb.Location{self.LeftInput, self.RightInput}
}

func (self *EPlanUnionNode) SetInputs(inputs []pb.Location) {
	self.LeftInput = inputs[0]
	self.RightInput = inputs[1]
}

func (self *EPlanUnionNode) GetOutputs() []pb.Location {
	return []pb.Location{self.Output}
}

func (self *EPlanUnionNode) GetLocation() pb.Location {
	return self.Location
}

func NewEPlanUnionNode(node *PlanUnionNode, leftInput, rightInput pb.Location, output pb.Location) *EPlanUnionNode {
	return &EPlanUnionNode{
		Location:   output,
		LeftInput:  leftInput,
		RightInput: rightInput,
		Output:     output,
		Operator:   node.Operator,
		Metadata:   node.GetMetadata(),
	}
}
