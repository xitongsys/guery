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
		filiterNode := NewPlanFiliterNode(res, wh)
		res.SetOutput(filiterNode)
		res = filiterNode
	}

	if gb := tt.GroupBy(); gb != nil {
		groupByNode := NewPlanGroupByNode(res, gb)
		res.SetOutput(groupByNode)
		res = groupByNode
	}

	selectNode := NewPlanSelectNode(res, tt.AllSelectItem())
	res.SetOutput(selectNode)
	res = selectNode

	if having := tt.GetHaving(); having != nil {
		havingNode := NewPlanHavingNode(res, having)
		res.SetOutput(having)
		res = havingNode
	}
	return res
}
