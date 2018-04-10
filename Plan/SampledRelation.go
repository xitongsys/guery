package Plan

import (
	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromSampleRelation(ctx *Context.Context, t parser.ISampledRelationContext) PlanNode {
	tt := t.(*parser.SampledRelationContext)
	return NewPlanNodeFromRelationPrimary(ctx, tt.RelationPrimary())
}
