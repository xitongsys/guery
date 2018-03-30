package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/DataSoruce"
)

type GroupByNode struct {
	tree             *antlr.Tree
	setQuantifier    *Quantifier
	groupingElements []*GroupingElementNode
}
