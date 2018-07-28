package plan

import (
	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromQueryPrimary(runtime *Config.ConfigRuntime, t parser.IQueryPrimaryContext) PlanNode {
	var res PlanNode
	tt := t.(*parser.QueryPrimaryContext)
	if tqs := tt.QuerySpecification(); tqs != nil {
		res = NewPlanNodeFromQuerySpecification(runtime, tqs)
	} else {
		res = NewPlanNodeFromQuery(runtime, tt.Query())
	}
	return res
}
