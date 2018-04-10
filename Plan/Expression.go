package Plan

import (
	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/parser"
)

type ExpressionNode struct {
	BooleanExpression *BooleanExpressionNode
}

func NewExpressionNode(ctx *Context.Context, t parser.IExpressionContext) *ExpressionNode {
	res := &ExpressionNode{
		BooleanExpression: NewBooleanExpressionNode(ctx, t.BooleanExpression()),
	}
	return res
}

func (self *ExpressionNode) Result(input DataSource.DataSource) interface{} {
	self.BooleanExpression.Result(input)
}
