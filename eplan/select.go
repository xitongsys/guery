package eplan

import (
	"github.com/xitongsys/guery/gtype"
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/pb"
	. "github.com/xitongsys/guery/plan"
)

type EPlanSelectNode struct {
	Location      pb.Location
	Input, Output pb.Location
	SetQuantifier *gtype.QuantifierType
	SelectItems   []*SelectItemNode
	Having        *BooleanExpressionNode
	Metadata      *metadata.Metadata
	IsAggregate   bool
}

func (self *EPlanSelectNode) GetNodeType() EPlanNodeType {
	return ESELECTNODE
}

func (self *EPlanSelectNode) GetInputs() []pb.Location {
	return []pb.Location{self.Input}
}

func (self *EPlanSelectNode) GetOutputs() []pb.Location {
	return []pb.Location{self.Output}
}

func (self *EPlanSelectNode) GetLocation() pb.Location {
	return self.Location
}

func NewEPlanSelectNode(node *PlanSelectNode, input, output pb.Location) *EPlanSelectNode {
	return &EPlanSelectNode{
		Location:      output,
		Input:         input,
		Output:        output,
		SetQuantifier: node.SetQuantifier,
		SelectItems:   node.SelectItems,
		Having:        node.Having,
		Metadata:      node.GetMetadata(),
		IsAggregate:   node.IsAggregate,
	}
}
