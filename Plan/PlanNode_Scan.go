package Plan

import (
	"github.com/xitongsys/guery/Catalog"
	"github.com/xitongsys/guery/Util"
)

type PlanScanNode struct {
	Name     string
	Metadata *Util.Metadata
	Output   PlanNode
}

func NewPlanScanNode(name string) *PlanScanNode {
	catalog, schema, table := Util.SplitName(name)
	res := &PlanScanNode{
		Name:     name,
		Metadata: Util.NewMetadata(catalog, schema, table, []string{}, []Util.Type{}),
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
	catalog, err := Catalog.NewCatalog(self.Metadata.Catalog, self.Metadata.Schema, self.Metadata.Table)
	if err != nil {
		return err
	}
	self.Metadata = catalog.GetMetadata()
	self.Metadata.Reset()

	return nil
}
