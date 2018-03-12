package main

import (
	"fmt"

	"./antlr/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	is, _ := antlr.NewInputStream("select * from rbp")
	lexer := parser.NewSqlLexer(is)
	fmt.Println(lexer)

}
