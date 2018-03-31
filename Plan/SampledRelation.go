package Plan

import (
	"github.com/xitongsys/guery/Common"
)

type SampledRelationNode struct {
	relationPrimary *RelationPrimaryNode
	as              *Common.As
	identifier      *IdentifierNode
	columnAliases   *ColumnAliasesNode
}
