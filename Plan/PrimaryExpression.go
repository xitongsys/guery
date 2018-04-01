package Plan

import (
	"github.com/xitongsys/guery/parser"
)

type PrimaryExpressionNode struct {
	Tree *parser.PrimaryExpressionContext
	//	Null         *NullNode
	Number                  *NumberNode
	BooleanValue            *BooleanValueNode
	StringValue             *StringValueNode
	Identifier              *IdentifierNode
	ParenthesizedExpression *ExpressionNode
}

func NewPrimaryExpressionNode(ctx *Context, t *parser.PrimaryExpressionContext) *PrimaryExpressionNode {
	res := &PrimaryExpressionNode{
		Tree: t,
	}
	children := t.GetChildren()
	if len(children) == 1 {
		child := children[0]
		switch child.(type) {
		case *parser.NumberContext:
			res.Number = NewNumberNode(ctx, child.(*parser.NumberContext))
		case *parser.BooleanValueContext:
			res.BooleanValue = NewBooleanValueNode(ctx, child.(*parser.BooleanValueContext))
		case *parser.StringValueContext:
			res.StringValue = NewStringValueNode(ctx, child.(*parser.StringValueContext))
		case *parser.IdentifierContext:
			res.Identifier = NewIdentifierNode(ctx, child.(*parser.IdentifierContext))
		default:
			//		self.Null = NewNullNode(ctx)
		}
	} else {
		res.ParenthesizedExpression = NewExpressionNode(ctx, children[1].(*parser.ExpressionContext))
	}
	return res
}

func (self *PrimaryExpressionNode) Result(ctx *Context) interface{} {
	if self.Number != nil {
		return self.Number.Result(ctx)
	} else if self.BooleanValue != nil {
		return self.BooleanValue.Result(ctx)
	} else if self.StringValue != nil {
		return self.StringValue.Result(ctx)
	} else if self.Identifier != nil {
		return self.Identifier.Result(ctx)
	} else if self.ParenthesizedExpression != nil {
		return self.ParenthesizedExpression.Result(ctx)
	}
	return nil
}
