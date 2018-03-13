package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Filter struct {
	ExpressionTree antlr.ParserRuleContext
}

func (self *Filter) GetColumns() []string {
	return nil
}

func (self *Filter) Eval(vals []interface{}) bool {
	return true
}
