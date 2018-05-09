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
		return NewPlanJoinNode(leftNode, rightNode, joinType, joinCriteriaNode)
	}
	return nil

}

func NewPlanNodeFromRelations(ts []parser.IRelationContext) PlanNode {
	if len(ts) == 1 {
		return NewPlanNodeFromRelation(ts[0])
	}
	res := &PlanCombineNode{
		Inputs: make([]PlanNode, 0),
	}
	for _, t := range ts {
		res.Inputs = append(res.Inputs, NewPlanNodeFromRelation(t))
	}
	return res
}
