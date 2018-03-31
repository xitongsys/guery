package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type PrimaryExpressionNode struct {
	tree               *antlr.Tree
	nullLiteral        *NullLiteralNode
	numericLiteral     *NumericLiteralNode
	booleanLiteral     *BooleanLiteralNode
	stringLiteral      *StringLiteralNode
	rowConstructor     *RowConstructorNode
	subqueryExpression *SubqueryExpressionNode
}

type NullLiteralNode struct{}

type NumericLiteralNode struct {
	number *NumberNode
}

type BooleanLiteralNode struct {
	res bool
}

type StringLiteralNode struct {
	res string
}

type RowConstructorNode struct {
	expressions []*ExpressionNode
}

type SubqueryExpressionNode struct {
	query *QueryNode
}
