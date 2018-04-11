package Plan

import (
	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromQueryTerm(ctx *Context.Context, name string, t parser.IQueryTermContext) PlanNode {
	var res PlanNode
	tt := t.(*parser.QueryTermContext)
	if tqp := tt.QueryPrimary(); tqp != nil {
		res = NewPlanNodeFromQueryPrimary(ctx, name, tqp)

	} else {
		left := NewPlanNodeFromQueryTerm(ctx, name+"left", tt.GetLeft())
		right := NewPlanNodeFromQueryTerm(ctx, name+"right", tt.GetRight())
		op := tt.GetOperator()
		res = NewPlanUnionNode(ctx, name, left, right, op)
	}

	return res
}
