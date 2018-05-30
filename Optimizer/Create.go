package Optimizer

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/parser"
)

func CreateLogicalTree(sqlStr string) (node Plan.PlanNode, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	is := antlr.NewInputStream(sqlStr)
	in := parser.NewCaseChangingStream(is, true)
	lexer := parser.NewSqlLexer(in)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewSqlParser(stream)
	tree := p.SingleStatement()
	logicalTree := Plan.NewPlanNodeFromSingleStatement(tree)

	if err = logicalTree.SetMetadata(); err != nil {
		return nil, err
	}

	//optimizer
	if err = DeleteRenameNode(logicalTree); err != nil {
		return logicalTree, err
	}

	if err = FilterColumns(logicalTree, []string{}); err != nil {
		return logicalTree, err
	}

	if err = PredicatePushDown(logicalTree, []*Plan.BooleanExpressionNode{}); err != nil {
		return logicalTree, err
	}

	return logicalTree, nil
}
