package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type PredicatedNode struct {
	valueExpression *ValueExpressionDefaultNode
	predicate       *PredicateNode
}

type PredicateNode struct {
	comparsion               *ComparisonNode
	quantifiedComparisonNode *QuantifiedComparisonNode
	between                  *BetweenNode
	inList                   *InListNode
	like                     *LikeNode
	nullPredicate            *NullPredicateNode
	distinctFrom             *DistinctFromNode
}

type ComparisonNode struct {
	comparsionOperator ComparisonOperator
	valueExpression    *ValueExpressionNode
}

type QuantifiedComparisonNode struct {
	comparisonOperator   ComparisonOperator
	comparsionQuantifier ComparisonQuantifier
	query                *QueryNode
}

type BetweenNode struct {
	not   *Not
	lower *ValueExpressionNode
	upper *ValueExpressionNode
}

type InListNode struct {
	not         *Not
	expressions []*ExpressionNode
}

type InSubqueryNode struct {
	not   *Not
	query *QueryNode
}

type LikeNode struct {
	not     *Not
	pattern *ValueExpressionNode
}

type NullPredicateNode struct {
	not *Not
}

type DistinctFromNode struct {
	not             *Not
	rightExpression *ValueExpressionNode
}
