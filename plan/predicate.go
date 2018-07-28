package plan

import (
	"fmt"

	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/gtype"
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/parser"
	"github.com/xitongsys/guery/row"
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

func (self *PredicateNode) ExtractAggFunc(res *[]*FuncCallNode) {
	self.RightValueExpression.ExtractAggFunc(res)
}

func (self *PredicateNode) GetColumns() ([]string, error) {
	return self.RightValueExpression.GetColumns()
}

func (self *PredicateNode) Init(md *Metadata.Metadata) error {
	if self.RightValueExpression != nil {
		return self.RightValueExpression.Init(md)

	}
	return nil
}

func (self *PredicateNode) Result(valsi interface{}, input *Row.RowsGroup) (interface{}, error) {
	if self.ComparisonOperator != nil && self.RightValueExpression != nil {
		resi, err := self.RightValueExpression.Result(input)
		if err != nil {
			return nil, err
		}
		vals, res := valsi.([]interface{}), resi.([]interface{})
		for i := 0; i < len(res); i++ {
			res[i] = Type.OperatorFunc(vals[i], res[i], *self.ComparisonOperator)
		}
		return res, nil
	} else {
		return false, fmt.Errorf("wrong PredicateNode")
	}
}
