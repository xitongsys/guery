package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/Common"
)

type GroupByNode struct {
	tree             *antlr.Tree
	setQuantifier    *Common.Quantifier
	groupingElements []*GroupingElementNode
}
