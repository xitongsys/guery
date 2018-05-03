package Plan

import (
	"github.com/xitongsys/guery/DataSource"
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

func (self *PredicatedNode) Result(input *DataSource.DataSource) interface{} {
	res := self.ValueExpression.Result(input)
	if self.Predicate == nil {
		return res
	}
	input.Reset()
	return self.Predicate.Result(res, input)
}

func (self *PredicatedNode) IsAggregate() bool {
	return self.ValueExpression.IsAggregate()
}
