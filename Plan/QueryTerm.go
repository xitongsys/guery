package Plan

import (
	"github.com/xitongsys/guery/DataSource"
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromQueryTerm(t parser.IQueryTermContext) PlanNode {
	var res PlanNode
	tt := t.(*parser.QueryTermContext)
	if tt.QueryPrimary() != nil {
		res = NewPlanNodeFromQueryPrimary(tt.QueryPrimary())

	} else {
		var op Common.Operator
		if tt.INTERSECT() != nil {
			op = Common.INTERSECT
		} else if tt.UNION() != nil {
			op = Common.UNION
		} else if tt.EXCEPT() != nil {
			op = Common.EXCEPT
		}

		left = NewPlanNodeFromQueryTerm(t.GetLeft())
		right = NewPlanNodeFromQueryTerm(t.GetRight())

		res = NewPlanUnionNode(left, right, op)
	}

	return res
}
