package Plan

import (
	"github.com/xitongsys/guery/DataSource"
	"github.com/xitongsys/guery/parser"
)

type QueryNode struct {
	Tree      *parser.QueryContext
	QueryTerm *QueryTermNode
	OrderBy   []*SortItemNode
	Limit     string
}

func NewQueryNode(ctx *Context, t *parser.QueryContext) *QueryNode {
	res := &QueryNode{
		Tree: t,
	}
	children := t.GetChildren()
	for i := 0; i < len(children); i++ {
		child := children[i]
		switch child.(type) {
		case *parser.QueryTermContext:
			res.QueryTerm = NewQueryTermNode(ctx,
				child.(*parser.QueryTermContext))

		case *parser.SortItemContext:
		}
	}
	return res
}

func (self *QueryNode) Result(ctx *Context) DataSource.DataSource {
	return self.QueryTerm.Result(ctx)
}
