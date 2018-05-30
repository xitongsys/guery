package parser

import (
	"fmt"
	"strings"

	. "github.com/antlr/antlr4/runtime/Go/antlr"
)

type GueryErrorListener struct {
	*DefaultErrorListener
	Msgs []string
}

func NewGueryErrorListener() *GueryErrorListener {
	return new(GueryErrorListener)
}

func (self *GueryErrorListener) SyntaxError(recognizer Recognizer, offendingSymbol interface{}, line, column int, msg string, e RecognitionException) {
	self.Msgs = append(self.Msgs, fmt.Sprintf("line %v:%v  ", line, column)+msg)
}

func (self *GueryErrorListener) HasError() bool {
	return len(self.Msgs) > 0
}

func (self *GueryErrorListener) GetErrorMsgs() string {
	return strings.Join(self.Msgs, "\n")
}
