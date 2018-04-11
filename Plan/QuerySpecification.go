package Plan

import (
	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromQuerySpecification(ctx *Context.Context, name string, t parser.IQuerySpecificationContext) PlanNode {
	tt := t.(*parser.QuerySpecificationContext)
	var res PlanNode
	if rels := tt.AllRelation(); rels != nil && len(rels) > 0 {
		res = NewPlanNodeFromRelations(ctx, rels)

	}
	if wh := tt.GetWhere(); wh != nil {
		res = NewPlanFiliterNode(ctx, res, wh)
	}
	res = NewPlanSelectNode(ctx, name, res, tt.AllSelectItem(), tt.GroupBy())
	if having := tt.GetHaving(); having != nil {
		res = NewPlanHavingNode(ctx, res, having)
	}
	return res
}
