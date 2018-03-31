package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/Common"
)

type PredicatedNode struct {
	tree            *antlr.Tree
	valueExpression *ValueExpressionDefaultNode
	predicate       *PredicateNode
}

type PredicateNode struct {
	tree                     *antlr.Tree
	comparsion               *ComparisonNode
	quantifiedComparisonNode *QuantifiedComparisonNode
	between                  *BetweenNode
	inList                   *InListNode
	like                     *LikeNode
	nullPredicate            *NullPredicateNode
	distinctFrom             *DistinctFromNode
}

type ComparisonNode struct {
	tree               *antlr.Tree
	comparsionOperator Common.ComparisonOperator
	valueExpression    *ValueExpressionNode
}

type QuantifiedComparisonNode struct {
	tree                 *antlr.Tree
	comparisonOperator   Common.ComparisonOperator
	comparsionQuantifier Common.Quantifier
	query                *QueryNode
}

type BetweenNode struct {
	tree  *antlr.Tree
	not   *Common.Not
	lower *ValueExpressionNode
	upper *ValueExpressionNode
}

type InListNode struct {
	tree        *antlr.Tree
	not         *Common.Not
	expressions []*ExpressionNode
}

type InSubqueryNode struct {
	tree  *antlr.Tree
	not   *Common.Not
	query *QueryNode
}

type LikeNode struct {
	not     *Common.Not
	pattern *ValueExpressionNode
}

type NullPredicateNode struct {
	not *Common.Not
}

type DistinctFromNode struct {
	not             *Common.Not
	rightExpression *ValueExpressionNode
}
