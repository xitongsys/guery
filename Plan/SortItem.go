package Plan

import (
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Row"
	"github.com/xitongsys/guery/Type"
	"github.com/xitongsys/guery/parser"
)

type SortItemNode struct {
	Expression *ExpressionNode
	OrderType  Type.OrderType
}

func NewSortItemNode(t parser.ISortItemContext) *SortItemNode {
	tt := t.(*parser.SortItemContext)
	res := &SortItemNode{
		Expression: NewExpressionNode(tt.Expression()),
		OrderType:  Type.ASC,
	}

	if ot := tt.GetOrdering(); ot != nil {
		if ot.GetText() != "ASC" {
			res.OrderType = Type.DESC
		}
	}

	return res
}

func (self *SortItemNode) GetColumns() ([]string, error) {
	return self.Expression.GetColumns()
}

func (self *SortItemNode) Result(input *Row.RowsGroup) (interface{}, error) {
	return self.Expression.Result(input)
}

func (self *SortItemNode) IsAggregate() bool {
	return self.Expression.IsAggregate()
}

func (self *SortItemNode) GetType(md *Metadata.Metadata) (Type.Type, error) {
	return self.Expression.GetType(md)
}
