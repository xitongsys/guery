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
	Metadata  *Metadata.Metadata
	SortItems []*SortItemNode
	OrderType Type.OrderType
}

func NewPlanOrderByNode(runtime *Config.ConfigRuntime, input PlanNode, items []parser.ISortItemContext) *PlanOrderByNode {
	res := &PlanOrderByNode{
		Input:     input,
		Metadata:  Metadata.NewMetadata(),
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

func (self *PlanOrderByNode) GetMetadata() *Metadata.Metadata {
	return self.Metadata
}

func (self *PlanOrderByNode) SetMetadata() error {
	if err := self.Input.SetMetadata(); err != nil {
		return err
	}
	self.Metadata = self.Input.GetMetadata().Copy()
	return nil
}
