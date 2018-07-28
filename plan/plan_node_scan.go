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
	Metadata      *Metadata.Metadata
	InputMetadata *Metadata.Metadata
	PartitionInfo *Partition.PartitionInfo
	Output        PlanNode
	Filters       []*BooleanExpressionNode
}

func NewPlanScanNode(runtime *Config.ConfigRuntime, name string) *PlanScanNode {
	catalog, schema, table := Metadata.SplitTableName(runtime, name)
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

func (self *PlanScanNode) GetMetadata() *Metadata.Metadata {
	return self.Metadata
}

func (self *PlanScanNode) SetMetadata() error {
	if self.Metadata != nil {
		return nil
	}
	connector, err := Connector.NewConnector(self.Catalog, self.Schema, self.Table)
	if err != nil {
		return err
	}

	md, err := connector.GetMetadata()
	if err != nil {
		return err
	}
	self.Metadata = md.Copy()
	self.Metadata.Reset()
	self.InputMetadata = md.Copy()
	self.InputMetadata.Reset()

	self.PartitionInfo, err = connector.GetPartitionInfo()

	return err
}
