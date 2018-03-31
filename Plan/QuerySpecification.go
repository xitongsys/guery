package Plan

import (
	"github.com/xitongsys/guery/Common"
	"github.com/xitongsys/guery/DataSource"
	"github.com/xitongsys/guery/parser"
)

type QuerySpecificationNode struct {
	tree          *parser.QuerySpecificationContext
	setQuantifier *Common.Quantifier
	selectItems   []*SelectItemNode
	relations     []*RelationNode
	where         *BooleanExpressionNode
	groupBy       *GroupByNode
	having        *BooleanExpressionNode
	result        DataSource.DataSource
}

func (self *QuerySpecificationNode) Result() DataSource.DataSource {
	return nil
}
