package Plan

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/Common"
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
	RENAMENODE
)

type JoinType int32

const (
	_ JoinType = iota
	LEFTJOIN
	RIGHTJOIN
	INNERJOIN
)

type PlanNode interface {
	GetNodeType() PlanNodeType
}

////////////////////////
type PlanHavingNode struct {
	Input             PlanNode
	BooleanExpression *BooleanExpressionNode
}

func NewPlanHavingNode(input PlanNode, be parser.IBooleanExpressionContext) *PlanHavingNode {
	return &PlanHavingNode{
		Input:             input,
		BooleanExpression: nil,
	}
}

func (self *PlanHavingNode) GetNodeType() PlanNodeType {
	return HAVINGNODE
}

////////////////////////
type PlanRenameNode struct {
	Rename string
	Input  PlanNode
}

func NewPlanRenameNode(tname string, input PlanNode) *PlanRenameNode {
	return &PlanRenameNode{
		Rename: tname,
		Input:  input,
	}
}

func (self *PlanRenameNode) GetNodeType() PlanNodeType {
	return RENAMENODE
}

///////////////////
type PlanJoinNode struct {
	LeftInput, RightInput PlanNode
	JoinType              JoinType
	JoinCriteria          *JoinCriteriaNode
}

func NewPlanJoinNode(leftInput PlanNode, rightInput PlanNode, joinType JoinType, joinCriteria *JoinCriteriaNode) *PlanJoinNode {
	res := &PlanJoinNode{
		LeftInput:    leftInput,
		RightInput:   rightInput,
		JoinType:     joinType,
		JoinCriteria: joinCriteria,
	}
	return res
}

func (self *PlanJoinNode) GetNodeType() PlanNodeType {
	return JOINNODE
}

/////////////////
type PlanOrderByNode struct {
	Input PlanNode
}

func NewPlanOrderByNode(input PlanNode, items []parser.ISortItemContext) *PlanOrderByNode {
	return &PlanOrderByNode{
		Input: input,
	}
}

func (self *PlanOrderByNode) GetNodeType() PlanNodeType {
	return ORDERBYNODE
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

func (self *PlanLimitNode) GetNodeType() PlanNodeType {
	return LIMITNODE
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

	res := &PlanUnionNode{
		LeftInput:  left,
		RightInput: right,
		Operator:   operator,
	}
	return res
}

func (self *PlanUnionNode) GetNodeType() PlanNodeType {
	return UNIONNODE
}

//////////////
type PlanFiliterNode struct {
	Input             PlanNode
	BooleanExpression *BooleanExpressionNode
}

func NewPlanFiliterNode(input PlanNode, t parser.IBooleanExpressionContext) *PlanFiliterNode {
	res := &PlanFiliterNode{
		Input:             input,
		BooleanExpression: NewBooleanExpressionNode(t),
	}
	return res
}

func (self *PlanFiliterNode) GetNodeType() PlanNodeType {
	return FILTERNODE
}

////////////////
type PlanSelectNode struct {
	Input       PlanNode
	SelectItems []*SelectItemNode
	GroupBy     *GroupByNode
}

func NewPlanSelectNode(input PlanNode, items []parser.ISelectItemContext, groupBy parser.IGroupByContext) *PlanSelectNode {
	res := &PlanSelectNode{
		Input:       input,
		SelectItems: []*SelectItemNode{},
		GroupBy:     NewGroupByNode(groupBy),
	}
	for i := 0; i < len(items); i++ {
		res.SelectItems = append(res.SelectItems, NewSelectItemNode(items[i]))
	}
	return res
}

func (self *PlanSelectNode) GetNodeType() PlanNodeType {
	return SELECTNODE
}

///////////////////
type PlanScanNode struct {
	Name string
}

func NewPlanScanNode(name string) *PlanScanNode {
	res := &PlanScanNode{
		Name: name,
	}
	return res
}

func (self *PlanScanNode) GetNodeType() PlanNodeType {
	return SCANNODE
}

//////////////////////
