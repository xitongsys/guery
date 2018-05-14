package Plan

import (
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/parser"
)

type OrderType int32

const (
	UNKNOWNORDERTYPE OrderType = iota
	ASC
	DESC
	FIRST
	LAST
)

type SortItemNode struct {
	Expression *ExpressionNode
	OrderType  OrderType
}

func NewSortItemNode(t parser.ISortItemContext) *SortItemNode {
	tt := t.(*parser.SortItemContext)
	res := &SortItemNode{
		Expression: NewExpressionNode(tt.Expression()),
		OrderType:  ASC,
	}

	if ot := tt.GetOrdering(); ot != nil {
		if ot.GetText() != "ASC" {
			res.OrderType = DESC
		}
	}

	return res
}

func (self *SortItemNode) Result(input *Util.RowsBuffer) (interface{}, error) {
	return self.Expression.Result(input)
}
