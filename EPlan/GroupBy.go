package EPlan

import (
	"github.com/xitongsys/guery/Logger"
	. "github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/pb"
)

type EPlanGroupByNode struct {
	Location        pb.Location
	Inputs, Outputs []pb.Location
	GroupBy         *GroupByNode
	Metadata        *Util.Metadata
}

func (self *EPlanGroupByNode) GetNodeType() EPlanNodeType {
	return EGROUPBYNODE
}

func (self *EPlanGroupByNode) GetInputs() []pb.Location {
	return self.Inputs
}

func (self *EPlanGroupByNode) GetOutputs() []pb.Location {
	return self.Outputs
}

func (self *EPlanGroupByNode) GetLocation() pb.Location {
	return self.Location
}

func NewEPlanGroupByNode(node *PlanGroupByNode, inputs, outputs []pb.Location) *EPlanGroupByNode {
	if len(outputs) <= 0 || len(inputs) <= 0 {
		Logger.Errorf("inputs/outputs number <= 0")
		return nil
	}

	return &EPlanGroupByNode{
		Location: outputs[0],
		Inputs:   inputs,
		Outputs:  outputs,
		GroupBy:  node.GroupBy,
		Metadata: node.GetMetadata(),
	}
}
