package Plan

import (
	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/DataSource"
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromSingleStatement(ctx *Context.Context, t *parser.SingleExpressionContext) PlanNode {
	return NewPlanNodeFromStatement(ctx, t.Statement())
}

func NewPlanNodeFromStatement(ctx *Context.Context, t parser.IStatementContext) PlanNode {
	tt := t.(*parser.StatementContext)
	return NewPlanNodeFromQuery(ctx, tt.Query())
}
