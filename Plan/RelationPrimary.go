package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func NewPlanNodeFromRelationPrimary(t parser.IRelationPrimaryContext) {
	tt := t.(*parser.RelationPrimaryContext)
	if tn := t.QualifiedName(); tn != nil {
		ttn := tn.(*parser.QualifiedNameContext)
		name := ttn.GetText()
		return NewPlanScanNode()

	} else if tq := t.Query(); tq != nil {
		return NewPlanNodeFromQuery(tq)

	} else if tr := t.Relation(); tr != nil {
		return NewPlanNodeFromRelation(tr)
	}

}
