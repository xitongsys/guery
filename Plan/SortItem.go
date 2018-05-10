package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/Util"
)

type SortItemNode struct {
	tree         *antlr.Tree
	expression   *ExpressionNode
	ordering     *Util.OrderType
	nullOrdering *Util.OrderType
}
