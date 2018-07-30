package plan

import (
	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromQuerySpecification(runtime *config.ConfigRuntime, t parser.IQuerySpecificationContext) PlanNode {
	tt := t.(*parser.QuerySpecificationContext)
	var res PlanNode
	if rels := tt.AllRelation(); rels != nil && len(rels) > 0 {
		res = NewPlanNodeFromRelations(runtime, rels)

	}
	if wh := tt.GetWhere(); wh != nil {
		filterNode := NewPlanFilterNode(runtime, res, wh)
		res.SetOutput(filterNode)
		res = filterNode
	}

	if gb := tt.GroupBy(); gb != nil {
		groupByNode := NewPlanGroupByNode(runtime, res, gb)
		res.SetOutput(groupByNode)
		res = groupByNode
	}

	selectNode := NewPlanSelectNode(runtime, res, tt.SetQuantifier(), tt.AllSelectItem(), tt.GetHaving())

	res.SetOutput(selectNode)

	res = selectNode
	return res
}
