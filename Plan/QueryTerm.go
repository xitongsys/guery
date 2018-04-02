package Plan

import (
	"github.com/xitongsys/guery/DataSource"
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromQueryTerm(t parser.IQueryTermContext) PlanNode {
	var res PlanNode
	tt := t.(*parser.QueryTermContext)
	if tqp := tt.QueryPrimary(); tqp != nil {
		res = NewPlanNodeFromQueryPrimary(tqp)

	} else {
		left := NewPlanNodeFromQueryTerm(tt.GetLeft())
		right := NewPlanNodeFromQueryTerm(tt.GetRight())
		op := tt.GetOperator()
		res = NewPlanUnionNode(left, right, op)
	}

	return res
}
