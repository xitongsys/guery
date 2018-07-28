package plan

import (
	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/gtype"
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/parser"
	"github.com/xitongsys/guery/row"
)

type ExpressionNode struct {
	Name              string
	BooleanExpression *BooleanExpressionNode
}

func NewExpressionNode(runtime *config.ConfigRuntime, t parser.IExpressionContext) *ExpressionNode {
	tt := t.(*parser.ExpressionContext)
	res := &ExpressionNode{
		Name:              "",
		BooleanExpression: NewBooleanExpressionNode(runtime, tt.BooleanExpression()),
	}
	res.Name = res.BooleanExpression.Name
	return res
}

func (self *ExpressionNode) ExtractAggFunc(res *[]*FuncCallNode) {
	self.BooleanExpression.ExtractAggFunc(res)
}

func (self *ExpressionNode) GetType(md *metadata.Metadata) (gtype.Type, error) {
	return self.BooleanExpression.GetType(md)
}

func (self *ExpressionNode) GetColumns() ([]string, error) {
	return self.BooleanExpression.GetColumns()
}

func (self *ExpressionNode) Init(md *metadata.Metadata) error {
	return self.BooleanExpression.Init(md)
}

func (self *ExpressionNode) Result(input *row.RowsGroup) (interface{}, error) {
	return self.BooleanExpression.Result(input)
}

func (self *ExpressionNode) IsAggregate() bool {
	return self.BooleanExpression.IsAggregate()
}
