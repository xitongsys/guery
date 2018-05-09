package Plan

import (
	"context"
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/parser"
)

type PlanOrderByNode struct {
	Input    PlanNode
	Metadata *Util.Metadata
}

func NewPlanOrderByNode(input PlanNode, items []parser.ISortItemContext) *PlanOrderByNode {
	return &PlanOrderByNode{
		Input:    input,
		Metadata: Util.NewDefaultMetadata(),
	}
}

func (self *PlanOrderByNode) GetNodeType() PlanNodeType {
	return ORDERBYNODE
}

func (self *PlanOrderByNode) String() string {
	res := "PlanOrderByNode {\n"
	res += "Input: " + self.Input.String() + "\n"
	res += "}\n"
	return res
}
