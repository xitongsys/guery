package Plan

import (
	"context"
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/Util"
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
	Inputs   []PlanNode
	Metadata *Util.Metadata
}

func NewPlanCombineNode(plans []PlanNode) *PlanCombineNode {
	return &PlanCombineNode{
		Inputs:   plans,
		Metadata: Util.NewDefaultMetadata(),
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

////////////////////////
type PlanRenameNode struct {
	Rename   string
	Metadata *Util.Metadata
	Input    PlanNode
}

func NewPlanRenameNode(tname string, input PlanNode) *PlanRenameNode {
	return &PlanRenameNode{
		Rename:   tname,
		Metadata: Util.NewDefaultMetadata(),
		Input:    input,
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
	Metadata              *Util.Metadata
	LeftInput, RightInput PlanNode
	JoinType              JoinType
	JoinCriteria          *JoinCriteriaNode
}

func NewPlanJoinNode(leftInput PlanNode, rightInput PlanNode, joinType JoinType, joinCriteria *JoinCriteriaNode) *PlanJoinNode {
	res := &PlanJoinNode{
		Metadata:     Util.NewDefaultMetadata(),
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

/////////////////
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

////////////////
type PlanUnionNode struct {
	LeftInput  PlanNode
	RightInput PlanNode
	Operator   UnionType
	Metadata   *Util.Metadata
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
		Metadata:   Util.NewDefaultMetadata(),
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
	Metadata          *Util.Metadata
	BooleanExpression *BooleanExpressionNode
}

func NewPlanFiliterNode(input PlanNode, t parser.IBooleanExpressionContext) *PlanFiliterNode {
	res := &PlanFiliterNode{
		Input:             input,
		Metadata:          Util.NewDefaultMetadata,
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
	Input    PlanNode
	Metadata *Util.Metadata
	GroupBy  *GroupByNode
}

func NewPlanGroupByNode(input PlanNode, groupBy parser.IGroupByContext) *PlanGroupByNode {
	return &PlanGroupByNode{
		Input:    input,
		Metadata: Util.NewDefaultMetadata(),
		GroupBy:  NewGroupByNode(groupBy),
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
	Metadata    *Util.Metadata
	SelectItems []*SelectItemNode
	IsAggregate bool
}

func NewPlanSelectNode(input PlanNode, items []parser.ISelectItemContext) *PlanSelectNode {
	res := &PlanSelectNode{
		Input:       input,
		Metadata:    Util.NewDefaultMetadata(),
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
	Name     string
	Metadata *Util.Metadata
}

func NewPlanScanNode(name string) *PlanScanNode {
	catalog, schema, table := Util.SplitName(name)
	res := &PlanScanNode{
		Name:     name,
		Metadata: Util.NewMetadata(catalog, schema, table, []string{}, []string{}),
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

func (self *PlanScanNode) GetMetadata() *Metadata {
	md := Util.GetMetadata(self.Metadata.Catalog, self.Metadata.Schema, self.Metadata.Table)
	self.Metadata.ColumnNames = md.ColumnNames
	self.Metadata.ColumnTypes = md.ColumnTypes
	self.Metadata.Reset()
	return self.Metadata
}

//////////////////////

func CreateLogicalTree(sqlStr string) (PlanNode, error) {
	is := antlr.NewInputStream(sqlStr)
	lexer := parser.NewSqlLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewSqlParser(stream)
	tree := p.SingleStatement()
	return NewPlanNodeFromSingleStatement(ctx, tree), nil
}
