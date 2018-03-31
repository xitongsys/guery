package Plan

import (
	"github.com/xitongsys/guery/DataSource"
	"github.com/xitongsys/guery/parser"
)

type SingleStatementNode struct {
	Tree      *parser.SingleStatementContext
	Statement *StatementDefaultNode
}

func NewSingleStatementNode(ctx *Context, t *parser.SingleStatementContext) *SingleStatementNode {
	res := &SingleStatementNode{
		Tree: t,
	}
	child := t.GetChildren()[0]
	res.Statement = NewStatementDefaultNode(ctx,
		child.(*parser.StatementDefaultContext))
	return res
}

func (self *SingleStatementNode) Result(ctx *Context) DataSource.DataSource {
	return self.Statement.Result(ctx)
}

/////
type StatementDefaultNode struct {
	Tree  *parser.StatementDefaultContext
	Query *QueryNode
}

func NewStatementDefaultNode(ctx *Context, t *parser.StatementDefaultContext) *StatementDefaultNode {
	res := &StatementDefaultNode{
		Tree: t,
		Query: NewQueryNode(ctx,
			t.Query().(*parser.QueryContext)),
	}
	return res
}

func (self *StatementDefaultNode) Result(ctx *Context) DataSource.DataSource {
	return self.Query.Result(ctx)
}

/////
type StatementNode struct {
	Tree  *parser.StatementContext
	Query *QueryNode
}

func NewStatementNode(ctx *Context, t *parser.StatementContext) *StatementNode {
	res := &StatementNode{
		Tree: t,
	}
	child := t.GetChildren()[0]
	res.Query = NewQueryNode(ctx,
		child.(*parser.StatementDefaultContext).Query().(*parser.QueryContext))
	return res
}

func (self *StatementNode) Result(ctx *Context) DataSource.DataSource {
	return self.Query.Result(ctx)
}
