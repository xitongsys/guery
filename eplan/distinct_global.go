package eplan

import (
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/pb"
	"github.com/xitongsys/guery/plan"
)

type EPlanDistinctGlobalNode struct {
	Location    pb.Location
	Inputs      []pb.Location
	Outputs     []pb.Location
	Expressions []*plan.ExpressionNode
	Metadata    *metadata.Metadata
}

func (self *EPlanDistinctGlobalNode) Init(md *metadata.Metadata) error {
	for _, e := range self.Expressions {
		if err := e.Init(md); err != nil {
			return err
		}
	}
	return nil
}

func (self *EPlanDistinctGlobalNode) GetNodeType() EPlanNodeType {
	return EDISTINCTGLOBALNODE
}

func (self *EPlanDistinctGlobalNode) GetInputs() []pb.Location {
	return self.Inputs
}

func (self *EPlanDistinctGlobalNode) GetOutputs() []pb.Location {
	return self.Outputs
}

func (self *EPlanDistinctGlobalNode) GetLocation() pb.Location {
	return self.Location
}

func NewEPlanDistinctGlobalNode(node *plan.PlanDistinctGlobalNode,
	inputs []pb.Location, outputs []pb.Location) *EPlanDistinctGlobalNode {

	res := &EPlanDistinctGlobalNode{
		Location:    outputs[0],
		Inputs:      inputs,
		Outputs:     outputs,
		Expressions: node.Expressions,
		Metadata:    node.GetMetadata(),
	}
	return res
}
