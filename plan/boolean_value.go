package plan

import (
	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/gtype"
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/parser"
	"github.com/xitongsys/guery/row"
)

type BooleanValueNode struct {
	Name string
	Bool bool
}

func NewBooleanValueNode(runtime *config.ConfigRuntime, t parser.IBooleanValueContext) *BooleanValueNode {
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

func (self *BooleanValueNode) Init(md *metadata.Metadata) error {
	return nil
}

func (self *BooleanValueNode) Result(input *row.RowsGroup) (interface{}, error) {
	rn := input.GetRowsNumber()
	res := make([]interface{}, rn)
	for i := 0; i < rn; i++ {
		res[i] = self.Bool
	}
	return res, nil
}

func (self *BooleanValueNode) GetType(md *metadata.Metadata) (gtype.Type, error) {
	return gtype.BOOL, nil
}
