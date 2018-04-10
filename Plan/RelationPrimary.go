package Plan

import (
	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromRelationPrimary(ctx *Context.Context, t parser.IRelationPrimaryContext) PlanNode {
	tt := t.(*parser.RelationPrimaryContext)
	if tn := tt.QualifiedName(); tn != nil {
		ttn := tn.(*parser.QualifiedNameContext)
		name := ttn.GetText()
		return NewPlanScanNode(ctx, name)

	} else if tq := tt.Query(); tq != nil {
		return NewPlanNodeFromQuery(ctx, tq)

	} else if tr := tt.Relation(); tr != nil {
		return NewPlanNodeFromRelation(ctx, tr)
	}
	return nil
}
