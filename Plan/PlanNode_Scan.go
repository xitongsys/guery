package Plan

import (
	"github.com/xitongsys/guery/Util"
)

type PlanScanNode struct {
	Name     string
	Metadata *Util.Metadata
}

func NewPlanScanNode(name string) *PlanScanNode {
	catalog, schema, table := Util.SplitName(name)
	res := &PlanScanNode{
		Name:     name,
		Metadata: Util.NewMetadata(catalog, schema, table, []string{}, []string{}),
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

func (self *PlanScanNode) SetMetadata() *Metadata {
	md := Util.GetMetadata(self.Metadata.Catalog, self.Metadata.Schema, self.Metadata.Table)
	self.Metadata.ColumnNames = md.ColumnNames
	self.Metadata.ColumnTypes = md.ColumnTypes
	self.Metadata.Reset()
	return self.Metadata
}
