package EPlan

import (
	. "github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/pb"
)

type EPlanFiliterNode struct {
	Location          pb.Location
	Input, Output     pb.Location
	BooleanExpression *BooleanExpressionNode
	Metadata          *Util.Metadata
}

func (self *EPlanFiliterNode) GetNodeType() EPlanNodeType {
	return EFILITERNODE
}

func (self *EPlanFiliterNode) GetOutputs() []pb.Location {
	return []pb.Location{self.Output}
}

func (self *EPlanFiliterNode) GetLocation() pb.Location {
	return self.Location
}

func NewEPlanFiliterNode(node *PlanFiliterNode, input, output pb.Location) *EPlanFiliterNode {
	return &EPlanFiliterNode{
		Location:          output,
		Input:             input,
		Output:            output,
		BooleanExpression: node.BooleanExpression,
		Metadata:          node.GetMetadata(),
	}
}
