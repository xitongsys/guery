package main

import (
	"fmt"

	"./parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func Visit(node antlr.Tree) {
	switch node.(type) {
	case antlr.ErrorNode:
		return
	case antlr.TerminalNode:
		return
	default:
		if _, ok := node.(parser.IQuerySpecificationContext); ok {
			children := node.GetChildren()
			qnode := node.(*parser.QuerySpecificationContext)
			if qnode.FROM() != nil {

			}

			for i := 0; i < len(children); i++ {
				c := children[i]
				if _, ok := c.(antlr.TerminalNode); ok {
					continue
				}
				Visit(children[i])

			}

		} else {
			s := node.(antlr.ParserRuleContext).GetText()
			fmt.Println(s)
			children := node.GetChildren()
			for i := 0; i < len(children); i++ {
				Visit(children[i])
			}
		}
	}
}

func main() {
	is := antlr.NewInputStream("SELECT NAME FROM STUDENT WHERE ID=1")
	lexer := parser.NewSqlLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewSqlParser(stream)
	tree := p.SingleStatement()
	Visit(tree)

}
