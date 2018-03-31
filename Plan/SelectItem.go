package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/DataSoruce"
)

type SelectItemNode struct {
	tree         *antlr.Tree
	selectSingle *SelectSingleNode
	selectAll    *SelectAllNode
}

type SelectSingleNode struct {
	expression *ExpressionNode
}

type SelectAllNode struct {
	qualifiedName *QualifiedNameNode
}
