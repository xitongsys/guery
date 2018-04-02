package Plan

import (
	"github.com/xitongsys/guery/DataSource"
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromSingleStatement(t *parser.SingleExpressionContext) PlanNode {
	child := t.GetChildren()[0]
	return NewPlanNodeFromStatement(child.(*parser.StatementContext))
}

func NewPlanNodeFromStatement(t *parser.StatementContext) PlanNode {
	child := t.GetChildren()[0]
	return NewPlanNodeFromQuery(child.(*parser.QueryContext))
}
