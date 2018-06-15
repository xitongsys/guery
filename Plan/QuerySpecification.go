package Plan

import (
	"github.com/xitongsys/guery/Config"
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromQuerySpecification(runtime *Config.ConfigRuntime, t parser.IQuerySpecificationContext) PlanNode {
	tt := t.(*parser.QuerySpecificationContext)
	var res PlanNode
	if rels := tt.AllRelation(); rels != nil && len(rels) > 0 {
		res = NewPlanNodeFromRelations(runtime, rels)

	}
	if wh := tt.GetWhere(); wh != nil {
		filiterNode := NewPlanFiliterNode(runtime, res, wh)
		res.SetOutput(filiterNode)
		res = filiterNode
	}

	if gb := tt.GroupBy(); gb != nil {
		groupByNode := NewPlanGroupByNode(runtime, res, gb, tt.GetHaving())
		res.SetOutput(groupByNode)
		res = groupByNode
	}

	selectNode := NewPlanSelectNode(runtime, res, tt.AllSelectItem())

	//for select sum/min/count/... without group
	if selectNode.IsAggregate && res.GetNodeType() != GROUPBYNODE {
		aggNode := NewPlanAggregateNode(runtime, res)
		aggNode.SetOutput(selectNode)
		selectNode.SetInputs([]PlanNode{aggNode})
		res.SetOutput(aggNode)

	} else {
		res.SetOutput(selectNode)
	}

	res = selectNode
	return res
}
