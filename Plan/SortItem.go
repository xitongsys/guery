package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/Util"
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

	if ot := tt.Ordering(); ot != nil {
		if ot.GetText() != "ASC" {
			res.OrderType = DESC
		}
	}

	return res

}
