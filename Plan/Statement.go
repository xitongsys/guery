package Plan

import (
	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromSingleStatement(ctx *Context.Context, t parser.ISingleStatementContext) PlanNode {
	tt := t.(*parser.SingleStatementContext)
	return NewPlanNodeFromStatement(ctx, tt.Statement())
}

func NewPlanNodeFromStatement(ctx *Context.Context, t parser.IStatementContext) PlanNode {
	tt := t.(*parser.StatementContext)
	return NewPlanNodeFromQuery(ctx, tt.Query())
}
