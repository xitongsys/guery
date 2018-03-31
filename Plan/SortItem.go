package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/Common"
)

type SortItemNode struct {
	tree         *antlr.Tree
	expression   *ExpressionNode
	ordering     *Common.Order
	nullOrdering *Common.Order
}
