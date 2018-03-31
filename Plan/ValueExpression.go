package Plan

import (
	"github.com/xitongsys/guery/parser"
)

type ValueExpressionNode struct {
	Tree              *parser.ValueExpressionContext
	PrimaryExpression *PrimaryExpressionNode
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
