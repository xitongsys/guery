package Plan

import (
	"fmt"

	"github.com/xitongsys/guery/Common"
	"github.com/xitongsys/guery/DataSource"
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromQuerySpecification(t parser.IQuerySpecificationContext) PlanNode {
	tt := t.(*parser.QuerySpecificationContext)
	var res PlanNode
	if rels := tt.AllRelation(); rels != nil && len(rels) > 0 {
		res = NewPlanNodeFromRelations(rels)
	}
	if be := tt.Where(); be != nil {
		res = NewPlanFiliterNode(res, be)
	}
	res = NewPlanSelectNode(res, tt.AllSelectItem(), tt.GroupBy())
	if having := tt.Having(); having != nil {
		res = NewPlanHavingNode(res, having)
	}
	return res
}
