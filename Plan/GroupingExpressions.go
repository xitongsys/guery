package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type GroupingExpressionsNode struct {
	tree        *antlr.Tree
	expressions []*ExpressionNode
}
