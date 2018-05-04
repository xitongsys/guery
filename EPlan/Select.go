package EPlan

import (
	"github.com/xitongsys/guery/Logger"
	. "github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/pb"
)

type EPlanSelectNode struct {
	Location      *pb.Location
	Input, Output *pb.Location
	SelectItems   []*SelectItemNode
}

func (self *EPlanSelectNode) GetNodeType() EPlanNodeType {
	return ESELECTNODE
}

func NewEPlanSelectNode(node *PlanSelectNode, input, output *pb.Location) *EPlanSelectNode {
	return &EPlanSelectNode{
		Location:    outputs[i],
		Input:       inputs[i],
		Output:      outputs[i],
		SelectItems: node.SelectItems,
	}
}
