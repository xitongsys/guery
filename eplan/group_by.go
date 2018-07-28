package eplan

import (
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/pb"
	. "github.com/xitongsys/guery/plan"
)

type EPlanGroupByNode struct {
	Location      pb.Location
	Input, Output pb.Location
	GroupBy       *GroupByNode
	Metadata      *metadata.Metadata
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
		Location: output,
		Input:    input,
		Output:   output,
		GroupBy:  node.GroupBy,
		Metadata: node.GetMetadata(),
	}
}
