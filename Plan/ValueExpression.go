package Plan

import (
	"github.com/xitongsys/guery/Common"
	"github.com/xitongsys/guery/parser"
)

type ValueExpressionNode struct {
	Tree                   *parser.ValueExpressionContext
	ValueExpressionDefault *ValueExpressionDefaultNode
	ArithmeticUnaryNode    *ArithmeticBinaryNode
	Concatenation          *ConcatenationNode
}

func NewValueExpressionNode(ctx *Context, t *parser.ValueExpressionContext) *ValueExpressionNode {
	res := &ValueExpressionNode{
		Tree: t,
	}
	children := t.GetChildren()
	switch children[0].(type) {
	case *parser.ValueExpressionDefaultContext:
		res.ValueExpressionDefault = NewValueExpressionDefaultNode(ctx, children[0].(*parser.ValueExpressionDefaultContext))
	}
	return res
}

func (self *ValueExpressionNode) Result(ctx *Context) interface{} {
	if self.ValueExpressionDefault != nil {
		return self.ValueExpressionDefault.Result(ctx)
	}
	return nil
}

//ValueExpressionDefaultNode
type ValueExpressionDefaultNode struct {
	Tree              *parser.ValueExpressionDefaultContext
	PrimaryExpression *PrimaryExpressionNode
}

func NewValueExpressionDefaultNode(ctx *Context, t *parser.ValueExpressionDefaultContext) *ValueExpressionDefaultNode {
	res := &ValueExpressionDefaultNode{
		Tree: t,
		PrimaryExpression: NewPrimaryExpressionNode(ctx,
			t.PrimaryExpression().(*parser.PrimaryExpressionContext)),
	}
	return res
}

func (self *ValueExpressionDefaultNode) Result(ctx *Context) interface{} {
	return self.PrimaryExpression.Result(ctx)
}

//
type ArithmeticUnaryNode struct {
	tree            *parser.ArithmeticUnaryContext
	operator        *Common.Operator
	valueExpression *ValueExpressionNode
}

type ArithmeticBinaryNode struct {
	leftExpression  *ValueExpressionNode
	rightExpression *ValueExpressionNode
	operator        *Common.Operator
}

type ConcatenationNode struct {
	leftExpression  *ValueExpressionNode
	rightExpression *ValueExpressionNode
}
