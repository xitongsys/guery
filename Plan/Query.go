package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/DataSource"
)

type QueryNode struct {
	tree      *antlr.Tree
	queryTerm *QueryTermNode
	orderBy   []*SortItemNode
	limit     string
	result    DataSource.DataSource
}

func (self *QueryNode) Result() DataSource.DataSource {
	return nil
}
