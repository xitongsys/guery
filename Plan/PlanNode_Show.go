package Plan

import (
	"github.com/xitongsys/guery/Config"
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Type"
)

type PlanShowNodeType int32

const (
	_ PlanShowNodeType = iota
	SHOWCATALOGS
	SHOWSCHEMAS
	SHOWTABLES
	SHOWCOLUMNS
	SHOWSTATS
	SHOWPARTITIONS
	SHOWCREATETABLE
	SHOWCREATEVIEW
)

func (self PlanShowNodeType) String() string {
	switch self {
	case SHOWCATALOGS:
		return "SHOWCATALOGS"
	case SHOWSCHEMAS:
		return "SHOWSCHEMAS"
	case SHOWTABLES:
		return "SHOWTABLES"
	case SHOWCOLUMNS:
		return "SHOWCOLUMNS"
	case SHOWSTATS:
		return "SHOWSTATS"
	case SHOWPARTITIONS:
		return "SHOWPARTITIONS"
	case SHOWCREATETABLE:
		return "SHOWCREATETABLE"
	case SHOWCREATEVIEW:
		return "SHOWCREATEVIEW"
	}
	return "UNKNOWNSHOWTYPE"
}

type PlanShowNode struct {
	Input    PlanNode
	Output   PlanNode
	Metadata *Metadata.Metadata
	ShowType PlanShowNodeType

	//show catalogs/schemas/tables/columns/createtable/createview
	Catalog     string
	Schema      string
	Table       string
	LikePattern *string
	Escape      *string
}

func NewPlanShowNodeTables(runtime *Config.ConfigRuntime, catalog, schema string, like, escape *string) *PlanShowNode {
	return &PlanShowNode{
		ShowType:    SHOWTABLES,
		Catalog:     catalog,
		Schema:      schema,
		LikePattern: like,
		Escape:      escape,
	}
}

func NewPlanShowNodeSchemas(runtime *Config.ConfigRuntime, catalog string, like, escape *string) *PlanShowNode {
	return &PlanShowNode{
		ShowType:    SHOWSCHEMAS,
		Catalog:     catalog,
		LikePattern: like,
		Escape:      escape,
	}
}

func NewPlanShowNodeColumns(runtime *Config.ConfigRuntime, catalog, schema, table string) *PlanShowNode {
	return &PlanShowNode{
		ShowType: SHOWCOLUMNS,
		Catalog:  catalog,
		Schema:   schema,
		Table:    table,
	}
}

func (self *PlanShowNode) GetNodeType() PlanNodeType {
	return SHOWNODE
}

func (self *PlanShowNode) SetMetadata() error {
	res := Metadata.NewMetadata()
	switch self.ShowType {
	case SHOWCATALOGS:
	case SHOWTABLES:
		col := Metadata.NewColumnMetadata(Type.STRING, self.Catalog, self.Schema, "*", "table")
		res.AppendColumn(col)
	case SHOWSCHEMAS:
		col := Metadata.NewColumnMetadata(Type.STRING, self.Catalog, "*", "*", "schema")
		res.AppendColumn(col)
	case SHOWCOLUMNS:
		col := Metadata.NewColumnMetadata(Type.STRING, self.Catalog, self.Schema, self.Table, "NAME")
		res.AppendColumn(col)
		col = Metadata.NewColumnMetadata(Type.STRING, self.Catalog, self.Schema, self.Table, "TYPE")
		res.AppendColumn(col)
	}

	self.Metadata = res

	return nil
}

func (self *PlanShowNode) GetMetadata() *Metadata.Metadata {
	return self.Metadata
}

func (self *PlanShowNode) GetOutput() PlanNode {
	return self.Output
}

func (self *PlanShowNode) SetOutput(output PlanNode) {
	self.Output = output
}

func (self *PlanShowNode) GetInputs() []PlanNode {
	return []PlanNode{self.Input}
}

func (self *PlanShowNode) SetInputs(inputs []PlanNode) {
}

func (self *PlanShowNode) String() string {
	res := "PlanShowNode {\n"
	res += "ShowType: " + self.ShowType.String() + "\n"
	res += "Catalog: " + self.Catalog + "\n"
	res += "Schema: " + self.Schema + "\n"
	res += "LikePattern: " + *self.LikePattern + "\n"
	res += "Escape: " + *self.Escape + "\n"
	res += "}\n"
	return res
}
