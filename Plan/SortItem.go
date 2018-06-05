package Plan

import (
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/parser"
)

type SortItemNode struct {
	Expression *ExpressionNode
	OrderType  Util.OrderType
}

func NewSortItemNode(t parser.ISortItemContext) *SortItemNode {
	tt := t.(*parser.SortItemContext)
	res := &SortItemNode{
		Expression: NewExpressionNode(tt.Expression()),
		OrderType:  Util.ASC,
	}

	if ot := tt.GetOrdering(); ot != nil {
		if ot.GetText() != "ASC" {
			res.OrderType = Util.DESC
		}
	}

	return res
}

func (self *SortItemNode) GetColumns() ([]string, error) {
	return self.Expression.GetColumns()
}

func (self *SortItemNode) Result(input *Util.RowsGroup) (interface{}, error) {
	return self.Expression.Result(input)
}
