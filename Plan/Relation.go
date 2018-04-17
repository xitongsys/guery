package Plan

import (
	"github.com/xitongsys/guery/Common"
	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/DataSource"
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromRelation(ctx *Context.Context, t parser.IRelationContext) PlanNode {
	tt := t.(*parser.RelationContext)
	if sr := tt.SampledRelation(); sr != nil {
		return NewPlanNodeFromSampleRelation(ctx, sr)

	} else { //join
		leftRelation, rightRelation := t.GetLeftRelation(), t.GetRightRelation()
		leftNode, rightNode := NewPlanNodeFromRelation(ctx, leftRelation), NewPlanNodeFromRelation(ctx, rightRelation)
		joinText := tt.JoinType().(*parser.JoinTypeContext).GetText()
		var joinType Common.JoinType
		if joinText == "" || joinText[0:1] == "I" {
			joinType = Common.INNERJOIN
		} else if joinText[0:1] == "L" {
			joinType = Common.LEFTJOIN
		} else if joinText[0:1] == "R" {
			joinType = Common.RIGHTJOIN
		}
		joinCriteriaNode := NewJoinCriteriaNode(ctx, tt.JoinCriteria())
		return NewPlanJoinNode(ctx, leftNode, rightNode, joinType, joinCriteriaNode)
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
	columnNames := []string{}
	name := ""
	for _, ds := range dss {
		name = name + ds.Name
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
		columnNames = append(columnNames, ds.ColumnNames...)
	}

	res := DataSource.NewDataSource(name, columnNames, columnBuffers)
	for k, v := range columnMap {
		res.ColumnMap[k] = v
	}

	return NewPlanScanNodeFromDataSource(ctx, res)
}
