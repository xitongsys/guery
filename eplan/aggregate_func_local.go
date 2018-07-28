package eplan

import (
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/pb"
	"github.com/xitongsys/guery/plan"
)

type EPlanAggregateFuncLocalNode struct {
	Location  pb.Location
	Input     pb.Location
	Output    pb.Location
	FuncNodes []*plan.FuncCallNode
	Metadata  *metadata.Metadata
}

func (self *EPlanAggregateFuncLocalNode) Init(md *metadata.Metadata) error {
	for _, f := range self.FuncNodes {
		if err := f.Init(md); err != nil {
			return err
		}
	}
	return nil
}

func (self *EPlanAggregateFuncLocalNode) GetNodeType() EPlanNodeType {
	return EAGGREGATEFUNCLOCALNODE
}

func (self *EPlanAggregateFuncLocalNode) GetInputs() []pb.Location {
	return []pb.Location{self.Input}
}

func (self *EPlanAggregateFuncLocalNode) GetOutputs() []pb.Location {
	return []pb.Location{self.Output}
}

func (self *EPlanAggregateFuncLocalNode) GetLocation() pb.Location {
	return self.Location
}

func NewEPlanAggregateFuncLocalNode(node *plan.PlanAggregateFuncLocalNode,
	input pb.Location, output pb.Location) *EPlanAggregateFuncLocalNode {

	res := &EPlanAggregateFuncLocalNode{
		Location:  output,
		Input:     input,
		Output:    output,
		FuncNodes: node.FuncNodes,
		Metadata:  node.GetMetadata(),
	}
	return res
}
