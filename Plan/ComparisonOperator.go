package Plan

import (
	"github.com/xitongsys/guery/Common"
	"github.com/xitongsys/guery/parser"
)

func NewComparisonOperator(t parser.IComparisonOperatorContext) *Common.ComparisonOperator {
	tt := t.(*parser.ComparisonOperatorContext)
	var op Common.ComparisonOperator
	if tt.EQ() != nil {
		op = Common.EQ
	} else if tt.NEQ() != nil {
		op = Common.NEQ
	} else if tt.LT() != nil {
		op = Common.LT
	} else if tt.LTE() != nil {
		op = Common.LTE
	} else if tt.GT() != nil {
		op = Common.GT
	} else if tt.GTE() != nil {
		op = Common.GTE
	}
	return &op
}
