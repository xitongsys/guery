package Plan

import (
	"context"
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/parser"
)

type PlanRenameNode struct {
	Rename   string
	Metadata *Util.Metadata
	Input    PlanNode
	Output   PlanNode
}

func NewPlanRenameNode(input, output PlanNode, tname string) *PlanRenameNode {
	return &PlanRenameNode{
		Rename:   tname,
		Metadata: Util.NewDefaultMetadata(),
		Input:    input,
		Output:   output,
	}
}

func (self *PlanRenameNode) GetNodeType() PlanNodeType {
	return RENAMENODE
}

func (self *PlanRenameNode) String() string {
	res := "PlanRenameNode {\n"
	res += "Rename: " + self.Rename + "\n"
	res += "Input: " + self.Input.String() + "\n"
	res += "}\n"
	return res
}
