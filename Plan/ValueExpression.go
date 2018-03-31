package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/Common"
)

type ValueExpressionNode struct {
	tree                   *antlr.Tree
	valueExpressionDefault *ValueExpressionDefaultNode
	arithmeticUnaryNode    *ArithmeticBinaryNode
	concatenation          *ConcatenationNode
}

type ValueExpressionDefaultNode struct {
	tree              *antlr.Tree
	primaryExpression *PrimaryExpressionNode
}

type ArithmeticUnaryNode struct {
	tree            *antlr.Tree
	operator        *Common.Operator
	valueExpression *ValueExpressionNode
}

type ArithmeticBinaryNode struct {
	leftExpression  *ValueExpressionNode
	rightExpression *ValueExpressionNode
	operator        *Common.Operator
}

type ConcatenationNode struct {
	leftExpression  *ValueExpressionNode
	rightExpression *ValueExpressionNode
}
