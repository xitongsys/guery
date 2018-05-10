package Plan

import (
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/parser"
)

type ExpressionNode struct {
	Name              string
	BooleanExpression *BooleanExpressionNode
}

func NewExpressionNode(t parser.IExpressionContext) *ExpressionNode {
	tt := t.(*parser.ExpressionContext)
	res := &ExpressionNode{
		Name:              "",
		BooleanExpression: NewBooleanExpressionNode(tt.BooleanExpression()),
	}
	res.Name = res.BooleanExpression.Name
	return res
}

func (self *ExpressionNode) GetType(md *Util.Metadata) (Util.Type, error) {
	return self.BooleanExpression.GetType(md)
}

func (self *ExpressionNode) Result(input *Util.RowsBuffer) (interface{}, error) {
	return self.BooleanExpression.Result(input)
}

func (self *ExpressionNode) IsAggregate() bool {
	return self.BooleanExpression.IsAggregate()
}
