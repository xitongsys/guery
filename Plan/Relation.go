package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/DataSoruce"
)

type RelationNode struct {
	sampledRelation *SampledRelationNode
	joinRelation    *JoinRelationNode
}

type JoinRelationNode struct {
	leftRelation  *RelationNode
	rightRelation *RelationNode
	joinType      *JointType
	joinCriteria  *JoinCriteriaNode
}
