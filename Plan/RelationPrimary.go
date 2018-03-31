package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type RelationPrimaryNode struct {
	tree                  *antlr.Tree
	tableName             *TableNameNode
	subqueryRelation      *SubqueryRelationNode
	unnest                *UnnestNode
	parenthesizedRelation *ParenthesizedRelationNode
}

type TableNameNode struct {
	qualifiedName *QualifiedNameNode
}

type SubqueryRelationNode struct {
	query *QueryNode
}

type UnnestNode struct {
	expressions []*ExpressionNode
}

type ParenthesizedRelationNode struct {
	relation *RelationNode
}
