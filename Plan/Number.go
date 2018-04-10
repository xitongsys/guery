package Plan

import (
	"fmt"

	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/DataSource"
	"github.com/xitongsys/guery/parser"
)

type NumberNode struct {
	DoubleVal *float64
	IntVal    *int64
}

func NewNumberNode(ctx *Context.Context, t parser.INumberContext) *NumberNode {
	tt := t.(*parser.NumberContext)
	res := &NumberNode{}
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

func (self *NumberNode) Result(input DataSource.DataSource) interface{} {
	if self.DoubleVal != nil {
		return *self.DoubleVal
	} else {
		return *self.IntVal
	}
}
