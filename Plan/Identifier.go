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
	res := &IdentifierNode{
		tree:       t,
		identifier: "",
	}
}

func (self *IdentifierNode) Result() string {
	return self.identifier
}
