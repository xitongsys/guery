package EPlan

import (
	"github.com/xitongsys/guery/Logger"
	. "github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/pb"
)

type EPlanFiliterNode struct {
	Location          *pb.Location
	Input, Output     *pb.Location
	BooleanExpression *BooleanExpressionNode
}

func (self *EPlanFiliterNode) GetNodeType() EPlanNodeType {
	return EFILITERNODE
}

func NewEPlanFiliterNode(node *PlanFiliterNode, input, output *pb.Location) *EPlanFiliterNode {
	return &EPlanFiliterNode{
		Location:          output,
		Input:             input,
		Output:            output,
		BooleanExpression: node.BooleanExpression,
	}
}
