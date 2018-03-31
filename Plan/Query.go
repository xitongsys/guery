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
		Tree:      t,
		QueryTerm: NewQueryTermNode(ctx, t.QueryTerm().(*parser.QueryTermContext)),
	}
	return res
}

func (self *QueryNode) Result(ctx *Context) DataSource.DataSource {
	return self.QueryTerm.Result(ctx)
}
