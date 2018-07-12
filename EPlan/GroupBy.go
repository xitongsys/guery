package EPlan

import (
	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/Metadata"
	. "github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/pb"
)

type EPlanGroupByNode struct {
	Location      pb.Location
	Input, Output []pb.Location
	GroupBy       *GroupByNode
	Metadata      *Metadata.Metadata
}

func (self *EPlanGroupByNode) GetNodeType() EPlanNodeType {
	return EGROUPBYNODE
}

func (self *EPlanGroupByNode) GetInputs() []pb.Location {
	return []pb.Location{self.Input}
}

func (self *EPlanGroupByNode) GetOutputs() []pb.Location {
	return []pb.Location{self.Output}
}

func (self *EPlanGroupByNode) GetLocation() pb.Location {
	return self.Location
}

func NewEPlanGroupByNode(node *PlanGroupByNode, input, output pb.Location) *EPlanGroupByNode {
	return &EPlanGroupByNode{
		Location: outputs[0],
		Input:    input,
		Output:   output,
		GroupBy:  node.GroupBy,
		Metadata: node.GetMetadata(),
	}
}
