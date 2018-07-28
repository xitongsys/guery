package eplan

import (
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/pb"
	. "github.com/xitongsys/guery/plan"
)

type EPlanLimitNode struct {
	Location    pb.Location
	Input       pb.Location
	Output      pb.Location
	LimitNumber *int64
	Metadata    *Metadata.Metadata
}

func (self *EPlanLimitNode) GetNodeType() EPlanNodeType {
	return ELIMITNODE
}

func (self *EPlanLimitNode) GetInputs() []pb.Location {
	return []pb.Location{self.Input}
}

func (self *EPlanLimitNode) GetOutputs() []pb.Location {
	return []pb.Location{self.Output}
}

func (self *EPlanLimitNode) GetLocation() pb.Location {
	return self.Location
}

func NewEPlanLimitNode(node *PlanLimitNode, input, output pb.Location) *EPlanLimitNode {
	return &EPlanLimitNode{
		Location:    output,
		Input:       input,
		Output:      output,
		LimitNumber: node.LimitNumber,
		Metadata:    node.GetMetadata(),
	}
}
