package Plan

import (
	"fmt"
	"strings"

	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/parser"
)

type PlanSelectNode struct {
	Input       PlanNode
	Metadata    *Util.Metadata
	SelectItems []*SelectItemNode
	IsAggregate bool
}

func NewPlanSelectNode(input PlanNode, items []parser.ISelectItemContext) *PlanSelectNode {
	res := &PlanSelectNode{
		Input:       input,
		Metadata:    Util.NewDefaultMetadata(),
		SelectItems: []*SelectItemNode{},
	}
	for i := 0; i < len(items); i++ {
		itemNode := NewSelectItemNode(items[i])
		res.SelectItems = append(res.SelectItems, itemNode)
		if itemNode.IsAggregate() {
			res.IsAggregate = true
		}
	}
	return res
}

func (self *PlanSelectNode) GetNodeType() PlanNodeType {
	return SELECTNODE
}

func (self *PlanSelectNode) String() string {
	res := "PlanSelectNode {\n"
	res += "Input: " + self.Input.String() + "\n"
	res += "SelectItems: " + fmt.Sprint(self.SelectItems) + "\n"
	res += "}\n"
	return res
}
