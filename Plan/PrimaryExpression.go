package Plan

import (
	"github.com/xitongsys/guery/parser"
)

type PrimaryExpressionNode struct {
	Tree            *parser.PrimaryExpressionContext
	NullLiteral     *NullLiteralNode
	NumericLiteral  *NumericLiteralNode
	BooleanLiteral  *BooleanLiteralNode
	StringLiteral   *StringLiteralNode
	ColumnReference *ColumnReferenceNode
}

func NewPrimaryExpressionNode(ctx *Context, t *parser.PrimaryExpressionContext) *PrimaryExpressionNode {
	res := &PrimaryExpressionNode{
		Tree: t,
	}
	return res
}

func (self *PrimaryExpressionNode) Result(ctx *Context) interface{} {
	if self.NullLiteral != nil {
		return self.NullLiteral.Result(ctx)
	} else if self.NumericLiteral != nil {
		return self.NumericLiteral.Result(ctx)
	} else if self.BooleanLiteral != nil {
		return self.BooleanLiteral.Result(ctx)
	} else if self.StringLiteral != nil {
		return self.StringLiteral.Result(ctx)
	} else if self.ColumnReference != nil {
		return self.ColumnReference.Result(ctx)
	}
	return nil
}

//NullLiteralNode
type NullLiteralNode struct{}

func NewNullLiteralNode(ctx *Context, t *parser.NullLiteralContext) *NullLiteralNode {
	return &NullLiteralNode{}
}

func (self *NullLiteralNode) Result(ctx *Context) interface{} {
	return nil
}

//NumericLiteralNode
type NumericLiteralNode struct {
	Tree   *parser.NumericLiteralContext
	Number *NumberNode
	Res    interface{}
}

func NewNumericLiteralNode(ctx *Context, t *parser.NumericLiteralContext) *NumericLiteralNode {
	res := &NumericLiteralNode{
		Tree:   t,
		Number: NewNumberNode(ctx, t.Number().(*parser.NumberContext)),
	}
	res.Res = res.Number.Result(ctx)
	return res
}

func (self *NumericLiteralNode) Result(ctx *Context) interface{} {
	return self.Res
}

//BooleanLiteralNode
type BooleanLiteralNode struct {
	Res bool
}

func NewBooleanLiteralNode(ctx *Context, t *parser.BooleanDefaultContext) *BooleanLiteralNode {
	res := &BooleanLiteralNode{}
	if t.GetText() == "TRUE" {
		res.Res = true
	} else {
		res.Res = false
	}
	return res
}

func (self *BooleanLiteralNode) Result(ctx *Context) bool {
	return self.Res
}

//StringLiteralNode
type StringLiteralNode struct {
	Res string
}

func NewStringLiteralNode(ctx *Context, t *parser.StringLiteralContext) *StringLiteralNode {
	res := &StringLiteralNode{}
	res.Res = t.GetText()
	return res
}

func (self *StringLiteralNode) Result(ctx *Context) string {
	return self.Res
}

//ColumnReference
type ColumnReferenceNode struct {
	Tree       *parser.ColumnReferenceContext
	Identifier *IdentifierNode
}

func NewColumnReferenceNode(ctx *Context, t *parser.ColumnReferenceContext) *ColumnReferenceNode {
	res := &ColumnReferenceNode{
		Tree:       t,
		Identifier: NewIdentifierNode(ctx, t.Identifier().(*parser.IdentifierContext)),
	}
	return res
}

func (self *ColumnReferenceNode) Result(ctx *Context) interface{} {
	name := self.Identifier.Result(ctx)
	return ctx.GetDataValue(name)
}
