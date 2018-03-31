package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/DataSource"
)

type JoinCriteriaNode struct {
	tree   *antlr.Tree
	on     *BooleanExpressionNode
	using  []*IdentifierNode
	result DataSource.DataSource
}
