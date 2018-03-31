package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/DataSource"
)

type SingleStatementNode struct {
	tree      *antlr.Tree
	statement *StatementNode
	result    DataSource.DataSource
}

func (self *SingleStatementNode) Result() DataSource.DataSource {
	if self.statement != nil && self.result == nil {
		self.result = self.statement.Result()
	}
	return nil
}

type StatementNode struct {
	tree   *antlr.Tree
	query  *QueryNode
	result DataSource.DataSource
}

func (self *StatementNode) Result() DataSource.DataSource {
	if self.query != nil && self.result == nil {
		self.result = self.query.Result()
	}
	return self.result
}
