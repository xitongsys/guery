package Plan

import (
	"github.com/xitongsys/guery/Util"
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

func (self *BooleanValueNode) Result(intput *Util.RowsGroup) (bool, error) {
	return self.Bool, nil
}

func (self *BooleanValueNode) GetType(md *Util.Metadata) (Util.Type, error) {
	return Util.BOOL, nil
}
