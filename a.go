package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/parser"
	"github.com/xitongsys/guery/pb"
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

	logicalPlanTree := Plan.NewPlanNodeFromSingleStatement(tree)
	freeExecutor := []pb.Location{}
	for i := 0; i < 100; i++ {
		freeExecutor = append(freeExecutor, pb.Location{Name: fmt.Sprintf("%v", i)})
	}
	res := []EPlan.ENode{}

	EPlan.CreateEPlan(logicalPlanTree, res, freeExecutor, 2)
	fmt.Println(logicalPlanTree, res)

}
