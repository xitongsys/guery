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
	return &StringValueNode{
		Str:  t.GetText(),
		Name: t.GetText(),
	}
}

func (self *StringValueNode) Result(intput *DataSource.DataSource) string {
	return self.Str
}
