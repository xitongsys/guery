package Plan

import (
	"github.com/xitongsys/guery/parser"
)

type PlanShowTablesNode struct {
	Catalog     string
	Schema      string
	LikePattern *string
	Escape      *string
}

func NewPlanShowTablesNode(catalog, schema string, like, escape *string) *PlanShowTablesNode {
	return &PlanShowTablesNode{
		Catalog:     catalog,
		Schema:      schema,
		LikePattern: like,
		Escape:      escape,
	}
}
