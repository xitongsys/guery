package EPlan

import (
	"github.com/xitongsys/guery/Metadata"
	. "github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/pb"
)

type EPlanGroupByLocalNode struct {
	Location      pb.Location
	Input, Output pb.Location
	GroupBy       *GroupByNode
	Metadata      *Metadata.Metadata
}

func (self *EPlanGroupByLocalNode) GetNodeType() EPlanNodeType {
	return EGROUPBYLOCALNODE
}

func (self *EPlanGroupByLocalNode) GetInputs() []pb.Location {
	return []pb.Location{self.Input}
}

func (self *EPlanGroupByLocalNode) GetOutputs() []pb.Location {
	return []pb.Location{self.Output}
}

func (self *EPlanGroupByLocalNode) GetLocation() pb.Location {
	return self.Location
}

func NewEPlanGroupByLocalNode(node *PlanGroupByNode, input, output pb.Location) *EPlanGroupByLocalNode {
	return &EPlanGroupByLocalNode{
		Location: output,
		Input:    input,
		Output:   output,
		GroupBy:  node.GroupBy,
		Metadata: node.GetMetadata(),
	}
}
