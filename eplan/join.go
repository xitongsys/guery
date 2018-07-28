package eplan

import (
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/pb"
	. "github.com/xitongsys/guery/plan"
)

type EPlanJoinNode struct {
	Location              pb.Location
	LeftInput, RightInput pb.Location
	Output                pb.Location
	JoinType              JoinType
	JoinCriteria          *JoinCriteriaNode
	Metadata              *Metadata.Metadata
}

func (self *EPlanJoinNode) GetNodeType() EPlanNodeType {
	return EJOINNODE
}

func (self *EPlanJoinNode) GetInputs() []pb.Location {
	return []pb.Location{self.LeftInput, self.RightInput}
}

func (self *EPlanJoinNode) GetOutputs() []pb.Location {
	return []pb.Location{self.Output}
}

func (self *EPlanJoinNode) GetLocation() pb.Location {
	return self.Location
}

func NewEPlanJoinNode(node *PlanJoinNode,
	leftInput, rightInput pb.Location, output pb.Location) *EPlanJoinNode {
	return &EPlanJoinNode{
		Location:     output,
		LeftInput:    leftInput,
		RightInput:   rightInput,
		Output:       output,
		JoinType:     node.JoinType,
		JoinCriteria: node.JoinCriteria,
		Metadata:     node.GetMetadata(),
	}
}
