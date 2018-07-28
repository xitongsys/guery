package plan

import (
	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/connector"
	"github.com/xitongsys/guery/gtype"
	"github.com/xitongsys/guery/metadata"
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
	Metadata *metadata.Metadata
	ShowType PlanShowNodeType

	//show catalogs/schemas/tables/columns/createtable/createview
	Catalog     string
	Schema      string
	Table       string
	LikePattern *string
	Escape      *string
}

func NewPlanShowNodeTables(runtime *config.ConfigRuntime, catalog, schema string, like, escape *string) *PlanShowNode {
	return &PlanShowNode{
		ShowType:    SHOWTABLES,
		Catalog:     catalog,
		Schema:      schema,
		LikePattern: like,
		Escape:      escape,
	}
}

func NewPlanShowNodeSchemas(runtime *config.ConfigRuntime, catalog string, like, escape *string) *PlanShowNode {
	return &PlanShowNode{
		ShowType:    SHOWSCHEMAS,
		Catalog:     catalog,
		LikePattern: like,
		Escape:      escape,
	}
}

func NewPlanShowNodeColumns(runtime *config.ConfigRuntime, catalog, schema, table string) *PlanShowNode {
	return &PlanShowNode{
		ShowType: SHOWCOLUMNS,
		Catalog:  catalog,
		Schema:   schema,
		Table:    table,
	}
}

func NewPlanShowNodePartitions(runtime *config.ConfigRuntime, catalog, schema, table string) *PlanShowNode {
	return &PlanShowNode{
		ShowType: SHOWPARTITIONS,
		Catalog:  catalog,
		Schema:   schema,
		Table:    table,
	}
}

func (self *PlanShowNode) GetNodeType() PlanNodeType {
	return SHOWNODE
}

func (self *PlanShowNode) SetMetadata() error {
	res := metadata.NewMetadata()
	switch self.ShowType {
	case SHOWCATALOGS:
	case SHOWTABLES:
		col := metadata.NewColumnMetadata(gtype.STRING, self.Catalog, self.Schema, "*", "table")
		res.AppendColumn(col)
	case SHOWSCHEMAS:
		col := metadata.NewColumnMetadata(gtype.STRING, self.Catalog, "*", "*", "schema")
		res.AppendColumn(col)
	case SHOWCOLUMNS:
		col := metadata.NewColumnMetadata(gtype.STRING, self.Catalog, self.Schema, self.Table, "NAME")
		res.AppendColumn(col)
		col = metadata.NewColumnMetadata(gtype.STRING, self.Catalog, self.Schema, self.Table, "TYPE")
		res.AppendColumn(col)
	case SHOWPARTITIONS:
		ctr, err := connector.NewConnector(self.Catalog, self.Schema, self.Table)
		if err != nil {
			return err
		}
		parInfo, err := ctr.GetPartitionInfo()
		if err != nil {
			return err
		}
		res = parInfo.Metadata

	}

	self.Metadata = res

	return nil
}

func (self *PlanShowNode) GetMetadata() *metadata.Metadata {
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
