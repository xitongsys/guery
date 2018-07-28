package plan

import (
	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromQuery(runtime *config.ConfigRuntime, t parser.IQueryContext) PlanNode {
	tt := t.(*parser.QueryContext)
	var res PlanNode
	queryNode := NewPlanNodeFromQueryTerm(runtime, tt.QueryTerm())
	res = queryNode

	if tt.ORDER() != nil {
		orderNode := NewPlanOrderByNode(runtime, res, tt.AllSortItem())
		res.SetOutput(orderNode)
		res = orderNode
	}

	if tt.LIMIT() != nil {
		if iv := tt.INTEGER_VALUE(); iv != nil {
			limitNode := NewPlanLimitNode(runtime, res, iv)
			res.SetOutput(limitNode)
			res = limitNode
		} else if ia := tt.ALL(); ia != nil {
			limitNode := NewPlanLimitNode(runtime, res, ia)
			res.SetOutput(limitNode)
			res = limitNode
		}
	}
	return res
}
