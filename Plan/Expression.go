package Plan

import (
	"github.com/xitongsys/guery/parser"
)

type ExpressionNode struct {
	Tree              *parser.ExpressionContext
	BooleanExpression *BooleanExpressionNode
}

func NewExpressionNode(ctx *Context, t *parser.ExpressionContext) *ExpressionNode {
	res := &ExpressionNode{
		Tree: t,
	}
	child := t.GetChildren()[0]
	res.BooleanExpression = NewBooleanExpressionNode(ctx, child.(*parser.BooleanExpressionContext))
	return res
}

func (self *ExpressionNode) Result(ctx *Context) interface{} {
	return self.BooleanExpression.Result(ctx)
}
