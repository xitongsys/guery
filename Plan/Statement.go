package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/DataSoruce"
)

type SingleStatementNode struct {
	tree      *antlr.Tree
	statement *StatementNode
	result    DataSource.DataSoruce
}

func (self *SingleStatementNode) Result() DataSoruce.DataSoruce {
	if self.statement != nil && self.result == nil {
		self.result = self.statement.Result()
	}
	return nil
}

type StatementNode struct {
	tree   *antlr.Tree
	query  *QueryNode
	result DataSoruce.DataSoruce
}

func (self *StatementNode) Result() DataSoruce.DataSoruce {
	if self.query != nil && self.result == nil {
		self.result = self.query.Result()
	}
	return self.result
}
