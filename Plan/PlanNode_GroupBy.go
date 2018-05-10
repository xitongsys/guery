package Plan

import (
	"context"
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/parser"
)

type PlanGroupByNode struct {
	Input    PlanNode
	Output   PlanNode
	Metadata *Util.Metadata
	GroupBy  *GroupByNode
}

func NewPlanGroupByNode(input PlanNode, groupBy parser.IGroupByContext) *PlanGroupByNode {
	return &PlanGroupByNode{
		Input:    input,
		Metadata: Util.NewDefaultMetadata(),
		GroupBy:  NewGroupByNode(groupBy),
	}
}

func (self *PlanGroupByNode) SetOutput(output PlanNode) {
	self.Output = output
}

func (self *PlanGroupByNode) GetNodeType() PlanNodeType {
	return GROUPBYNODE
}

func (self *PlanGroupByNode) SetMetadata() (err error) {
	if err = self.Input.SetMetadata(); err != nil {
		return err
	}
	self.Metadata.Copy(self.Input.GetMetadata())
}

func (self *PlanGroupByNode) String() string {
	res := "PlanGroupByNode {\n"
	res += "Input: " + self.Input.String() + "\n"
	res += "GroupBy: " + fmt.Sprint(self.GroupBy) + "\n"
	res += "}/n"
	return res
}
