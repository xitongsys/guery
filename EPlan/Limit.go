package EPlan

import (
	"github.com/xitongsys/guery/Logger"
	. "github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/pb"
)

type EPlanLimitNode struct {
	Location    *pb.Location
	Inputs      []*pb.Location
	Outputs     []*pb.Location
	LimitNumber *int64
}

func (self *EPlanLimitNode) GetNodeType() EPlanNodeType {
	return ELIMITNODE
}

func NewEPlanLimitNode(node *PlanLimitNode, inputs []*pb.Location, outputs []*pb.Location) *EPlanLimitNode {
	return &EPlanUnionNode{
		Location:    outputs[0],
		Inputs:      inputs,
		Outputs:     outputs,
		LimitNumber: node.LimitNumber,
	}
}
