package eplan

import (
	"github.com/xitongsys/guery/filesystem/partition"
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/pb"
	. "github.com/xitongsys/guery/plan"
)

type EPlanScanNode struct {
	Location      pb.Location
	Catalog       string
	Schema        string
	Table         string
	Metadata      *Metadata.Metadata
	InputMetadata *Metadata.Metadata
	PartitionInfo *Partition.PartitionInfo

	Outputs []pb.Location
	Filters []*BooleanExpressionNode
}

func (self *EPlanScanNode) GetNodeType() EPlanNodeType {
	return ESCANNODE
}

func (self *EPlanScanNode) GetInputs() []pb.Location {
	return []pb.Location{}
}

func (self *EPlanScanNode) GetOutputs() []pb.Location {
	return self.Outputs
}

func (self *EPlanScanNode) GetLocation() pb.Location {
	return self.Location
}

func NewEPlanScanNode(node *PlanScanNode, parInfo *Partition.PartitionInfo, loc pb.Location, outputs []pb.Location) *EPlanScanNode {
	parInfo.Encode()
	return &EPlanScanNode{
		Location:      loc,
		Catalog:       node.Catalog,
		Schema:        node.Schema,
		Table:         node.Table,
		Outputs:       outputs,
		Metadata:      node.GetMetadata(),
		InputMetadata: node.InputMetadata,
		PartitionInfo: parInfo,
		Filters:       node.Filters,
	}
}
