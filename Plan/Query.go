package Plan

import (
	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromQuery(ctx *Context.Context, t parser.IQueryContext) PlanNode {
	var res PlanNode
	queryNode := NewPlanNodeFromQueryTerm(ctx, t.Query())
	res = queryNode

	if t.ORDER() != nil {
		res = NewPlanOrderByNode(ctx, res, t.AllSortItem())
	}

	if t.LIMIT() != nil {
		if t.INTEGER_VALUE() != nil {
			res = NewPlanLimitNode(ctx, res, t.INTEGER_VALUE())
		} else if t.ALL() != nil {
			res = NewPlanLimitNode(ctx, res, t.ALL())
		}
	}

	return res
}
