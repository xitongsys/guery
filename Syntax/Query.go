package Plan

import (
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromQuery(t parser.IQueryContext) PlanNode {
	tt := t.(*parser.QueryContext)
	var res PlanNode
	queryNode := NewPlanNodeFromQueryTerm(tt.QueryTerm())
	res = queryNode

	if tt.ORDER() != nil {
		res = NewPlanOrderByNode(res, tt.AllSortItem())
	}

	if tt.LIMIT() != nil {
		if iv := tt.INTEGER_VALUE(); iv != nil {
			res = NewPlanLimitNode(res, iv)
		} else if ia := tt.ALL(); ia != nil {
			res = NewPlanLimitNode(res, ia)
		}
	}

	return res
}
