package eplan

import (
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/pb"
	. "github.com/xitongsys/guery/plan"
)

type EPlanHashJoinNode struct {
	Location                pb.Location
	LeftInputs, RightInputs []pb.Location
	Output                  pb.Location
	JoinType                JoinType
	JoinCriteria            *JoinCriteriaNode
	LeftKeys, RightKeys     []*ValueExpressionNode
	Metadata                *metadata.Metadata
}

func (self *EPlanHashJoinNode) GetNodeType() EPlanNodeType {
	return EHASHJOINNODE
}

func (self *EPlanHashJoinNode) GetInputs() []pb.Location {
	res := []pb.Location{}
	res = append(res, self.LeftInputs...)
	res = append(res, self.RightInputs...)
	return res
}

func (self *EPlanHashJoinNode) GetOutputs() []pb.Location {
	return []pb.Location{self.Output}
}

func (self *EPlanHashJoinNode) GetLocation() pb.Location {
	return self.Location
}

func NewEPlanHashJoinNode(node *PlanHashJoinNode,
	leftInputs, rightInputs []pb.Location, output pb.Location) *EPlanHashJoinNode {
	return &EPlanHashJoinNode{
		Location:     output,
		LeftInputs:   leftInputs,
		RightInputs:  rightInputs,
		Output:       output,
		JoinType:     node.JoinType,
		JoinCriteria: node.JoinCriteria,
		LeftKeys:     node.LeftKeys,
		RightKeys:    node.RightKeys,
		Metadata:     node.GetMetadata(),
	}
}
