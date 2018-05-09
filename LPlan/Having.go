package Plan

import (
	"context"
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/parser"
)

type PlanHavingNode struct {
	Input             PlanNode
	Metadata          *Util.Metadata
	BooleanExpression *BooleanExpressionNode
}

func NewPlanHavingNode(input PlanNode, be parser.IBooleanExpressionContext) *PlanHavingNode {
	return &PlanHavingNode{
		Input:             input,
		Metadata:          Util.NewDefaultMetadata(),
		BooleanExpression: nil,
	}
}

func (self *PlanHavingNode) GetNodeType() PlanNodeType {
	return HAVINGNODE
}

func (self *PlanHavingNode) String() string {
	res := "PlanHavingNode {\n"
	res += "Input: " + self.Input.String() + "\n"
	res += "BooleanExpression: " + fmt.Sprint(self.BooleanExpression) + "\n"
	res += "}\n"
	return res
}
