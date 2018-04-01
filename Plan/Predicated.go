package Plan

import (
	"github.com/xitongsys/guery/parser"
)

type PredicatedNode struct {
	Tree            *parser.PredicatedContext
	ValueExpression *ValueExpressionNode
	Predicate       *PredicateNode
}

func NewPredicatedNode(ctx *Context, t *parser.PredicatedContext) *PredicatedNode {
	res := &PredicatedNode{
		Tree: t,
	}
	res.ValueExpression = NewValueExpressionNode(ctx, t.ValueExpression().(*parser.ValueExpressionContext))
	if t.Predicate() != nil {
		res.Predicate = NewPredicateNode(ctx, t.Predicate().(*parser.PredicateContext))
	}
	return res
}

func (self *PredicatedNode) Result(ctx *Context) interface{} {
	return self.ValueExpression.Result(ctx)
}
