package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/DataSoruce"
)

type QueryPrimaryNode struct {
	tree               *antlr.Tree
	querySpecification *QuerySpecificationNode
	subQuery           *Query
	result             DataSoruce.DataSoruce
}

func (self *QueryPrimaryNode) Result() DataSoruce.DataSoruce {
	if self.result == nil {
		if self.querySpecification != nil {
			self.result = self.querySpecification.Result()
		} else {
			self.result = self.subQuery.Result()
		}
	}
	return self.result
}
