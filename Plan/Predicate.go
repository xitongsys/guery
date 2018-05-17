package Plan

import (
	"fmt"

	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/parser"
)

type PredicateNode struct {
	ComparisonOperator   *Util.Operator
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

func (self *PredicateNode) GetType(md *Util.Metadata) (Util.Type, error) {
	return Util.BOOL, nil
}

func (self *PredicateNode) GetColumns() ([]string, error) {
	return self.RightValueExpression.GetColumns()
}

func (self *PredicateNode) Result(val interface{}, input *Util.RowsBuffer) (bool, error) {
	if self.ComparisonOperator != nil && self.RightValueExpression != nil {
		res, err := self.RightValueExpression.Result(input)
		if err != nil {
			return false, err
		}
		return Util.OperatorFunc(val, res, *self.ComparisonOperator).(bool), nil

	}
	return false, fmt.Errorf("wrong PredicateNode")
}
