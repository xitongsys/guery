package Plan

import (
	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/DataSource"
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromQueryPrimary(ctx *Context.Context, t parser.IQueryPrimaryContext) PlanNode {
	var res PlanNode
	tt := t.(*parser.QueryPrimaryContext)
	if tqs := tt.QuerySpecification(); tqs != nil {
		res = NewPlanNodeFromQuerySpecification(ctx, tqs)
	} else {
		res = NewPlanNodeFromQuery(ctx, tt.Query())
	}
	return res
}
