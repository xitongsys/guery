package plan

import (
	"fmt"

	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/gtype"
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/parser"
	"github.com/xitongsys/guery/row"
)

type NumberNode struct {
	Name      string
	DoubleVal *float64
	IntVal    *int64
}

func NewNumberNode(runtime *config.ConfigRuntime, t parser.INumberContext) *NumberNode {
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

func (self *NumberNode) Init(md *metadata.Metadata) error {
	return nil
}

func (self *NumberNode) Result(input *row.RowsGroup) (interface{}, error) {
	rn := input.GetRowsNumber()
	res := make([]interface{}, rn)
	if self.DoubleVal != nil {
		for i := 0; i < rn; i++ {
			res[i] = *self.DoubleVal
		}
	} else if self.IntVal != nil {
		for i := 0; i < rn; i++ {
			res[i] = *self.IntVal
		}
	} else {
		return nil, fmt.Errorf("wrong NumberNode")
	}
	return res, nil
}

func (self *NumberNode) GetType(md *metadata.Metadata) (gtype.Type, error) {
	if self.DoubleVal != nil {
		return gtype.FLOAT64, nil
	} else if self.IntVal != nil {
		return gtype.INT64, nil
	}
	return gtype.UNKNOWNTYPE, fmt.Errorf("wrong NumberNode")
}
