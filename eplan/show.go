package eplan

import (
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/pb"
	. "github.com/xitongsys/guery/plan"
)

type EPlanShowNode struct {
	Location pb.Location
	Output   pb.Location
	Metadata *metadata.Metadata
	ShowType PlanShowNodeType

	//show catalogs/schemas/tables/columns/createtable/createview
	Catalog     string
	Schema      string
	Table       string
	LikePattern *string
	Escape      *string
}

func (self *EPlanShowNode) GetNodeType() EPlanNodeType {
	return ESHOWNODE
}

func (self *EPlanShowNode) GetInputs() []pb.Location {
	return []pb.Location{}
}

func (self *EPlanShowNode) GetOutputs() []pb.Location {
	return []pb.Location{self.Output}
}

func (self *EPlanShowNode) GetLocation() pb.Location {
	return self.Location
}

func NewEPlanShowNode(node *PlanShowNode, output pb.Location) *EPlanShowNode {
	return &EPlanShowNode{
		Location: output,
		Output:   output,
		Metadata: node.GetMetadata(),
		ShowType: node.ShowType,

		Catalog:     node.Catalog,
		Schema:      node.Schema,
		Table:       node.Table,
		LikePattern: node.LikePattern,
		Escape:      node.Escape,
	}
}
