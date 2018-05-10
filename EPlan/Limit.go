package EPlan

import (
	. "github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/pb"
)

type EPlanLimitNode struct {
	Location    pb.Location
	Inputs      []pb.Location
	Outputs     []pb.Location
	LimitNumber *int64
	Metadata    *Util.Metadata
}

func (self *EPlanLimitNode) GetNodeType() EPlanNodeType {
	return ELIMITNODE
}

func (self *EPlanLimitNode) GetOutputs() []pb.Location {
	return self.Outputs
}

func (self *EPlanLimitNode) GetLocation() pb.Location {
	return self.Location
}

func NewEPlanLimitNode(node *PlanLimitNode, inputs, outputs []pb.Location) *EPlanLimitNode {
	return &EPlanLimitNode{
		Location:    outputs[0],
		Inputs:      inputs,
		Outputs:     outputs,
		LimitNumber: node.LimitNumber,
		Metadata:    node.GetMetadata(),
	}
}
