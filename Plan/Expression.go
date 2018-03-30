package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/DataSoruce"
)

type ExpressionNode struct {
	tree              *antlr.Tree
	booleanExpression *BooleanExpressionNode
}
