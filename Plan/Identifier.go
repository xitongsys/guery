package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/parser"
)

type IdentifierNode struct {
	tree       *parser.IdentifierContext
	identifier string
}

func NewIdentifierNode(t *parser.IdentifierContext) *IdentifierNode {
}

//UnquotedIdentifierNode
type UnquotedIdentifierNode struct {
	tree       *parser.UnquotedIdentifierContext
	identifier string
}

func NewUnquotedIdentifierNode(t *parser.UnquotedIdentifierContext) *UnquotedIdentifierNode {
	id := t.IDENTIFIER().GetText()
	return &UnquotedIdentifierNode{
		tree:       t,
		identifier: id,
	}
}

func (self *UnquotedIdentifierNode) Result() string {
	return self.identifier
}

//QuotedIdentifier
type QuotedIdentifierNode struct {
	tree       *parser.QuotedIdentifierContext
	identifier string
}

func NewQuotedIdentifierNode(t *parser.QuotedIdentifierContext) *QuotedIdentifierNode {
	id := t.IDENTIFIER().GetText()
	id = id[1 : len(id)-1]
	return &QuotedIdentifierNode{
		tree:       t,
		identifier: id,
	}
}

func (self *QuotedIdentifierNode) Result() string {
	return self.identifier
}

//BackquotedIdentifier
type BackQuotedIdentifierNode struct {
	tree       *parser.BackQuotedIdentifierContext
	identifier string
}

func NewBackQuotedIdentifierNode(t *parser.BackQuotedIdentifierContext) *BackQuotedIdentifierNode {
	id := t.IDENTIFIER().GetText()
	id = id[1 : len(id)-1]
	return &BackQuotedIdentifierNode{
		tree:       t,
		identifier: id,
	}
}

func (self *BackQuotedIdentifierNode) Result() string {
	return self.identifier
}

//DigitIdentifier
type DigitIdentifierNode struct {
	tree       *parser.DigitIdentifierContext
	identifier string
}

func NewDigitIdentifierNode(t *parser.DigitIdentifierContext) *DigitIdentifierNode {
	id := t.IDENTIFIER().GetText()
	return &DigitIdentifierNode{
		tree:       t,
		identifier: id,
	}
}

func (self *DigitIdentifierNode) Result() string {
	return self.identifier
}
