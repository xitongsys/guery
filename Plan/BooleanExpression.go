package Plan

import (
	"github.com/xitongsys/guery/parser"
)

type BooleanExpressionNode struct {
	Tree       *parser.BooleanExpressionContext
	Predicated *PredicatedNode
}

func NewBooleanExpressionNode(ctx *Context, t *parser.BooleanExpressionContext) *BooleanExpressionNode {
	res := &BooleanExpressionNode{
		Tree: t,
	}
	children := t.GetChildren()
	if len(children) == 1 {
		res.Predicated = NewPredicatedNode(ctx,
			t.Predicated().(*parser.PredicatedContext))
	}
	return res
}

func (self *BooleanExpressionNode) Result(ctx *Context) interface{} {
	if self.Predicated != nil {
		return self.Predicated.Result(ctx)
	}
	return nil
}
