package Plan

import (
	"context"
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/parser"
)

func NewPlanCombineNode(plans []PlanNode) *PlanCombineNode {
	return &PlanCombineNode{
		Inputs:   plans,
		Metadata: Util.NewDefaultMetadata(),
	}
}

func (self *PlanCombineNode) GetNodeType() PlanNodeType {
	return COMBINENODE
}

func (self *PlanCombineNode) String() string {
	res := "PlanCombineNode {\n"
	for _, n := range self.Inputs {
		res += n.String()
	}
	res += "}\n"
	return res
}
