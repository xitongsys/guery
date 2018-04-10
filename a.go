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
	is := antlr.NewInputStream("SELECT * from t1 ")
	lexer := parser.NewSqlLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewSqlParser(stream)
	tree := p.SingleStatement()

	names, vals := []string{"name", "id"}, []interface{}{"a", int64(1)}
	tb := DataSource.NewTableSource("t1", names)
	tb.Append(vals)

	ctx := Context.NewContext()
	ctx.AddTable("t1", tb)

	q := Plan.NewPlanNodeFromSingleStatement(ctx, tree)
	fmt.Println(q.Execute())
}
