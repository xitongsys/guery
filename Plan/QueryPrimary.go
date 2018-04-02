package Plan

import (
	"github.com/xitongsys/guery/DataSource"
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromQueryPrimary(t parser.IQueryPrimaryContext) PlanNode {
	var res PlanNode
	tt := t.(*parser.QueryPrimaryContext)
	if tqs := tt.QuerySpecification(); tqs != nil {
		res = NewPlanNodeFromQuerySpecification(tqs)
	} else {
		res = NewPlanNodeFromQuery(tt.Query())
	}
	return res
}
