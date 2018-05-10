package Plan

import (
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/parser"
)

func NewComparisonOperator(t parser.IComparisonOperatorContext) *Util.Operator {
	tt := t.(*parser.ComparisonOperatorContext)
	var op Util.Operator
	if tt.EQ() != nil {
		op = Util.EQ
	} else if tt.NEQ() != nil {
		op = Util.NEQ
	} else if tt.LT() != nil {
		op = Util.LT
	} else if tt.LTE() != nil {
		op = Util.LTE
	} else if tt.GT() != nil {
		op = Util.GT
	} else if tt.GTE() != nil {
		op = Util.GTE
	}
	return &op
}
