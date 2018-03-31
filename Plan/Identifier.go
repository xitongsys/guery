package Plan

import (
	"github.com/xitongsys/guery/parser"
)

type IdentifierNode struct {
	Tree *parser.IdentifierContext
	Str  string
}

func NewIdentifierNode(ctx *Context, t *parser.IdentifierContext) *IdentifierNode {
	res := &IdentifierNode{
		Tree: t,
	}
	if t.IDENTIFIER() != nil {
		res.Str = t.IDENTIFIER().GetText()

	} else if t.QUOTED_IDENTIFIER() != nil {
		res.Str = t.QUOTED_IDENTIFIER().GetText()
		ln := len(res.Str)
		res.Str = res.Str[1 : ln-1]

	} else if t.NonReserved() != nil {
		res.Str = t.NonReserved().GetText()

	} else if t.BACKQUOTED_IDENTIFIER() != nil {
		res.Str = t.BACKQUOTED_IDENTIFIER().GetText()
		ln := len(res.Str)
		res.Str = res.Str[1 : ln-1]

	} else if t.DIGIT_IDENTIFIER() != nil {
		res.Str = t.DIGIT_IDENTIFIER().GetText()
		ln := len(res.Str)
		res.Str = res.Str[1 : ln-1]
	}
	return res
}

func (self *IdentifierNode) Result(ctx *Context) string {
	return self.Str
}
