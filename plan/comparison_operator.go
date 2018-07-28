package plan

import (
	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/gtype"
	"github.com/xitongsys/guery/parser"
)

func NewComparisonOperator(runtime *config.ConfigRuntime, t parser.IComparisonOperatorContext) *gtype.Operator {
	tt := t.(*parser.ComparisonOperatorContext)
	var op gtype.Operator
	if tt.EQ() != nil {
		op = gtype.EQ
	} else if tt.NEQ() != nil {
		op = gtype.NEQ
	} else if tt.LT() != nil {
		op = gtype.LT
	} else if tt.LTE() != nil {
		op = gtype.LTE
	} else if tt.GT() != nil {
		op = gtype.GT
	} else if tt.GTE() != nil {
		op = gtype.GTE
	}
	return &op
}
