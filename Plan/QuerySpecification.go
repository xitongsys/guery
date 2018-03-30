package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/DataSoruce"
)

type QuerySpecificationNode struct {
	tree          *antlr.Tree
	setQuantifier *Quantifier
	selectItems   []*SelectItemNode
	relations     []*RelationNode
	where         *BooleanExpressionNode
	groupBy       *GroupByNode
	having        *BooleanExpressionNode
	result        DataSoruce.DataSoruce
}

func (self *QuerySpecificationNode) Result() DataSoruce.DataSoruce {
	return nil
}
