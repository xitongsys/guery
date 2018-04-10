package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/DataSource"
	"github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/parser"
)

func main() {
	is := antlr.NewInputStream("SELECT NAME,ID FROM T1 ")
	lexer := parser.NewSqlLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewSqlParser(stream)
	tree := p.SingleStatement()

	names := []string{"NAME", "ID"}

	tb := DataSource.NewTableSource("T1", names)
	vals := []interface{}{"a", int64(1)}
	tb.Append(vals)
	vals = []interface{}{"b", int64(2)}
	tb.Append(vals)

	ctx := Context.NewContext()
	ctx.AddTable("T1", tb)

	q := Plan.NewPlanNodeFromSingleStatement(ctx, tree)
	res := q.Execute()
	for !res.IsEnd() {
		fmt.Println(res.ReadRow())
		res.Next()
	}
}
