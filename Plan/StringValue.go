package Plan

import (
	"github.com/xitongsys/guery/parser"
)

type StringValueNode struct {
	Tree *parser.StringValueContext
	Str  string
}

func NewStringValueNode(ctx *Context, t *parser.StringValueContext) *StringValueNode {
	return &StringValueNode{
		Tree: t,
		Str:  t.GetText(),
	}
}

func (self *StringValueNode) Result(ctx *Context) string {
	return self.Str
}
