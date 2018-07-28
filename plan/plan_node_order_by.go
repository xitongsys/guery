package plan

import (
	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/gtype"
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/parser"
)

type PlanOrderByNode struct {
	Input     PlanNode
	Output    PlanNode
	Metadata  *metadata.Metadata
	SortItems []*SortItemNode
	OrderType gtype.OrderType
}

func NewPlanOrderByNode(runtime *config.ConfigRuntime, input PlanNode, items []parser.ISortItemContext) *PlanOrderByNode {
	res := &PlanOrderByNode{
		Input:     input,
		Metadata:  metadata.NewMetadata(),
		SortItems: []*SortItemNode{},
	}
	for _, item := range items {
		itemNode := NewSortItemNode(runtime, item)
		res.SortItems = append(res.SortItems, itemNode)
	}
	return res
}

func (self *PlanOrderByNode) GetInputs() []PlanNode {
	return []PlanNode{self.Input}
}

func (self *PlanOrderByNode) SetInputs(inputs []PlanNode) {
	self.Input = inputs[0]
}

func (self *PlanOrderByNode) GetOutput() PlanNode {
	return self.Output
}

func (self *PlanOrderByNode) SetOutput(output PlanNode) {
	self.Output = output
}

func (self *PlanOrderByNode) GetNodeType() PlanNodeType {
	return ORDERBYNODE
}

func (self *PlanOrderByNode) String() string {
	res := "PlanOrderByNode {\n"
	res += "Input: " + self.Input.String() + "\n"
	res += "}\n"
	return res
}

func (self *PlanOrderByNode) GetMetadata() *metadata.Metadata {
	return self.Metadata
}

func (self *PlanOrderByNode) SetMetadata() error {
	if err := self.Input.SetMetadata(); err != nil {
		return err
	}
	self.Metadata = self.Input.GetMetadata().Copy()
	return nil
}
