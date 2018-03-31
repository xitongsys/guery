package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/Common"
	"github.com/xitongsys/guery/DataSource"
)

type QueryTermNode struct {
	tree             *antlr.Tree
	queryTermDefault *QueryTermDefaultNode
	setOperation     *SetOperation
	result           DataSource.DataSource
}

func (self *QueryTermNode) Result() DataSource.DataSource {
	if self.result == nil {
		if self.queryTermDefault != nil {
			self.result = self.queryTermDefault.Result()
		} else {
			self.result = self.setOperation.Result()
		}
	}
	return self.result
}

type QueryTermDefaultNode struct {
	tree         *antlr.Tree
	queryPrimary *QueryPrimaryNode
	result       DataSource.DataSource
}

func (self *QueryTermDefaultNode) Result() DataSource.DataSource {
	return nil
}

type SetOperation struct {
	tree           *antlr.Tree
	leftQueryTerm  *QueryTermNode
	rightQueryTerm *QueryTermNode
	operator       *Common.Operator
	result         DataSource.DataSource
}

func (self *SetOperation) Result() DataSource.DataSource {
	return self.result
}
