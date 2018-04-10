package Plan

import (
	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromRelationPrimary(ctx *Context.Context, t parser.IRelationPrimaryContext) {
	tt := t.(*parser.RelationPrimaryContext)
	if tn := t.QualifiedName(); tn != nil {
		ttn := tn.(*parser.QualifiedNameContext)
		name := ttn.GetText()
		return NewPlanScanNode(ctx, name)

	} else if tq := t.Query(); tq != nil {
		return NewPlanNodeFromQuery(tq)

	} else if tr := t.Relation(); tr != nil {
		return NewPlanNodeFromRelation(tr)
	}
}
