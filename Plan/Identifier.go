package Plan

import (
	"github.com/xitongsys/guery/parser"
)

type IdentifierNode struct {
	Tree       *parser.IdentifierContext
	Identifier string

	UnquotedIdentifier   *UnquotedIdentifierNode
	QuotedIdentifier     *QuotedIdentifierNode
	BackQuotedIdentifier *BackQuotedIdentifierNode
	DigitIdentifier      *DigitIdentifierNode
}

func NewIdentifierNode(ctx *Context, t *parser.IdentifierContext) *IdentifierNode {
	children := t.GetChildren()
	if len(children) <= 0 {
		return nil
	}
	res := &IdentifierNode{
		Tree: t,
	}

	switch children[0].(type) {
	case *parser.UnquotedIdentifierContext:
		res.UnquotedIdentifier = NewUnquotedIdentifierNode(ctx, children[0].(*parser.UnquotedIdentifierContext))
		res.Identifier = res.UnquotedIdentifier.Result(ctx)

	case *parser.QuotedIdentifierContext:
		res.QuotedIdentifier = NewQuotedIdentifierNode(ctx, children[0].(*parser.QuotedIdentifierContext))
		res.Identifier = res.QuotedIdentifier.Result(ctx)

	case *parser.BackQuotedIdentifierContext:
		res.BackQuotedIdentifier = NewBackQuotedIdentifierNode(ctx, children[0].(*parser.BackQuotedIdentifierContext))
		res.Identifier = res.BackQuotedIdentifier.Result(ctx)

	case *parser.DigitIdentifierContext:
		res.DigitIdentifier = NewDigitIdentifierNode(ctx, children[0].(*parser.DigitIdentifierContext))
		res.Identifier = res.DigitIdentifier.Result(ctx)
	}

	return res
}

func (self *IdentifierNode) Result(ctx *Context) string {
	return self.Identifier
}

//UnquotedIdentifierNode
type UnquotedIdentifierNode struct {
	tree       *parser.UnquotedIdentifierContext
	identifier string
}

func NewUnquotedIdentifierNode(ctx *Context, t *parser.UnquotedIdentifierContext) *UnquotedIdentifierNode {
	id := t.IDENTIFIER().GetText()
	return &UnquotedIdentifierNode{
		tree:       t,
		identifier: id,
	}
}

func (self *UnquotedIdentifierNode) Result(ctx *Context) string {
	return self.identifier
}

//QuotedIdentifier
type QuotedIdentifierNode struct {
	tree       *parser.QuotedIdentifierContext
	identifier string
}

func NewQuotedIdentifierNode(ctx *Context, t *parser.QuotedIdentifierContext) *QuotedIdentifierNode {
	id := t.GetText()
	id = id[1 : len(id)-1]
	return &QuotedIdentifierNode{
		tree:       t,
		identifier: id,
	}
}

func (self *QuotedIdentifierNode) Result(ctx *Context) string {
	return self.identifier
}

//BackquotedIdentifier
type BackQuotedIdentifierNode struct {
	tree       *parser.BackQuotedIdentifierContext
	identifier string
}

func NewBackQuotedIdentifierNode(ctx *Context, t *parser.BackQuotedIdentifierContext) *BackQuotedIdentifierNode {
	id := t.GetText()
	id = id[1 : len(id)-1]
	return &BackQuotedIdentifierNode{
		tree:       t,
		identifier: id,
	}
}

func (self *BackQuotedIdentifierNode) Result(ctx *Context) string {
	return self.identifier
}

//DigitIdentifier
type DigitIdentifierNode struct {
	tree       *parser.DigitIdentifierContext
	identifier string
}

func NewDigitIdentifierNode(ctx *Context, t *parser.DigitIdentifierContext) *DigitIdentifierNode {
	id := t.GetText()
	return &DigitIdentifierNode{
		tree:       t,
		identifier: id,
	}
}

func (self *DigitIdentifierNode) Result(ctx *Context) string {
	return self.identifier
}
