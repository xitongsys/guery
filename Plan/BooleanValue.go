package Plan

import (
	"github.com/xitongsys/guery/Config"
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Row"
	"github.com/xitongsys/guery/Type"
	"github.com/xitongsys/guery/parser"
)

type BooleanValueNode struct {
	Name string
	Bool bool
}

func NewBooleanValueNode(runtime *Config.ConfigRuntime, t parser.IBooleanValueContext) *BooleanValueNode {
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

func (self *BooleanValueNode) Init(md *Metadata.Metadata) error {
	return nil
}

func (self *BooleanValueNode) Result(input *Row.RowsGroup) (interface{}, error) {
	rn := input.GetRowsNumber()
	res := make([]interface{}, rn)
	for i := 0; i < rn; i++ {
		res[i] = self.Bool
	}
	return res, nil
}

func (self *BooleanValueNode) GetType(md *Metadata.Metadata) (Type.Type, error) {
	return Type.BOOL, nil
}
