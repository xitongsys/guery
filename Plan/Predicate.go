package Plan

import (
	"fmt"

	"github.com/xitongsys/guery/Config"
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Row"
	"github.com/xitongsys/guery/Type"
	"github.com/xitongsys/guery/parser"
)

type PredicateNode struct {
	ComparisonOperator   *Type.Operator
	RightValueExpression *ValueExpressionNode
}

func NewPredicateNode(runtime *Config.ConfigRuntime, t parser.IPredicateContext) *PredicateNode {
	tt := t.(*parser.PredicateContext)
	res := &PredicateNode{}
	if iopc, ve := tt.ComparisonOperator(), tt.GetRight(); iopc != nil && ve != nil {
		res.ComparisonOperator = NewComparisonOperator(runtime, iopc)
		res.RightValueExpression = NewValueExpressionNode(runtime, ve)
	}
	return res
}

func (self *PredicateNode) GetType(md *Metadata.Metadata) (Type.Type, error) {
	return Type.BOOL, nil
}

func (self *PredicateNode) GetColumns() ([]string, error) {
	return self.RightValueExpression.GetColumns()
}

func (self *PredicateNode) Result(val interface{}, input *Row.RowsGroup) (bool, error) {
	if self.ComparisonOperator != nil && self.RightValueExpression != nil {
		res, err := self.RightValueExpression.Result(input)
		if err != nil {
			return false, err
		}
		return Type.OperatorFunc(val, res, *self.ComparisonOperator).(bool), nil

	}
	return false, fmt.Errorf("wrong PredicateNode")
}
