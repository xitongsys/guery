package Plan

import (
	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/DataSource"
	"github.com/xitongsys/guery/parser"
)

type StringValueNode struct {
	Name string
	Str  string
}

func NewStringValueNode(ctx *Context.Context, t parser.IStringValueContext) *StringValueNode {
	s := t.GetText()
	ls := len(s)
	return &StringValueNode{
		Str:  s[1 : ls-1],
		Name: s[1 : ls-1],
	}
}

func (self *StringValueNode) Result(intput *DataSource.DataSource) string {
	return self.Str
}
