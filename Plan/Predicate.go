package Plan

import (
	"github.com/xitongsys/guery/parser"
)

type PredicateNode struct{}

func NewPredicateNode(ctx *Context, t *parser.PredicateContext) *PredicateNode {
	res := &PredicateNode{}
	return res
}

func (self *PredicateNode) Result(ctx *Context) interface{} {
	return nil
}
