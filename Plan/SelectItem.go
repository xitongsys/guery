package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/DataSoruce"
)

type SelectItemNode struct {
	tree       *antlr.Tree
	expression *ExpressionNode
}
