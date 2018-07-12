package EPlan

import (
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/pb"
)

type EPlanAggregateFuncLocalNode struct {
	Location  pb.Location
	Input     pb.Location
	Output    pb.Location
	FuncNodes []*FuncCallNode
	Metadata  *Metadata.Metadata
}

func (self *EPlanAggregateFuncLocalNode) Init(md *Metadata.Metadata) error {
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

func NewEPlanAggregateFuncLocalNode(node *PlanAggregateFuncLocalNode,
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
