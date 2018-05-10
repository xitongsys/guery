package Plan

import (
	"fmt"

	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/parser"
)

type NumberNode struct {
	Name      string
	DoubleVal *float64
	IntVal    *int64
}

func NewNumberNode(t parser.INumberContext) *NumberNode {
	tt := t.(*parser.NumberContext)
	res := &NumberNode{}
	res.Name = tt.GetText()
	if dv := tt.DOUBLE_VALUE(); dv != nil {
		var v float64
		fmt.Sscanf(dv.GetText(), "%f", &v)
		res.DoubleVal = &v

	} else if iv := tt.INTEGER_VALUE(); iv != nil {
		var v int64
		fmt.Sscanf(iv.GetText(), "%d", &v)
		res.IntVal = &v
	}
	return res
}

func (self *NumberNode) Result(input *Util.RowsBuffer) (interface{}, error) {
	if self.DoubleVal != nil {
		return *self.DoubleVal, nil
	} else if self.IntVal != nil {
		return *self.IntVal, nil
	}
	return nil, fmt.Errorf("wrong NumberNode")
}

func (self *NumberNode) GetType(md *Util.Metadata) (Util.Metadata, error) {
	if self.DoubleVal != nil {
		return Util.DOUBLE, nil
	} else if self.IntVal != nil {
		return Util.INT, nil
	}
	return Util.UNKNOWNTYPE, fmt.Errorf("wrong NumberNode")
}
