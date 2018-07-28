package plan

import (
	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromQueryTerm(runtime *Config.ConfigRuntime, t parser.IQueryTermContext) PlanNode {
	var res PlanNode
	tt := t.(*parser.QueryTermContext)
	if tqp := tt.QueryPrimary(); tqp != nil {
		res = NewPlanNodeFromQueryPrimary(runtime, tqp)

	} else {
		left := NewPlanNodeFromQueryTerm(runtime, tt.GetLeft())
		right := NewPlanNodeFromQueryTerm(runtime, tt.GetRight())
		op := tt.GetOperator()
		unionNode := NewPlanUnionNode(runtime, left, right, op)
		left.SetOutput(unionNode)
		right.SetOutput(unionNode)
		res = unionNode
	}

	return res
}
