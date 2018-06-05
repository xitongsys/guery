package Plan

import (
	"fmt"

	"github.com/xitongsys/guery/Connector"
	"github.com/xitongsys/guery/Util"
)

type PlanScanNode struct {
	Catalog       string
	Schema        string
	Table         string
	Name          string
	Metadata      *Util.Metadata
	PartitionInfo *Util.PartitionInfo
	Output        PlanNode
	Filiters      []*BooleanExpressionNode
}

func NewPlanScanNode(name string) *PlanScanNode {
	catalog, schema, table := Util.SplitTableName(name)
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
	res += "Filiters:" + fmt.Sprintf("%v", self.Filiters) + "\n"
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
	if self.Metadata != nil {
		return nil
	}
	connector, err := Connector.NewConnector(self.Catalog, self.Schema, self.Table)
	if err != nil {
		return err
	}
	self.Metadata = connector.GetMetadata()
	self.Metadata.Reset()

	self.PartitionInfo = connector.GetPartitionInfo()

	return nil
}
