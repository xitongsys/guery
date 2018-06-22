package EPlan

import (
	"github.com/xitongsys/guery/Metadata"
	. "github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/pb"
)

type EPlanHashJoinNode struct {
	Location              pb.Location
	LeftInput, RightInput pb.Location
	Output                pb.Location
	JoinType              JoinType
	JoinCriteria          *JoinCriteriaNode
	LeftKeys, RightKeys   []*ValueExpressionNode
	Metadata              *Metadata.Metadata
}

func (self *EPlanHashJoinNode) GetNodeType() EPlanNodeType {
	return EHASHJOINNODE
}

func (self *EPlanHashJoinNode) GetInputs() []pb.Location {
	return []pb.Location{self.LeftInput, self.RightInput}
}

func (self *EPlanHashJoinNode) SetInputs(inputs []pb.Location) {
	self.LeftInput = inputs[0]
	self.RightInput = inputs[1]
}

func (self *EPlanHashJoinNode) GetOutputs() []pb.Location {
	return []pb.Location{self.Output}
}

func (self *EPlanHashJoinNode) GetLocation() pb.Location {
	return self.Location
}

func NewEPlanHashJoinNode(node *PlanHashJoinNode,
	leftInput, rightInput pb.Location, output pb.Location) *EPlanHashJoinNode {
	return &EPlanHashJoinNode{
		Location:     output,
		LeftInput:    leftInput,
		RightInput:   rightInput,
		Output:       output,
		JoinType:     node.JoinType,
		JoinCriteria: node.JoinCriteria,
		LeftKeys:     node.LeftKeys,
		RightKeys:    node.RightKeys,
		Metadata:     node.GetMetadata(),
	}
}
