package Plan

import (
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Row"
	"github.com/xitongsys/guery/Type"
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

func (self *StringValueNode) Result(intput *Row.RowsGroup) (string, error) {
	return self.Str, nil
}

func (self *StringValueNode) GetType(md *Metadata.Metadata) (Type.Type, error) {
	return Type.STRING, nil
}
