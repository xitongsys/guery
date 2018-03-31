package Plan

import (
	"fmt"

	"github.com/xitongsys/guery/parser"
)

type NumberNode struct {
	Tree      *parser.NumberContext
	DoubleVal *float64
	IntVal    *int
}

func NewNumberNode(ctx *Context, t *parser.NumberContext) *NumberNode {
	res := &NumberNode{
		Tree: t,
	}
	children := t.GetChildren()
	switch children[0].(type) {
	case *parser.DoubleLiteralContext:
		s := children[0].(*parser.DoubleLiteralContext).GetText()
		var v float64
		fmt.Sscanf(s, "%f", &v)
		res.DoubleVal = &v

	case *parser.IntegerLiteralContext:
		s := children[0].(*parser.IntegerLiteralContext).GetText()
		var v int
		fmt.Sscanf(s, "%d", &v)
		res.IntVal = &v
	}
	return res
}

func (self *NumberNode) Result(ctx *Context) interface{} {
	if self.DoubleVal != nil {
		return *self.DoubleVal
	} else {
		return *self.IntVal
	}
}
