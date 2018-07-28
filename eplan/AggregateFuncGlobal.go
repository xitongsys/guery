package EPlan

import (
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/pb"
)

type EPlanAggregateFuncGlobalNode struct {
	Location  pb.Location
	Inputs    []pb.Location
	Output    pb.Location
	FuncNodes []*Plan.FuncCallNode
	Metadata  *Metadata.Metadata
}

func (self *EPlanAggregateFuncGlobalNode) Init(md *Metadata.Metadata) error {
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

func NewEPlanAggregateFuncGlobalNode(node *Plan.PlanAggregateFuncGlobalNode,
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
