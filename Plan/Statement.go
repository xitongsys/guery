package Plan

import (
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromSingleStatement(t parser.ISingleStatementContext) PlanNode {
	tt := t.(*parser.SingleStatementContext)
	return NewPlanNodeFromStatement(tt.Statement())
}

func NewPlanNodeFromStatement(t parser.IStatementContext) PlanNode {
	tt := t.(*parser.StatementContext)
	return NewPlanNodeFromQuery(tt.Query())
}
