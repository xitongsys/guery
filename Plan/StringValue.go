package Plan

import (
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/parser"
)

type StringValueNode struct {
	Name string
	Str  string
}

func NewStringValueNode(t parser.IStringValueContext) *StringValueNode {
	s := t.GetText()
	ls := len(s)
	return &StringValueNode{
		Str:  s[1 : ls-1],
		Name: s[1 : ls-1],
	}
}

func (self *StringValueNode) Result(intput *Util.RowsBuffer) (string, error) {
	return self.Str, nil
}
