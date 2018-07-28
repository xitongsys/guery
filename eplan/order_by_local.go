package eplan

import (
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/pb"
	. "github.com/xitongsys/guery/plan"
)

type EPlanOrderByLocalNode struct {
	Location  pb.Location
	Input     pb.Location
	Output    pb.Location
	SortItems []*SortItemNode
	Metadata  *Metadata.Metadata
}

func (self *EPlanOrderByLocalNode) GetNodeType() EPlanNodeType {
	return EORDERBYLOCALNODE
}

func (self *EPlanOrderByLocalNode) GetInputs() []pb.Location {
	return []pb.Location{self.Input}
}

func (self *EPlanOrderByLocalNode) GetOutputs() []pb.Location {
	return []pb.Location{self.Output}
}

func (self *EPlanOrderByLocalNode) GetLocation() pb.Location {
	return self.Location
}

func NewEPlanOrderByLocalNode(node *PlanOrderByNode, input pb.Location, output pb.Location) *EPlanOrderByLocalNode {
	return &EPlanOrderByLocalNode{
		Location:  output,
		Input:     input,
		Output:    output,
		SortItems: node.SortItems,
		Metadata:  node.GetMetadata(),
	}
}
