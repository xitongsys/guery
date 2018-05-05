package EPlan

import (
	. "github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/pb"
)

type EPlanOrderByNode struct {
	Location pb.Location
	Inputs   []pb.Location
	Outputs  []pb.Location
}

func (self *EPlanOrderByNode) GetNodeType() EPlanNodeType {
	return EORDERBYNODE
}

func (self *EPlanOrderByNode) GetOutputs() []pb.Location {
	return self.Outputs
}

func NewEPlanOrderByNode(node *PlanOrderByNode, inputs []pb.Location, outputs []pb.Location) *EPlanOrderByNode {
	return &EPlanOrderByNode{
		Location: outputs[0],
		Inputs:   inputs,
		Outputs:  outputs,
	}
}
