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
	HAVINGNODE
	UNITNODE
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

func NewPlanUnionNode(left PlanNode, right PlanNode, op antlr.Token) *PlanUnionNode {
	var operator Common.Operator
	switch op.GetText() {
	case "INTERSECT":
		operator = Common.INTERSECT
	case "UNION":
		operator = Common.UNION
	case "EXCEPT":
		operator = Common.EXCEPT
	}

	res := &PlanUniontNode{
		LeftInput:  left,
		RightInput: right,
		Operator:   operator,
	}
	return res
}

func (self *PlanUnionNode) Execute() DataSource.DataSource {
	return self.LeftInput.Execute()
}

//////////////
type PlanFiliterNode struct {
	Input PlanNode
}

func NewPlanFiliterNode(input PlanNode, t antlr.IBooleanExpressionContext) *PlanFiliterNode {
	res := &PlanFiliterNode{
		Input: input,
	}
	return res
}

func (self *PlanFiliterNode) Execute() DataSource.DataSource {
	return self.Input.Execute()
}

//////////////
type PlanHavingNode struct {
	Input PlanNode
}

func NewPlanHavingNode(input PlanNode, t antlr.IBooleanExpressionContext) *PlanHavingNode {
	res := &PlanHavingNode{
		Input: input,
	}
	return res
}

func (self *PlanHavingNode) Execute() DataSource.DataSource {
	return self.Input.Execute()
}

////////////////
type PlanSelectNode struct {
	Input PlanNode
}

func NewPlanSelectNode(intput PlanNode, items []parser.ISelectItemContext, groupBy parser.IGroupBy) *PlanSelectNode {
	res := &PlanSelectNode{
		Input: input,
	}
	return res
}

func (self *PlanSelectNode) Execute() DataSource.DataSource {
	return self.Input.Execute()
}

///////////////////
type PlanScanNode struct {
	Input DataSource.DataSource
}

func NewPlanScanNode(input DataSource.DataSource, r parser.IRelationContext) *PlanScanNode {
	res := &PlanScanNode{
		Input: input,
	}
	return res
}

func (self *PlanScanNode) Execute() DataSource.DataSource {
	return self.Input
}
