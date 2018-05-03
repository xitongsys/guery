package Plan

import (
	"github.com/xitongsys/guery/Common"
	"github.com/xitongsys/guery/DataSource"
	"github.com/xitongsys/guery/parser"
)

type PredicateNode struct {
	ComparisonOperator   *Common.ComparisonOperator
	RightValueExpression *ValueExpressionNode
}

func NewPredicateNode(t parser.IPredicateContext) *PredicateNode {
	tt := t.(*parser.PredicateContext)
	res := &PredicateNode{}
	if iopc, ve := tt.ComparisonOperator(), tt.GetRight(); iopc != nil && ve != nil {
		res.ComparisonOperator = NewComparisonOperator(iopc)
		res.RightValueExpression = NewValueExpressionNode(ve)
	}
	return res
}

func (self *PredicateNode) Result(val interface{}, input *DataSource.DataSource) bool {
	if self.ComparisonOperator != nil && self.RightValueExpression != nil {
		cp := Common.Cmp(val, self.RightValueExpression.Result(input))

		switch *self.ComparisonOperator {
		case Common.EQ:
			return cp == 0
		case Common.NEQ:
			return cp != 0
		case Common.LT:
			return cp < 0
		case Common.LTE:
			return cp <= 0
		case Common.GT:
			return cp > 0
		case Common.GTE:
			return cp >= 0
		}
	}
	return false
}
