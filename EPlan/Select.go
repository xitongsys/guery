package EPlan

import (
	. "github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/pb"
)

type EPlanSelectNode struct {
	Location      pb.Location
	Input, Output pb.Location
	SelectItems   []*SelectItemNode
}

func (self *EPlanSelectNode) GetNodeType() EPlanNodeType {
	return ESELECTNODE
}

func (self *EPlanSelectNode) GetOutputs() []pb.Location {
	return []pb.Location{self.Output}
}

func NewEPlanSelectNode(node *PlanSelectNode, input, output pb.Location) *EPlanSelectNode {
	return &EPlanSelectNode{
		Location:    output,
		Input:       input,
		Output:      output,
		SelectItems: node.SelectItems,
	}
}
