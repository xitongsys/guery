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
	ds := make([]DataSource.DataSource, len(ts))
	names := []string{}
	size := int64(0)

	for i := 0; i < len(ts); i++ {
		nodes[i] = NewPlanNodeFromRelation(ctx, ts[i])
		ds[i] = nodes[i].Execute()
		names = append(names, ds[i].Names()...)
		if ds[i].Size() > size {
			size = ds[i].Size()
		}
	}

	tb := DataSource.NewTableSource("", names)
	for i := int64(0); i < size; i++ {
		vals := []interface{}{}

		for j := 0; j < len(ds); j++ {
			vals = append(vals, ds[j].ReadRow()...)
		}
		tb.Append(vals)
	}

	return NewPlanScanNodeFromDataSource(ctx, tb)
}
