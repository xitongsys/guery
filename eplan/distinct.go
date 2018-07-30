package eplan

import (
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/pb"
	"github.com/xitongsys/guery/plan"
)

type EPlanDistinctNode struct {
	Location    pb.Location
	Inputs      []pb.Location
	Outputs     []pb.Location
	Expressions []*plan.ExpressionNode
	Metadata    *metadata.Metadata
}

func (self *EPlanDistinctNode) Init(md *metadata.Metadata) error {
	for _, e := range self.Expressions {
		if err := e.Init(md); err != nil {
			return err
		}
	}
	return nil
}

func (self *EPlanDistinctNode) GetNodeType() EPlanNodeType {
	return EDISTINCTNODE
}

func (self *EPlanDistinctNode) GetInputs() []pb.Location {
	return self.Inputs
}

func (self *EPlanDistinctNode) GetOutputs() []pb.Location {
	return self.Outputs
}

func (self *EPlanDistinctNode) GetLocation() pb.Location {
	return self.Location
}

func NewEPlanDistinctNode(node *plan.PlanDistinctNode,
	inputs []pb.Location, outputs []pb.Location) *EPlanDistinctNode {

	res := &EPlanDistinctNode{
		Location:    outputs[0],
		Inputs:      inputs,
		Outputs:     outputs,
		Expressions: node.Expressions,
		Metadata:    node.GetMetadata(),
	}
	return res
}
