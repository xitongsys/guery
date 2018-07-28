package plan

import (
	"fmt"

	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/gtype"
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/parser"
)

type PlanGroupByNode struct {
	Input    PlanNode
	Output   PlanNode
	Metadata *metadata.Metadata
	GroupBy  *GroupByNode
}

func NewPlanGroupByNode(runtime *config.ConfigRuntime, input PlanNode, groupBy parser.IGroupByContext) *PlanGroupByNode {
	return &PlanGroupByNode{
		Input:    input,
		Metadata: metadata.NewMetadata(),
		GroupBy:  NewGroupByNode(runtime, groupBy),
	}
}

func (self *PlanGroupByNode) GetInputs() []PlanNode {
	return []PlanNode{self.Input}
}

func (self *PlanGroupByNode) SetInputs(inputs []PlanNode) {
	self.Input = inputs[0]
}

func (self *PlanGroupByNode) GetOutput() PlanNode {
	return self.Output
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
	self.Metadata = self.Input.GetMetadata().Copy()
	self.Metadata.AppendKeyByType(gtype.STRING)
	return nil
}

func (self *PlanGroupByNode) GetMetadata() *metadata.Metadata {
	return self.Metadata
}

func (self *PlanGroupByNode) String() string {
	res := "PlanGroupByNode {\n"
	res += "Input: " + self.Input.String() + "\n"
	res += "GroupBy: " + fmt.Sprint(self.GroupBy) + "\n"
	res += "}/n"
	return res
}
