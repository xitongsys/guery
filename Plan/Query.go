package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/DataSoruce"
)

type QueryNode struct {
	tree      *antlr.Tree
	queryTerm *QueryTermNode
	orderBy   []*SortItemNode
	limit     string
	result    DataSoruce.DataSoruce
}

func (self *QueryNode) Result() DataSoruce.DataSoruce {
	return nil
}
