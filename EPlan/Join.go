package EPlan

import (
	. "github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/pb"
)

type EPlanJoinNode struct {
	Location              pb.Location
	LeftInput, RightInput pb.Location
	Output                pb.Location
	JoinType              JoinType
	JoinCriteria          *JoinCriteriaNode
}

func (self *EPlanJoinNode) GetNodeType() EPlanNodeType {
	return EJOINNODE
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
	}
}
