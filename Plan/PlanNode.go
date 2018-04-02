package Plan

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/parser"
)

type PlanNodeType int32

const (
	_ PlanNodeType = iota
	SCANNODE
	JOINNODE
	FILTERNODE
	ORDERBYNODE
	LIMITNODE
	SELECTNODE
	UNIONNODE
)

type PlanNode interface {
	Execute() DataSource.DataSource
}

//////////////////
type PlanOrderByNode struct {
	Input PlanNode
}

func NewPlanOrderByNode(input PlanNode, sortItems []parser.ISortItemContext) *PlanOrderByNode {
	res := &PlanOrderByNode{
		Input: input,
	}
	return res
}

func (self *PlanOrderByNode) Execute() DataSource.DataSource {
	return self.Input.Execute()
}

/////////////////
type PlanLimitNode struct {
	Input       PlanNode
	LimitNumber *int64
}

func NewPlanLimitNode(input PlanNode, t antlr.TerminalNode) *PlanLimitNode {
	res := &PlanLimitNode{
		Input: input,
	}
	if ns := t.GetText(); ns != "ALL" {
		var num int64
		fmt.Sscanf(ns, "%d", &num)
		res.LimitNumber = &num
	}
	return res
}

func (self *PlanLimitNode) Execute() DataSource.DataSource {
	return self.Input.Execute()
}

////////////////
type PlanUnionNode struct {
	LeftInput  PlanNode
	RightInput PlanNode
	Operator   Common.Operator
}

func NewPlanUnionNode(left PlanNode, right PlanNode, op Common.Operator) *PlanUnionNode {
	res := &PlanUniontNode{
		LeftInput:  left,
		RightInput: right,
		Operator:   op,
	}
	return res
}

func (self *PlanUnionNode) Execute() DataSource.DataSource {
	return self.LeftInput.Execute()
}
