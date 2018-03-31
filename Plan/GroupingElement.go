package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type GroupingElementNode struct {
	tree                *antlr.Tree
	groupingExpressions *GroupingExpressionsNode
}
