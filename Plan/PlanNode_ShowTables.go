package Plan

import (
	"github.com/xitongsys/guery/Config"
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Type"
)

type PlanShowTablesNode struct {
	Input       PlanNode
	Output      PlanNode
	Catalog     string
	Schema      string
	LikePattern *string
	Escape      *string

	Metadata *Metadata.Metadata
}

func NewPlanShowTablesNode(runtime *Config.ConfigRuntime, catalog, schema string, like, escape *string) *PlanShowTablesNode {
	return &PlanShowTablesNode{
		Catalog:     catalog,
		Schema:      schema,
		LikePattern: like,
		Escape:      escape,
	}
}

func (self *PlanShowTablesNode) GetNodeType() PlanNodeType {
	return SHOWTABLESNODE
}

func (self *PlanShowTablesNode) SetMetadata() error {
	res := Metadata.NewMetadata()
	col := Metadata.NewColumnMetadata(Type.STRING, self.Catalog, self.Schema, "", "table")
	res.AppendColumn(col)
	self.Metadata = res
	return nil
}

func (self *PlanShowTablesNode) GetMetadata() *Metadata.Metadata {
	return self.Metadata
}

func (self *PlanShowTablesNode) GetOutput() PlanNode {
	return self.Output
}

func (self *PlanShowTablesNode) SetOutput(output PlanNode) {
	self.Output = output
}

func (self *PlanShowTablesNode) GetInputs() []PlanNode {
	return []PlanNode{self.Input}
}

func (self *PlanShowTablesNode) SetInputs(inputs []PlanNode) {
}

func (self *PlanShowTablesNode) String() string {
	res := "PlanShowTablesNode {\n"
	res += "Catalog: " + self.Catalog + "\n"
	res += "Schema: " + self.Schema + "\n"
	res += "LikePattern: " + *self.LikePattern + "\n"
	res += "Escape: " + *self.Escape + "\n"
	res += "}\n"
	return res
}
