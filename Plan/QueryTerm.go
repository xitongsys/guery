package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/DataSoruce"
)

type QueryTermNode struct {
	tree             *antlr.Tree
	queryTermDefault *QueryTermDefaultNode
	setOperation     *SetOperation
	result           DataSoruce.DataSoruce
}

func (self *QueryTermNode) Result() DataSoruce.DataSoruce {
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
	result       DataSoruce.DataSoruce
}

func (self *QueryTermDefaultNode) Result() DataSoruce.DataSoruce {
	return nil
}

type SetOperation struct {
	tree           *antlr.Tree
	leftQueryTerm  *QueryTermNode
	rightQueryTerm *QueryTermNode
	operator       *Operator
	result         DataSoruce.DataSoruce
}

func (self *SetOperation) Result() DataSoruce.DataSoruce {
	return self.result
}
