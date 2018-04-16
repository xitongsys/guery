package Plan

import (
	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/DataSource"
	"github.com/xitongsys/guery/parser"
)

type ExpressionNode struct {
	BooleanExpression *BooleanExpressionNode
}

func NewExpressionNode(ctx *Context.Context, t parser.IExpressionContext) *ExpressionNode {
	tt := t.(*parser.ExpressionContext)
	res := &ExpressionNode{
		BooleanExpression: NewBooleanExpressionNode(ctx, tt.BooleanExpression()),
	}
	return res
}

func (self *ExpressionNode) Result(input *DataSource.DataSource) interface{} {
	return self.BooleanExpression.Result(input)
}

func (self *ExpressionNode) IsAggregate() bool {
	return self.BooleanExpression.IsAggregate()
}
