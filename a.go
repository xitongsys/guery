package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/parser"
)

func main() {
	sql := `SELECT NEWNAME FROM (SELECT 
            NAME,
            (CASE 
            WHEN NAME='A' THEN 'ANAME'
            WHEN NAME='B' THEN 'BNAME'
            ELSE 'OTHER' END) AS NEWNAME
            FROM T1)AS T2`
	fmt.Println(sql)

	is := antlr.NewInputStream(sql)
	lexer := parser.NewSqlLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewSqlParser(stream)
	tree := p.SingleStatement()
	q := Plan.NewPlanNodeFromSingleStatement(tree)
	fmt.Println(q)

}
