package Plan

import (
	"github.com/xitongsys/guery/Common"
	"github.com/xitongsys/guery/parser"
)

type SelectItemNode struct {
	Tree          *parser.SelectItemContext
	Expression    *ExpressionNode
	As            *Common.As
	Identifier    *IdentifierNode
	QualifiedName *QualifiedNameNode
}

func NewSelectItemNode(ctx *Context, t *parser.SelectItemContext) *SelectItemNode {
	res := &SelectItemNode{
		Tree: t,
	}
	child := t.GetChildren()[0]
	switch child.(type) {
	case *parser.ExpressionContext:
		res.Expression = NewExpressionNode(ctx, child.(*parser.ExpressionContext))
	}
	return res
}

func (self *SelectItemNode) Result(ctx *Context) interface{} {
	if self.Expression != nil {
		return self.Expression.Result(ctx)
	}
	return nil
}
