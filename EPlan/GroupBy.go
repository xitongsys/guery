package EPlan

import (
	"github.com/xitongsys/guery/Logger"
	. "github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/pb"
)

type EPlanGroupByNode struct {
	Location        *pb.Location
	Inputs, Outputs []*pb.Location
	GroupBy         *GroupByNode
}

func (self *EPlanGroupByNode) GetNodeType() EPlanNodeType {
	return EGROUPBYNODE
}

func CreateEPlanGroupByNode(node *PlanGroupByNode, pn int32, inputs, outputs []*pb.Location) *EPlanGroupByNode {
	if len(outputs) <= 0 || len(inputs) <= 0 {
		Logger.Errorf("inputs/outputs number <= 0")
		return nil
	}

	return &EPlanGroupByNode{
		Location: outputs[0],
		Inputs:   inputs,
		Outputs:  outputs,
		GroupBy:  node.GroupBy,
	}
}
