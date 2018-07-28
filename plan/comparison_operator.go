package plan

import (
	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/gtype"
	"github.com/xitongsys/guery/parser"
)

func NewComparisonOperator(runtime *Config.ConfigRuntime, t parser.IComparisonOperatorContext) *Type.Operator {
	tt := t.(*parser.ComparisonOperatorContext)
	var op Type.Operator
	if tt.EQ() != nil {
		op = Type.EQ
	} else if tt.NEQ() != nil {
		op = Type.NEQ
	} else if tt.LT() != nil {
		op = Type.LT
	} else if tt.LTE() != nil {
		op = Type.LTE
	} else if tt.GT() != nil {
		op = Type.GT
	} else if tt.GTE() != nil {
		op = Type.GTE
	}
	return &op
}
