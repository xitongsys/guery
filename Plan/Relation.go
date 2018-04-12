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
	size := 0

	for i := 0; i < len(ts); i++ {
		nodes[i] = NewPlanNodeFromRelation(ctx, ts[i])
		dss[i] = nodes[i].Execute()
		if dss[i].GetRowNum() > size {
			size = dss[i].GetRowNum()
		}
	}

	columnBuffers := []DataSource.ColumnBuffer{}
	columnMap := make(map[string]int)
	for _, ds := range dss {
		for name, index := range ds.ColumnMap {
			columnMap[name] = index + len(columnBuffers)
		}

		for _, buf := range ds.ColumnBuffers {
			memBuf := DataSource.NewMemColumnBuffer()
			for j := 0; j < buf.Size(); j++ {
				memBuf.Append(buf.Read())
			}
			columnBuffers = append(columnBuffers, memBuf)
		}
	}

	res := &DataSource.DataSource{
		Names:         []string{},
		ColumnMap:     columnMap,
		ColumnBuffers: columnBuffers,
		Vals:          []interface{}{},
		CurIndex:      -1,
		RowNum:        size,
	}

	return NewPlanScanNodeFromDataSource(ctx, res)
}
