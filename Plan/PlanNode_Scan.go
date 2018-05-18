package Plan

import (
	"github.com/xitongsys/guery/Catalog"
	"github.com/xitongsys/guery/Util"
)

type PlanScanNode struct {
	Catalog  string
	Schema   string
	Table    string
	Name     string
	Metadata *Util.Metadata
	Output   PlanNode
}

func NewPlanScanNode(name string) *PlanScanNode {
	catalog, schema, table := Util.SplitTableName(name)
	res := &PlanScanNode{
		Catalog:  catalog,
		Schema:   schema,
		Table:    table,
		Name:     name,
		Metadata: Util.NewMetadata(),
	}
	return res
}

func (self *PlanScanNode) GetNodeType() PlanNodeType {
	return SCANNODE
}

func (self *PlanScanNode) String() string {
	res := "PlanScanNode {\n"
	res += "Name: " + self.Name + "\n"
	res += "}\n"
	return res
}

func (self *PlanScanNode) GetInputs() []PlanNode {
	return []PlanNode{}
}

func (self *PlanScanNode) SetInputs(inputs []PlanNode) {
}

func (self *PlanScanNode) GetOutput() PlanNode {
	return self.Output
}

func (self *PlanScanNode) SetOutput(output PlanNode) {
	self.Output = output
}

func (self *PlanScanNode) GetMetadata() *Util.Metadata {
	return self.Metadata
}

func (self *PlanScanNode) SetMetadata() error {
	catalog, err := Catalog.NewCatalog(self.Catalog, self.Schema, self.Table)

	if err != nil {
		return err
	}
	self.Metadata = catalog.GetMetadata()
	self.Metadata.Reset()

	return nil
}
