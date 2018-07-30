package eplan

import (
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/pb"
	"github.com/xitongsys/guery/plan"
)

type EPlanDistinctLocalNode struct {
	Location    pb.Location
	Inputs      []pb.Location
	Outputs     []pb.Location
	Expressions []*plan.ExpressionNode
	Metadata    *metadata.Metadata
}

func (self *EPlanDistinctLocalNode) Init(md *metadata.Metadata) error {
	for _, e := range self.Expressions {
		if err := e.Init(md); err != nil {
			return err
		}
	}
	return nil
}

func (self *EPlanDistinctLocalNode) GetNodeType() EPlanNodeType {
	return EDISTINCTLOCALNODE
}

func (self *EPlanDistinctLocalNode) GetInputs() []pb.Location {
	return self.Inputs
}

func (self *EPlanDistinctLocalNode) GetOutputs() []pb.Location {
	return self.Outputs
}

func (self *EPlanDistinctLocalNode) GetLocation() pb.Location {
	return self.Location
}

func NewEPlanDistinctLocalNode(node *plan.PlanDistinctLocalNode,
	inputs []pb.Location, outputs []pb.Location) *EPlanDistinctLocalNode {

	res := &EPlanDistinctLocalNode{
		Location:    outputs[0],
		Inputs:      inputs,
		Outputs:     outputs,
		Expressions: node.Expressions,
		Metadata:    node.GetMetadata(),
	}
	return res
}
