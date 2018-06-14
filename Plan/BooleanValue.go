package Plan

import (
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Row"
	"github.com/xitongsys/guery/Type"
	"github.com/xitongsys/guery/parser"
)

type BooleanValueNode struct {
	Name string
	Bool bool
}

func NewBooleanValueNode(t parser.IBooleanValueContext) *BooleanValueNode {
	s := t.GetText()
	b := true
	if s != "TRUE" {
		b = false
	}
	return &BooleanValueNode{
		Bool: b,
		Name: s,
	}
}

func (self *BooleanValueNode) Result(intput *Row.RowsGroup) (bool, error) {
	return self.Bool, nil
}

func (self *BooleanValueNode) GetType(md *Metadata.Metadata) (Type.Type, error) {
	return Type.BOOL, nil
}
