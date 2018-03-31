package Plan

import (
	"github.com/xitongsys/guery/DataSource"
	"github.com/xitongsys/guery/parser"
)

type QueryTermNode struct {
	Tree         *parser.QueryTermContext
	QueryPrimary *QueryPrimaryNode
}

func NewQueryTermNode(ctx *Context, t *parser.QueryTermContext) *QueryTermNode {
	res := &QueryTermNode{
		Tree: t,
	}
	children := t.GetChildren()
	if len(children) == 1 {
		res.QueryPrimary = NewQueryPrimaryNode(ctx,
			t.QueryPrimary().(*parser.QueryPrimaryContext))
	}

	return res
}

func (self *QueryTermNode) Result(ctx *Context) DataSource.DataSource {
	if self.QueryPrimary != nil {
		return self.QueryPrimary.Result(ctx)
	}
	return nil
}
