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
	is := antlr.NewInputStream("SELECT ID,NAME,SUM(AGE) FROM T1 GROUP BY ID,NAME")
	lexer := parser.NewSqlLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewSqlParser(stream)
	tree := p.SingleStatement()

	names := []string{"NAME", "ID", "AGE"}
	columnBuf1 := DataSource.NewMemColumnBuffer()
	columnBuf1.Append("A", "A", "C")
	columnBuf2 := DataSource.NewMemColumnBuffer()
	columnBuf2.Append(int64(1), int64(1), int64(3))
	columnBuf3 := DataSource.NewMemColumnBuffer()
	columnBuf3.Append(int64(4), int64(5), int64(6))

	columnBuffers := []DataSource.ColumnBuffer{}
	columnBuffers = append(columnBuffers, columnBuf1, columnBuf2, columnBuf3)

	ds := DataSource.NewDataSource([]string{"T1"}, names, columnBuffers)

	ctx := Context.NewContext()
	ctx.AddTable("T1", ds)

	q := Plan.NewPlanNodeFromSingleStatement(ctx, tree)
	res := q.Execute()
	for !res.IsEnd() {
		fmt.Println(res.GetRawVals())
		res.Next()
	}
}
