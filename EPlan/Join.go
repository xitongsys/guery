package EPlan

import (
	"github.com/xitongsys/guery/Logger"
	. "github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/pb"
)

type EPlanJoinNode struct {
	Location              *pb.Location
	LeftInput, RightInput *pb.Location
	Output                *pb.Location
	JoinType              JoinType
	JoinCriteria          *JoinCriteriaNode
}

func (self *EPlanJoinNode) GetNodeType() EPlanNodeType {
	return EJOINNODE
}

func NewEPlanJoinNode(node *PlanJoinNode,
	leftInput, rightInput *pb.Location, output *pb.Location) *EPlanJoinNode {
	return &EPlanJoinNode{
		Location:     output,
		LeftInput:    leftInput,
		RightInput:   rightInput,
		Output:       output,
		JoinType:     node.JoinType,
		JoinCriteria: node.JoinCriteria,
	}
}
