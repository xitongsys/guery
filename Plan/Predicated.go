package Plan

import (
	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/parser"
)

type PredicatedNode struct {
	ValueExpression *ValueExpressionNode
	Predicate       *PredicateNode
}

func NewPredicatedNode(ctx *Context.Context, t parser.IPredicatedContext) *PredicatedNode {
	tt := t.(*parser.PredicateContext)
	res := &PredicatedNode{}
	res.ValueExpression = NewValueExpressionNode(ctx, tt.ValueExpression())
	if t.Predicate() != nil {
		res.Predicate = NewPredicateNode(ctx, t.Predicate())
	}
	return res
}

func (self *PredicatedNode) Result(input DataSource.DataSource) interface{} {
	return self.ValueExpression.Result(input)
}
