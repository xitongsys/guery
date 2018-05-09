package Plan

import (
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromRelationPrimary(t parser.IRelationPrimaryContext) PlanNode {
	tt := t.(*parser.RelationPrimaryContext)
	if tn := tt.QualifiedName(); tn != nil {
		ttn := tn.(*parser.QualifiedNameContext)
		qname := ttn.GetText()
		return NewPlanScanNode(qname)

	} else if tq := tt.Query(); tq != nil {
		return NewPlanNodeFromQuery(tq)

	} else if tr := tt.Relation(); tr != nil {
		return NewPlanNodeFromRelation(tr)
	}
	return nil
}
