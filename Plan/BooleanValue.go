package Plan

import (
	"github.com/xitongsys/guery/parser"
)

type BooleanValueNode struct {
	Tree *parser.BooleanValueContext
	Bool bool
}

func NewBooleanValueNode(ctx *Context, t *parser.BooleanValueContext) *BooleanValueNode {
	s := t.GetText()
	b := true
	if s != "TRUE" {
		b = false
	}
	return &BooleanValueNode{
		Tree: t,
		Bool: b,
	}
}

func (self *BooleanValueNode) Result(ctx *Context) bool {
	return self.Bool
}
