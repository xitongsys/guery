package main

import (
	"fmt"
	"reflect"

	"./parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func Visit(node antlr.Tree) {
	switch node.(type) {
	case antlr.ErrorNode:
		return

	case antlr.TerminalNode:
		fmt.Println("===", reflect.TypeOf(node), node)
		return

	default:
		/*
			if _, ok := node.(*parser.ValueExpressionDefaultContext); ok {
				children := node.GetChildren()
				for _, child := range children {
					fmt.Println("----", child, reflect.TypeOf(child))
					if _, ok := child.(antlr.ParserRuleContext); ok {
						fmt.Println("=======", child.(antlr.ParserRuleContext).GetText())
					}

				}
			} else {
		*/
		s := node.(antlr.ParserRuleContext).GetText()
		fmt.Println(s, reflect.TypeOf(node))
		children := node.GetChildren()
		for i := 0; i < len(children); i++ {
			fmt.Printf("[%v] ", reflect.TypeOf(children[i]))
		}
		fmt.Println("")
		fmt.Println("")

		for i := 0; i < len(children); i++ {
			Visit(children[i])
		}
		//		}
	}
}

func main() {
	is := antlr.NewInputStream("SELECT ABC WHERE TRUE AND FALSE ORDER BY A")
	lexer := parser.NewSqlLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewSqlParser(stream)
	tree := p.SingleStatement()
	Visit(tree)

}
