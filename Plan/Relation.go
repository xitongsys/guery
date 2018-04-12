package Plan

import (
	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/DataSource"
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromRelation(ctx *Context.Context, t parser.IRelationContext) PlanNode {
	tt := t.(*parser.RelationContext)
	if sr := tt.SampledRelation(); sr != nil {
		return NewPlanNodeFromSampleRelation(ctx, sr)

	} else {

	}
	return nil

}

func NewPlanNodeFromRelations(ctx *Context.Context, ts []parser.IRelationContext) PlanNode {
	nodes := make([]PlanNode, len(ts))
	dss := make([]*DataSource.DataSource, len(ts))
	names := []string{}
	size := 0

	for i := 0; i < len(ts); i++ {
		nodes[i] = NewPlanNodeFromRelation(ctx, ts[i])
		dss[i] = nodes[i].Execute()
		if dss[i].GetRowNum() > size {
			size = ds[i].Size()
		}
	}

	return NewPlanScanNodeFromDataSource(ctx, tb)
}
