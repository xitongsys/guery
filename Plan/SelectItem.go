package Plan

import (
	"github.com/xitongsys/guery/Common"
	"github.com/xitongsys/guery/parser"
)

type SelectItemNode struct {
	Tree         *parser.SelectItemContext
	SelectSingle *SelectSingleNode
	SelectAll    *SelectAllNode
}

func NewSelectItemNode(ctx *Context, t *parser.SelectItemContext) *SelectItemNode {
	res := &SelectItemNode{
		Tree: t,
	}
	child := t.GetChildren()[0]
	switch child.(type) {
	case *parser.SelectSingleContext:
		res.SelectSingle = NewSelectSingleNode(ctx, child.(*parser.SelectSingleContext))

	case *parser.SelectAllContext:
	}
	return res
}

func (self *SelectItemNode) Result(ctx *Context) interface{} {
	if self.SelectSingle != nil {
		self.SelectSingle.Result(ctx)
	}
	return nil
}

//SelectSingleNode
type SelectSingleNode struct {
	Tree       *parser.SelectSingleContext
	Expression *ExpressionNode
	As         *Common.As
	Identifier *IdentifierNode
}

func NewSelectSingleNode(ctx *Context, t *parser.SelectSingleContext) *SelectSingleNode {
	res := &SelectSingleNode{
		Tree: t,
	}
	child := t.GetChildren()[0]
	res.Expression = NewExpressionNode(ctx, child.(*parser.ExpressionContext))
	return res
}

func (self *SelectSingleNode) Result(ctx *Context) interface{} {
	return self.Expression.Result(ctx)
}

///
type SelectAllNode struct {
	qualifiedName *QualifiedNameNode
}
