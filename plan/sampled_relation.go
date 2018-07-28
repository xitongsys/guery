package plan

import (
	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromSampleRelation(runtime *config.ConfigRuntime, t parser.ISampledRelationContext) PlanNode {
	tt := t.(*parser.SampledRelationContext)
	res := NewPlanNodeFromRelationPrimary(runtime, tt.RelationPrimary())
	if id := tt.Identifier(); id != nil {
		idNode := NewIdentifierNode(runtime, id)
		rename := idNode.GetText()
		renameNode := NewPlanRenameNode(runtime, res, rename)
		res.SetOutput(renameNode)
		res = renameNode
	}
	return res
}
