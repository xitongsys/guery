package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/Common"
)

type RelationNode struct {
	tree            *antlr.Tree
	sampledRelation *SampledRelationNode
	joinRelation    *JoinRelationNode
}

type JoinRelationNode struct {
	leftRelation  *RelationNode
	rightRelation *RelationNode
	joinType      *Common.JoinType
	joinCriteria  *JoinCriteriaNode
}
