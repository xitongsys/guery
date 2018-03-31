package Plan

import (
	"github.com/xitongsys/guery/Common"
	"github.com/xitongsys/guery/DataSource"
	"github.com/xitongsys/guery/parser"
)

type QueryTermNode struct {
	Tree             *parser.QueryTermContext
	QueryTermDefault *QueryTermDefaultNode
	SetOperation     *SetOperation
}

func NewQueryTermNode(ctx *Context, t *parser.QueryTermContext) *QueryTermNode {
	res := &QueryTermNode{
		Tree: t,
	}
	child := t.GetChildren()[0]
	switch child.(type) {
	case *parser.QueryTermDefaultContext:
		res.QueryTermDefault = NewQueryTermDefaultNode(ctx,
			child.(*parser.QueryTermDefaultContext))
	}
	return res
}

func (self *QueryTermNode) Result(ctx *Context) DataSource.DataSource {
	if self.QueryTermDefault != nil {
		return self.QueryTermDefault.Result(ctx)
	} else {
		return self.SetOperation.Result(ctx)
	}
}

///
type QueryTermDefaultNode struct {
	Tree         *parser.QueryTermDefaultContext
	QueryPrimary *QueryPrimaryNode
}

func NewQueryTermDefaultNode(ctx *Context, t *parser.QueryTermDefaultContext) *QueryTermDefaultNode {
	res := &QueryTermDefaultNode{
		Tree:         t,
		QueryPrimary: NewQueryPrimaryNode(ctx, t.QueryPrimary().(*parser.QueryPrimaryContext)),
	}
	return res
}

func (self *QueryTermDefaultNode) Result(ctx *Context) DataSource.DataSource {
	return self.QueryPrimary.Result(ctx)
}

///
type SetOperation struct {
	Tree           *parser.SetOperationContext
	LeftQueryTerm  *QueryTermNode
	RightQueryTerm *QueryTermNode
	Operator       *Common.Operator
}

func (self *SetOperation) Result(ctx *Context) DataSource.DataSource {
	return nil
}
