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
		Metadata:  Util.NewDefaultMetadata(),
		SortItems: []*SortItemNode{},
	}
	for _, item := range items {
		itemNode := NewSortItemNode(item)
		res.SortItems = append(res.SortItems, itemNode)
	}
	return res
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
	err := self.Input.SetMetadata()
	if err != nil {
		return err
	}
	self.Metadata.Copy(self.Input.GetMetadata())
	return nil
}
