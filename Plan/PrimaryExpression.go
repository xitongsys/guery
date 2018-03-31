package Plan

import (
	"github.com/xitongsys/guery/parser"
)

type PrimaryExpressionNode struct {
	Tree *parser.PrimaryExpressionContext
	//	Null         *NullNode
	Number       *NumberNode
	BooleanValue *BooleanValueNode
	StringValue  *StringValueNode
	Identifier   *IdentifierNode
}

func NewPrimaryExpressionNode(ctx *Context, t *parser.PrimaryExpressionContext) *PrimaryExpressionNode {
	res := &PrimaryExpressionNode{
		Tree: t,
	}
	child := t.GetChildren()[0]
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
	}
	return nil
}
