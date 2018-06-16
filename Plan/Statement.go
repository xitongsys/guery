package Plan

import (
	"github.com/xitongsys/guery/Config"
	"github.com/xitongsys/guery/parser"
)

func NewPlanNodeFromSingleStatement(runtime *Config.ConfigRuntime, t parser.ISingleStatementContext) PlanNode {
	tt := t.(*parser.SingleStatementContext)
	return NewPlanNodeFromStatement(runtime, tt.Statement())
}

func NewPlanNodeFromStatement(runtime *Config.ConfigRuntime, t parser.IStatementContext) PlanNode {
	tt := t.(*parser.StatementContext)
	if tt.Query() != nil {
		return NewPlanNodeFromQuery(runtime, tt.Query())
	}

	//use
	if tt.USE() != nil {
		catalog, schema := runtime.Catalog, runtime.Schema

		if ct := tt.GetCatalog(); ct != nil {
			catalogNode := NewIdentifierNode(runtime, ct)
			catalog = catalogNode.GetText()
		}

		if sc := tt.GetSchema(); sc != nil {
			schemaNode := NewIdentifierNode(runtime, sc)
			schema = schemaNode.GetText()
		}

		return NewPlanUseNode(runtime, catalog, schema)
	}

	//show tables
	if tt.SHOW() != nil && tt.TABLES() != nil {
		if qname := tt.QualifiedName(); qname != nil {

		} else {
		}
	}

	return nil
}
