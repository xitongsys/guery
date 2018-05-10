package Plan

import (
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromRelation(t parser.IRelationContext) PlanNode {
	tt := t.(*parser.RelationContext)
	if sr := tt.SampledRelation(); sr != nil {
		return NewPlanNodeFromSampleRelation(sr)

	} else { //join
		leftRelation, rightRelation := t.GetLeftRelation(), t.GetRightRelation()
		leftNode, rightNode := NewPlanNodeFromRelation(leftRelation), NewPlanNodeFromRelation(rightRelation)
		joinText := tt.JoinType().(*parser.JoinTypeContext).GetText()
		var joinType JoinType
		if joinText == "" || joinText[0:1] == "I" {
			joinType = INNERJOIN
		} else if joinText[0:1] == "L" {
			joinType = LEFTJOIN
		} else if joinText[0:1] == "R" {
			joinType = RIGHTJOIN
		}
		joinCriteriaNode := NewJoinCriteriaNode(tt.JoinCriteria())
		res := NewPlanJoinNode(leftNode, rightNode, joinType, joinCriteriaNode)
		leftNode.SetOutput(res)
		rightNode.SetOutput(res)
		return res
	}
	return nil

}

func NewPlanNodeFromRelations(ts []parser.IRelationContext) PlanNode {
	if len(ts) == 1 {
		return NewPlanNodeFromRelation(ts[0])
	}
	res := NewPlanCombineNode([]PlanNode{})
	for _, t := range ts {
		relationNode := NewPlanNodeFromRelation(t)
		res.Inputs = append(res.Inputs, relationNode)
		relationNode.SetOutput(res)
	}
	return res
}
