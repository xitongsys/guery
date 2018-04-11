package Plan

import (
	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromQueryPrimary(ctx *Context.Context, name string, t parser.IQueryPrimaryContext) PlanNode {
	var res PlanNode
	tt := t.(*parser.QueryPrimaryContext)
	if tqs := tt.QuerySpecification(); tqs != nil {
		res = NewPlanNodeFromQuerySpecification(ctx, name, tqs)
	} else {
		res = NewPlanNodeFromQuery(ctx, name, tt.Query())
	}
	return res
}
