package Plan

import (
	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/parser"
)

type StringValueNode struct {
	Str string
}

func NewStringValueNode(ctx *Context.Context, t *parser.StringValueContext) *StringValueNode {
	return &StringValueNode{
		Str: t.GetText(),
	}
}

func (self *StringValueNode) Result() string {
	return self.Str
}
