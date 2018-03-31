package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/DataSource"
)

type QueryPrimaryNode struct {
	tree               *antlr.Tree
	querySpecification *QuerySpecificationNode
	subQuery           *QueryNode
	result             DataSource.DataSource
}

func (self *QueryPrimaryNode) Result() DataSource.DataSource {
	if self.result == nil {
		if self.querySpecification != nil {
			self.result = self.querySpecification.Result()
		} else {
			self.result = self.subQuery.Result()
		}
	}
	return self.result
}
