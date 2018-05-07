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
	RENAMENODE
	COMBINENODE
	GROUPBYNODE
)

type JoinType int32

const (
	_ JoinType = iota
	LEFTJOIN
	RIGHTJOIN
	INNERJOIN
)

type UnionType int32

const (
	_ UnionType = iota
	INTERSECT
	UNION
	EXCEPT
)

////////////////////

type PlanNode interface {
	GetNodeType() PlanNodeType
	String() string
}

////////////////////////
type PlanCombineNode struct {
	Inputs []PlanNode
}

func NewPlanCombineNode(plans []PlanNode) *PlanCombineNode {
	return &PlanCombineNode{
		Inputs: plans,
	}
}

func (self *PlanCombineNode) GetNodeType() PlanNodeType {
	return COMBINENODE
}

func (self *PlanCombineNode) String() string {
	res := "PlanCombineNode {\n"
	for _, n := range self.Inputs {
		res += n.String()
	}
	res += "}\n"
	return res
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

func (self *PlanHavingNode) String() string {
	res := "PlanHavingNode {\n"
	res += "Input: " + self.Input.String() + "\n"
	res += "BooleanExpression: " + fmt.Sprint(self.BooleanExpression) + "\n"
	res += "}\n"
	return res
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

func (self *PlanRenameNode) String() string {
	res := "PlanRenameNode {\n"
	res += "Rename: " + self.Rename + "\n"
	res += "Input: " + self.Input.String() + "\n"
	res += "}\n"
	return res
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

func (self *PlanJoinNode) String() string {
	res := "PlanJoinNode {\n"
	res += "LeftInput: " + self.LeftInput.String() + "\n"
	res += "RightInput: " + self.RightInput.String() + "\n"
	res += "JoinType: " + fmt.Sprint(self.JoinType) + "\n"
	res += "JoinCriteria: " + fmt.Sprint(self.JoinCriteria) + "\n"
	res += "}\n"
	return res
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

func (self *PlanOrderByNode) String() string {
	res := "PlanOrderByNode {\n"
	res += "Input: " + self.Input.String() + "\n"
	res += "}\n"
	return res
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

func (self *PlanLimitNode) String() string {
	res := "PlanLimitNode {\n"
	res += "Input: " + self.Input.String() + "\n"
	res += "LimitNubmer: " + fmt.Sprint(*self.LimitNumber) + "\n"
	res += "}\n"
	return res
}

////////////////
type PlanUnionNode struct {
	LeftInput  PlanNode
	RightInput PlanNode
	Operator   UnionType
}

func NewPlanUnionNode(left PlanNode, right PlanNode, op antlr.Token) *PlanUnionNode {
	var operator UnionType
	switch op.GetText() {
	case "INTERSECT":
		operator = INTERSECT
	case "UNION":
		operator = UNION
	case "EXCEPT":
		operator = EXCEPT
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

func (self *PlanUnionNode) String() string {
	res := "PlanUnionNode {\n"
	res += "LeftInput: " + self.LeftInput.String() + "\n"
	res += "RightInput: " + self.RightInput.String() + "\n"
	res += "Operator: " + fmt.Sprint(self.Operator) + "\n"
	res += "}\n"
	return res
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

func (self *PlanFiliterNode) String() string {
	res := "PlanFiliterNode {\n"
	res += "Input: " + self.Input.String() + "\n"
	res += "BooleanExpression: " + fmt.Sprint(self.BooleanExpression) + "\n"
	res += "}\n"
	return res
}

////////////////
type PlanGroupByNode struct {
	Input   PlanNode
	GroupBy *GroupByNode
}

func NewPlanGroupByNode(input PlanNode, groupBy parser.IGroupByContext) *PlanGroupByNode {
	return &PlanGroupByNode{
		Input:   input,
		GroupBy: NewGroupByNode(groupBy),
	}
}

func (self *PlanGroupByNode) GetNodeType() PlanNodeType {
	return GROUPBYNODE
}

func (self *PlanGroupByNode) String() string {
	res := "PlanGroupByNode {\n"
	res += "Input: " + self.Input.String() + "\n"
	res += "GroupBy: " + fmt.Sprint(self.GroupBy) + "\n"
	res += "}/n"
	return res
}

////////////////
type PlanSelectNode struct {
	Input       PlanNode
	SelectItems []*SelectItemNode
	IsAggregate bool
}

func NewPlanSelectNode(input PlanNode, items []parser.ISelectItemContext) *PlanSelectNode {
	res := &PlanSelectNode{
		Input:       input,
		SelectItems: []*SelectItemNode{},
	}
	for i := 0; i < len(items); i++ {
		itemNode := NewSelectItemNode(items[i])
		res.SelectItems = append(res.SelectItems, itemNode)
		if itemNode.IsAggregate() {
			res.IsAggregate = true
		}
	}
	return res
}

func (self *PlanSelectNode) GetNodeType() PlanNodeType {
	return SELECTNODE
}

func (self *PlanSelectNode) String() string {
	res := "PlanSelectNode {\n"
	res += "Input: " + self.Input.String() + "\n"
	res += "SelectItems: " + fmt.Sprint(self.SelectItems) + "\n"
	res += "}\n"
	return res
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

func (self *PlanScanNode) String() string {
	res := "PlanScanNode {\n"
	res += "Name: " + self.Name + "\n"
	res += "}\n"
	return res
}

//////////////////////

func CreateLogicalTree(sqlStr string) (PlanNode, error) {
	is := antlr.NewInputStream(sqlStr)
	lexer := parser.NewSqlLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewSqlParser(stream)
	tree := p.SingleStatement()
	return Plan.NewPlanNodeFromSingleStatement(tree), nil
}
