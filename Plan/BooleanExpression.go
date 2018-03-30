package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/DataSoruce"
)

type BooleanExpressionNode struct {
	tree           *antlr.Tree
	booleanDefault *BooleanDefaultNode
	logicalNot     *LocicalNotNode
	logicalBinary  *LogicalBinaryNode
}

type BooleanDefaultNode struct {
	predicated *PredicatedNode
}

type LocicalNotNode struct {
	booleanExpression *BooleanExpressionNode
}

type LogicalBinaryNode struct {
	leftExpression  *BooleanExpressionNode
	rightExpression *BooleanExpressionNode
	operator        *Operator
}
