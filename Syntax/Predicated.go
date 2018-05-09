package Plan

import (
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/parser"
)

type PredicatedNode struct {
	Name            string
	ValueExpression *ValueExpressionNode
	Predicate       *PredicateNode
}

func NewPredicatedNode(t parser.IPredicatedContext) *PredicatedNode {
	tt := t.(*parser.PredicatedContext)
	res := &PredicatedNode{}
	res.ValueExpression = NewValueExpressionNode(tt.ValueExpression())
	if tp := tt.Predicate(); tp != nil {
		res.Predicate = NewPredicateNode(tp)
	}
	res.Name = res.ValueExpression.Name
	return res
}

func (self *PredicatedNode) Result(input *Util.RowsBuffer) (interface{}, error) {
	res, err := self.ValueExpression.Result(input)
	if err != nil {
		return nil, err
	}
	if self.Predicate == nil {
		return res, nil
	}
	input.Reset()
	return self.Predicate.Result(res, input)
}

func (self *PredicatedNode) IsAggregate() bool {
	return self.ValueExpression.IsAggregate()
}
