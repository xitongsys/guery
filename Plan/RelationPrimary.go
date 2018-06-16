package Plan

import (
	"github.com/xitongsys/guery/Config"
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromRelationPrimary(runtime *Config.ConfigRuntime, t parser.IRelationPrimaryContext) PlanNode {
	tt := t.(*parser.RelationPrimaryContext)
	if tn := tt.QualifiedName(); tn != nil {
		ttn := tn.(*parser.QualifiedNameContext)
		qname := ttn.GetText()
		return NewPlanScanNode(runtime, qname)

	} else if tq := tt.Query(); tq != nil {
		return NewPlanNodeFromQuery(runtime, tq)

	} else if tr := tt.Relation(); tr != nil {
		return NewPlanNodeFromRelation(runtime, tr)
	}
	return nil
}
