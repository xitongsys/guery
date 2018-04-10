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
	res := &NumberNode{}
	if t.DOUBLE_VALUE() != nil {
		var v float64
		fmt.Sscanf(t.DOUBLE_VALUE().GetText(), "%f", &v)
		res.DoubleVal = &v

	} else if t.INTEGER_VALUE() != nil {
		var v int64
		fmt.Sscanf(t.INTEGER_VALUE().GetText(), "%d", &v)
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
