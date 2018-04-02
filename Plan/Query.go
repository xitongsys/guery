package Plan

import (
	"github.com/xitongsys/guery/DataSource"
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromQuery(t parser.IQueryContext) PlanNode {
	var res PlanNode
	queryNode := NewPlanNodeFromQueryTerm(t.Query())
	res = queryNode

	if t.ORDER() != nil {
		res = NewPlanOrderByNode(res, t.AllSortItem())
	}

	if t.LIMIT() != nil {
		if t.INTEGER_VALUE() != nil {
			res = NewPlanLimitNode(res, t.INTEGER_VALUE())
		} else if t.ALL() != nil {
			res = NewPlanLimitNode(res, t.ALL())
		}
	}

	return res
}
