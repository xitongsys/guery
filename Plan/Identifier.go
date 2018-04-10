package Plan

import (
	"fmt"

	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/DataSource"
	"github.com/xitongsys/guery/parser"
)

type IdentifierNode struct {
	Str         *string
	Digit       *int
	NonReserved *string
}

func NewIdentifierNode(ctx *Context.Context, t parser.IIdentifierContext) *IdentifierNode {
	res := &IdentifierNode{}
	var (
		str string
		dig int
	)

	if t.IDENTIFIER() != nil {
		str = t.IDENTIFIER().GetText()
		res.Str = &str

	} else if t.QUOTED_IDENTIFIER() != nil {
		str = t.QUOTED_IDENTIFIER().GetText()
		ln := len(res.Str)
		str = res.Str[1 : ln-1]
		res.Str = &str

	} else if t.NonReserved() != nil {
		str = t.NonReserved().GetText()
		res.NonReserved = &str

	} else if t.DIGIT_IDENTIFIER() != nil {
		str = t.DIGIT_IDENTIFIER().GetText()
		fmt.Sscanf(str, "%d", &dig)
		res.Digit = &dig
	}
	return res
}

func (self *IdentifierNode) Result(intput DataSource.DataSource) interface{} {
	if self.Str != nil {
		return intput.First().ReadColumnByName(self.Str)
	} else if self.Digit != nil {
		return input.First().ReadColumnByIndex(self.Digit)
	}
}
