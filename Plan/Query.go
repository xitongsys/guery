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
		orderNode := NewPlanOrderByNode(res, tt.AllSortItem())
		res.SetOutput(orderNode)
		res = orderNode
	}

	if tt.LIMIT() != nil {
		if iv := tt.INTEGER_VALUE(); iv != nil {
			limitNode := NewPlanLimitNode(res, iv)
			res.SetOutput(limitNode)
			res = limitNode
		} else if ia := tt.ALL(); ia != nil {
			limitNode := NewPlanLimitNode(res, ia)
			res.SetOutput(limitNode)
			res = limitNode
		}
	}
	return res
}
