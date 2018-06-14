package Plan

import (
	"github.com/xitongsys/guery/Config"
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromSingleStatement(t parser.ISingleStatementContext) PlanNode {
	tt := t.(*parser.SingleStatementContext)
	return NewPlanNodeFromStatement(tt.Statement())
}

func NewPlanNodeFromStatement(t parser.IStatementContext) PlanNode {
	tt := t.(*parser.StatementContext)
	if tt.Query() != nil {
		return NewPlanNodeFromQuery(tt.Query())
	}

	if tt.USE() != nil {
		catalog, schema := Config.Conf.Runtime.Catalog, Config.Conf.Runtime.Schema

		if ct := tt.GetCatalog(); ct != nil {
			catalogNode := NewIdentifierNode(ct)
			catalog = catalogNode.GetText()
		}

		if sc := tt.GetSchema(); sc != nil {
			schemaNode := NewIdentifierNode(sc)
			schema = schemaNode.GetText()
		}

		return NewPlanUseNode(catalog, schema)
	}

	return nil
}
