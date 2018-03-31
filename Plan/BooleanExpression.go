package Plan

import (
	"github.com/xitongsys/guery/Common"
	"github.com/xitongsys/guery/parser"
)

type BooleanExpressionNode struct {
	Tree           *parser.BooleanExpressionContext
	BooleanDefault *BooleanDefaultNode
	LogicalNot     *LocicalNotNode
	LogicalBinary  *LogicalBinaryNode
}

func NewBooleanExpressionNode(ctx *Context, t *parser.BooleanExpressionContext) *BooleanExpressionNode {
	res := &BooleanExpressionNode{
		Tree: t,
	}
	child := t.GetChildren()[0]
	switch child.(type) {
	case *parser.BooleanDefaultContext:
		res.BooleanDefault = NewBooleanDefaultNode(ctx, child.(*parser.BooleanDefaultContext))
	}
	return res
}

func (self *BooleanExpressionNode) Result(ctx *Context) interface{} {
	if self.BooleanDefault != nil {
		return self.BooleanDefault.Result(ctx)
	}
	return nil
}

//BooleanDefaultNode
type BooleanDefaultNode struct {
	Tree       *parser.BooleanDefaultContext
	Predicated *PredicatedNode
}

func NewBooleanDefaultNode(ctx *Context, t *parser.BooleanDefaultContext) *BooleanDefaultNode {
	res := &BooleanDefaultNode{
		Tree:       t,
		Predicated: NewPredicatedNode(ctx, t.Predicated().(*parser.PredicatedContext)),
	}
	return res
}

func (self *BooleanDefaultNode) Result(ctx *Context) interface{} {
	return self.Predicated.Result(ctx)
}

//LocicalNotNode
type LocicalNotNode struct {
	booleanExpression *BooleanExpressionNode
}

type LogicalBinaryNode struct {
	leftExpression  *BooleanExpressionNode
	rightExpression *BooleanExpressionNode
	operator        *Common.Operator
}
