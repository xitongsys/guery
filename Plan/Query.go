package Plan

import (
	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromQuery(ctx *Context.Context, name string, t parser.IQueryContext) PlanNode {
	tt := t.(*parser.QueryContext)
	var res PlanNode
	queryNode := NewPlanNodeFromQueryTerm(ctx, name, tt.QueryTerm())
	res = queryNode

	if tt.ORDER() != nil {
		res = NewPlanOrderByNode(ctx, res, tt.AllSortItem())
	}

	if tt.LIMIT() != nil {
		if iv := tt.INTEGER_VALUE(); iv != nil {
			res = NewPlanLimitNode(ctx, res, iv)
		} else if ia := tt.ALL(); ia != nil {
			res = NewPlanLimitNode(ctx, res, ia)
		}
	}

	return res
}
