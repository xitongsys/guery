package Plan

import (
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/parser"
)

type PlanOrderByNode struct {
	Input     PlanNode
	Output    PlanNode
	Metadata  *Util.Metadata
	SortItems []*SortItemNode
	OrderType Util.OrderType
}

func NewPlanOrderByNode(input PlanNode, items []parser.ISortItemContext) *PlanOrderByNode {
	res := &PlanOrderByNode{
		Input:     input,
		Metadata:  Util.NewMetadata(),
		SortItems: []*SortItemNode{},
	}
	for _, item := range items {
		itemNode := NewSortItemNode(item)
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

func (self *PlanOrderByNode) GetMetadata() *Util.Metadata {
	return self.Metadata
}

func (self *PlanOrderByNode) SetMetadata() error {
	if err := self.Input.SetMetadata(); err != nil {
		return err
	}
	self.Metadata = self.Input.GetMetadata().Copy()
	self.Metadata.ClearKeys()
	for _, item := range self.SortItems {
		t, err := item.GetType(self.Metadata)
		if err != nil {
			return err
		}
		self.Metadata.AppendKeyByType(t)
	}
	return nil
}
