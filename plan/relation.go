package plan

import (
	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromRelation(runtime *config.ConfigRuntime, t parser.IRelationContext) PlanNode {
	tt := t.(*parser.RelationContext)
	if sr := tt.SampledRelation(); sr != nil {
		return NewPlanNodeFromSampleRelation(runtime, sr)

	} else { //join
		leftRelation, rightRelation := t.GetLeftRelation(), t.GetRightRelation()
		leftNode, rightNode := NewPlanNodeFromRelation(runtime, leftRelation), NewPlanNodeFromRelation(runtime, rightRelation)
		joinText := tt.JoinType().(*parser.JoinTypeContext).GetText()
		var joinType JoinType
		if joinText == "" || joinText[0:1] == "I" {
			joinType = INNERJOIN
		} else if joinText[0:1] == "L" {
			joinType = LEFTJOIN
		} else if joinText[0:1] == "R" {
			joinType = RIGHTJOIN
		}
		joinCriteriaNode := NewJoinCriteriaNode(runtime, tt.JoinCriteria())
		res := NewPlanJoinNode(runtime, leftNode, rightNode, joinType, joinCriteriaNode)
		leftNode.SetOutput(res)
		rightNode.SetOutput(res)
		return res
	}
	return nil

}

func NewPlanNodeFromRelations(runtime *config.ConfigRuntime, ts []parser.IRelationContext) PlanNode {
	if len(ts) == 1 {
		return NewPlanNodeFromRelation(runtime, ts[0])
	}
	res := NewPlanCombineNode(runtime, []PlanNode{})
	for _, t := range ts {
		relationNode := NewPlanNodeFromRelation(runtime, t)
		res.Inputs = append(res.Inputs, relationNode)
		relationNode.SetOutput(res)
	}
	return res
}
