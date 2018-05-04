package EPlan

import (
	"github.com/xitongsys/guery/Logger"
	. "github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/pb"
)

type EPlanUnionNode struct {
	Location                *pb.Location
	LeftInputs, RightInputs []*pb.Location
	Outputs                 []*pb.Location
	Operator                UnionType
}

func (self *EPlanUnionNode) GetNodeType() EPlanNodeType {
	return UNIONNODE
}

func NewEPlanUnionNode(node *PlanUnionNode, leftInputs, rightInputs []*pb.Location, outputs []*pb.Location) *EPlanUnionNode {
	return &EPlanUnionNode{
		Location:    outputs[0],
		LeftInputs:  leftInputs,
		RightInputs: rightInputs,
		Outputs:     outputs,
		Operator:    node.Operator,
	}
}
