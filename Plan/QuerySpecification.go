package Plan

import (
	"fmt"

	"github.com/xitongsys/guery/Common"
	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/DataSource"
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromQuerySpecification(ctx *Context.Context, t parser.IQuerySpecificationContext) PlanNode {
	tt := t.(*parser.QuerySpecificationContext)
	var res PlanNode
	if rels := tt.AllRelation(); rels != nil && len(rels) > 0 {
		res = NewPlanNodeFromRelations(ctx, rels)
	}
	if be := tt.Where(); be != nil {
		res = NewPlanFiliterNode(ctx, res, be)
	}
	res = NewPlanSelectNode(ctx, res, tt.AllSelectItem(), tt.GroupBy())
	if having := tt.Having(); having != nil {
		res = NewPlanHavingNode(ctx, res, having)
	}
	return res
}
