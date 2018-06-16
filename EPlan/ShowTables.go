package EPlan

import (
	"github.com/xitongsys/guery/Metadata"
	. "github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/pb"
)

type EPlanShowTablesNode struct {
	Catalog       string
	Schema        string
	Location      pb.Location
	Input, Output pb.Location
	Metadata      *Metadata.Metadata
}

func (self *EPlanShowTablesNode) GetNodeType() EPlanNodeType {
	return ESHOWTABLESNODE
}

func (self *EPlanShowTablesNode) GetInputs() []pb.Location {
	return []pb.Location{}
}

func (self *EPlanShowTablesNode) GetOutputs() []pb.Location {
	return []pb.Location{self.Output}
}

func (self *EPlanShowTablesNode) GetLocation() pb.Location {
	return self.Location
}

func NewEPlanShowTablesNode(node *PlanShowTablesNode, output pb.Location) *EPlanShowTablesNode {
	return &EPlanShowTablesNode{
		Location: output,
		Output:   output,
		Catalog:  node.Catalog,
		Schema:   node.Schema,
		Metadata: node.GetMetadata(),
	}
}
