package Plan

import (
	"fmt"

	"github.com/xitongsys/guery/Common"
	"github.com/xitongsys/guery/Util"
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

func (self *PredicateNode) Result(val interface{}, input *Util.RowsBuffer) (bool, error) {
	if self.ComparisonOperator != nil && self.RightValueExpression != nil {
		res, err := self.RightValueExpression.Result(input)
		if err != nil {
			return false, err
		}
		cp := Common.Cmp(val, res)

		switch *self.ComparisonOperator {
		case Common.EQ:
			return cp == 0, nil
		case Common.NEQ:
			return cp != 0, nil
		case Common.LT:
			return cp < 0, nil
		case Common.LTE:
			return cp <= 0, nil
		case Common.GT:
			return cp > 0, nil
		case Common.GTE:
			return cp >= 0, nil
		}
	}
	return false, fmt.Errorf("wrong PredicateNode")
}
