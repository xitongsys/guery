package Plan

import (
	"github.com/xitongsys/guery/parser"
)

type PredicatedNode struct {
	ValueExpression *ValueExpressionNode
	Predicate       *PredicateNode
}

func NewPredicatedNode(t parser.IPredicatedContext) *PredicatedNode {
	tt := t.(*parser.PredicateContext)
	res := &PredicatedNode{}
	res.ValueExpression = NewValueExpressionNode(tt.ValueExpression())
	if t.Predicate() != nil {
		res.Predicate = NewPredicateNode(t.Predicate())
	}
	return res
}

func (self *PredicatedNode) Result(input DataSource.DataSource) interface{} {
	return self.ValueExpression.Result(input)
}
