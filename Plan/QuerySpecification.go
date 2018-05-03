package Plan

import (
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromQuerySpecification(t parser.IQuerySpecificationContext) PlanNode {
	tt := t.(*parser.QuerySpecificationContext)
	var res PlanNode
	if rels := tt.AllRelation(); rels != nil && len(rels) > 0 {
		res = NewPlanNodeFromRelations(rels)

	}
	if wh := tt.GetWhere(); wh != nil {
		res = NewPlanFiliterNode(res, wh)
	}

	res = NewPlanSelectNode(res, tt.AllSelectItem(), tt.GroupBy())
	if having := tt.GetHaving(); having != nil {
		res = NewPlanHavingNode(res, having)
	}
	return res
}
