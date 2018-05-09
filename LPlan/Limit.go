package Plan

import (
	"context"
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/parser"
)

type PlanLimitNode struct {
	Input       PlanNode
	Metadata    *Util.Metadata
	LimitNumber *int64
}

func NewPlanLimitNode(input PlanNode, t antlr.TerminalNode) *PlanLimitNode {
	res := &PlanLimitNode{
		Input:    input,
		Metadata: Util.NewDefaultMetadata(),
	}
	if ns := t.GetText(); ns != "ALL" {
		var num int64
		fmt.Sscanf(ns, "%d", &num)
		res.LimitNumber = &num
	}
	return res
}

func (self *PlanLimitNode) GetNodeType() PlanNodeType {
	return LIMITNODE
}

func (self *PlanLimitNode) String() string {
	res := "PlanLimitNode {\n"
	res += "Input: " + self.Input.String() + "\n"
	res += "LimitNubmer: " + fmt.Sprint(*self.LimitNumber) + "\n"
	res += "}\n"
	return res
}
