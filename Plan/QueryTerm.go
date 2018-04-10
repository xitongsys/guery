package Plan

import (
	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromQueryTerm(ctx *Context.Context, t parser.IQueryTermContext) PlanNode {
	var res PlanNode
	tt := t.(*parser.QueryTermContext)
	if tqp := tt.QueryPrimary(); tqp != nil {
		res = NewPlanNodeFromQueryPrimary(ctx, tqp)

	} else {
		left := NewPlanNodeFromQueryTerm(ctx, tt.GetLeft())
		right := NewPlanNodeFromQueryTerm(ctx, tt.GetRight())
		op := tt.GetOperator()
		res = NewPlanUnionNode(ctx, left, right, op)
	}

	return res
}
