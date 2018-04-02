package Plan

import (
	"github.com/xitongsys/guery/DataSource"
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromSingleStatement(t *parser.SingleExpressionContext) PlanNode {
	return NewPlanNodeFromStatement(t.Statement())
}

func NewPlanNodeFromStatement(t parser.IStatementContext) PlanNode {
	tt := t.(*parser.StatementContext)
	return NewPlanNodeFromQuery(tt.Query())
}
