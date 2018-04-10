package Plan

import (
	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/parser"
)

type PredicateNode struct{}

func NewPredicateNode(ctx *Context.Context, t parser.IPredicateContext) *PredicateNode {
	res := &PredicateNode{}
	return res
}
