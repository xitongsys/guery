package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/DataSoruce"
)

type GroupingElementNode struct {
	tree                *antlr.Tree
	groupingExpressions *GroupingExpressionsNode
}
