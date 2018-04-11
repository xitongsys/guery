package Plan

import (
	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromRelationPrimary(ctx *Context.Context, name string, t parser.IRelationPrimaryContext) PlanNode {
	tt := t.(*parser.RelationPrimaryContext)
	if tn := tt.QualifiedName(); tn != nil {
		ttn := tn.(*parser.QualifiedNameContext)
		qname := ttn.GetText()
		if name != "" {
			ctx.AddTableRename(name, qname)
		}
		return NewPlanScanNode(ctx, qname)

	} else if tq := tt.Query(); tq != nil {
		return NewPlanNodeFromQuery(ctx, name, tq)

	} else if tr := tt.Relation(); tr != nil {
		return NewPlanNodeFromRelation(ctx, tr)
	}
	return nil
}
