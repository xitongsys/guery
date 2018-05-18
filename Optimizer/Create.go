package Optimizer

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/parser"
)

func CreateLogicalTree(sqlStr string) (Plan.PlanNode, error) {
	is := antlr.NewInputStream(sqlStr)
	in := parser.NewCaseChangingStream(is, true)
	lexer := parser.NewSqlLexer(in)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewSqlParser(stream)
	tree := p.SingleStatement()
	logicalTree := Plan.NewPlanNodeFromSingleStatement(tree)

	if err := logicalTree.SetMetadata(); err != nil {
		return nil, err
	}

	//optimizer
	if err := DeleteRenameNode(logicalTree); err != nil {
		return logicalTree, err
	}

	return logicalTree, nil
}
