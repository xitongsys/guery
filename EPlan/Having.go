package EPlan

import (
	"github.com/xitongsys/guery/Logger"
	. "github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/pb"
)

type EPlanHavingNode struct {
	Location          *pb.Location
	Input             *pb.Location
	Output            *pb.Location
	BooleanExpression *BooleanExpressionNode
}

func (self *EPlanHavingNode) GetNodeType() EPlanNodeType {
	return EHAVINGNODE
}

func NewEPlanHavingNode(node *PlanHavingNode, input, output *pb.Location) *EPlanHavingNode {
	return &EPlanHavingNode{
		Location:          output,
		Input:             input,
		Output:            output,
		BooleanExpression: node.BooleanExpression,
	}
}
