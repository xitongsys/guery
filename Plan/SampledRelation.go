package Plan

import (
	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromSampleRelation(ctx *Context.Context, t parser.ISampledRelationContext) PlanNode {
	tt := t.(*parser.SampledRelationContext)
	rename := "cur"
	if id := tt.Identifier(); id != nil {
		idNode := NewIdentifierNode(ctx, id)
		rename = idNode.GetText()
	}

	return NewPlanNodeFromRelationPrimary(ctx, rename, tt.RelationPrimary())
}
