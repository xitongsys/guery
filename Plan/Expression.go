package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type ExpressionNode struct {
	tree              *antlr.Tree
	booleanExpression *BooleanExpressionNode
}
