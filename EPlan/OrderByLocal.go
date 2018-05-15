package EPlan

import (
	. "github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/pb"
)

type EPlanOrderByLocalNode struct {
	Location  pb.Location
	Input     pb.Location
	Output    pb.Location
	SortItems []*SortItemNode
	Metadata  *Util.Metadata
}

func (self *EPlanOrderByLocalNode) GetNodeType() EPlanNodeType {
	return EORDERBYLOCALNODE
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
