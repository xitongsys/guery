package plan

import (
	"fmt"

	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/connector"
	"github.com/xitongsys/guery/filesystem/partition"
	"github.com/xitongsys/guery/metadata"
)

type PlanScanNode struct {
	Catalog       string
	Schema        string
	Table         string
	Name          string
	Metadata      *metadata.Metadata
	InputMetadata *metadata.Metadata
	PartitionInfo *partition.PartitionInfo
	Output        PlanNode
	Filters       []*BooleanExpressionNode
}

func NewPlanScanNode(runtime *config.ConfigRuntime, name string) *PlanScanNode {
	catalog, schema, table := metadata.SplitTableName(runtime, name)
	res := &PlanScanNode{
		Catalog: catalog,
		Schema:  schema,
		Table:   table,
		Name:    name,
	}
	return res
}

func (self *PlanScanNode) GetNodeType() PlanNodeType {
	return SCANNODE
}

func (self *PlanScanNode) String() string {
	res := "PlanScanNode {\n"
	res += "Name: " + self.Name + "\n"
	res += "Metadata:" + fmt.Sprintf("%v", self.Metadata) + "\n"
	res += "Filters:" + fmt.Sprintf("%v", self.Filters) + "\n"
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

func (self *PlanScanNode) GetMetadata() *metadata.Metadata {
	return self.Metadata
}

func (self *PlanScanNode) SetMetadata() error {
	if self.Metadata != nil {
		return nil
	}
	ctr, err := connector.NewConnector(self.Catalog, self.Schema, self.Table)
	if err != nil {
		return err
	}

	md, err := ctr.GetMetadata()
	if err != nil {
		return err
	}
	self.Metadata = md.Copy()
	self.Metadata.Reset()
	self.InputMetadata = md.Copy()
	self.InputMetadata.Reset()

	self.PartitionInfo, err = ctr.GetPartitionInfo()

	return err
}
