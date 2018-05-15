package Optimizer

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/parser"
)

func CreateLogicalTree(sqlStr string) (Plan.PlanNode, error) {
	is := antlr.NewInputStream(sqlStr)
	lexer := parser.NewSqlLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewSqlParser(stream)
	tree := p.SingleStatement()
	logicalTree := Plan.NewPlanNodeFromSingleStatement(tree)

	//optimizer
	if err := DeleteRenameNode(logicalTree); err != nil {
		return logicalTree, err
	}

	return logicalTree, nil
}
