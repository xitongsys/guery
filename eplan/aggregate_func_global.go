package eplan

import (
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/pb"
	"github.com/xitongsys/guery/plan"
)

type EPlanAggregateFuncGlobalNode struct {
	Location  pb.Location
	Inputs    []pb.Location
	Output    pb.Location
	FuncNodes []*plan.FuncCallNode
	Metadata  *metadata.Metadata
}

func (self *EPlanAggregateFuncGlobalNode) Init(md *metadata.Metadata) error {
	for _, f := range self.FuncNodes {
		if err := f.Init(md); err != nil {
			return err
		}
	}
	return nil
}

func (self *EPlanAggregateFuncGlobalNode) GetNodeType() EPlanNodeType {
	return EAGGREGATEFUNCGLOBALNODE
}

func (self *EPlanAggregateFuncGlobalNode) GetInputs() []pb.Location {
	return self.Inputs
}

func (self *EPlanAggregateFuncGlobalNode) GetOutputs() []pb.Location {
	return []pb.Location{self.Output}
}

func (self *EPlanAggregateFuncGlobalNode) GetLocation() pb.Location {
	return self.Location
}

func NewEPlanAggregateFuncGlobalNode(node *plan.PlanAggregateFuncGlobalNode,
	inputs []pb.Location, output pb.Location) *EPlanAggregateFuncGlobalNode {

	res := &EPlanAggregateFuncGlobalNode{
		Location:  output,
		Inputs:    inputs,
		Output:    output,
		FuncNodes: node.FuncNodes,
		Metadata:  node.GetMetadata(),
	}
	return res
}
