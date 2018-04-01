package Plan

import (
	"github.com/xitongsys/guery/Common"
	"github.com/xitongsys/guery/parser"
)

type ValueExpressionNode struct {
	Tree                  *parser.ValueExpressionContext
	PrimaryExpression     *PrimaryExpressionNode
	ValueExpression       *ValueExpressionNode
	BinaryVauleExpression *BinaryValueExpressionNode
}

func NewValueExpressionNode(ctx *Context, t *parser.ValueExpressionContext) *ValueExpressionNode {
	res := &ValueExpressionNode{
		Tree: t,
	}
	children := t.GetChildren()
	switch children[0].(type) {
	case *parser.PrimaryExpressionContext:
		res.PrimaryExpression = NewPrimaryExpressionNode(ctx, children[0].(*parser.PrimaryExpressionContext))
	}
	return res
}

func (self *ValueExpressionNode) Result(ctx *Context) interface{} {
	if self.PrimaryExpression != nil {
		return self.PrimaryExpression.Result(ctx)
	}
	return nil
}

/////////////////
type BinaryValueExpressionNode struct {
	LeftValueExpression  *ValueExpressionNode
	RightValueExpression *ValueExpressionNode
	Operator             *Common.Operator
}

func NewBinaryValueExpressionNode(ctx *Context,
	left *parser.ValueExpressionContext, right *parser.ValueExpressionContext,
	op *Common.Operator) *BinaryValueExpressionNode {

	res := &BinaryValueExpressionNode{
		LeftValueExpression:  NewValueExpressionNode(ctx, left),
		RightValueExpression: NewValueExpressionNode(ctx, right),
		Operator:             op,
	}
	return res
}

func (self *BinaryValueExpressionNode) Result(ctx *Context) interface{} {
	leftVal, rightVal := self.LeftValueExpression.Result(ctx), self.RightValueExpression.Result(ctx)
	return Common.Arithmetic(leftVal, rightVal, *self.Operator)
}
