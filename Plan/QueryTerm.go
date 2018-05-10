package Plan

import (
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
		unionNode := NewPlanUnionNode(left, right, op)
		left.SetOutput(unionNode)
		right.SetOutput(unionNode)
		res = unionNode
	}

	return res
}
