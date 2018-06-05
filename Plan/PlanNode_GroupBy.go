package Plan

import (
	"fmt"

	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/parser"
)

type PlanGroupByNode struct {
	Input    PlanNode
	Output   PlanNode
	Metadata *Util.Metadata
	GroupBy  *GroupByNode
}

func NewPlanGroupByNode(input PlanNode, groupBy parser.IGroupByContext, having parser.IBooleanExpressionContext) *PlanGroupByNode {
	return &PlanGroupByNode{
		Input:    input,
		Metadata: Util.NewMetadata(),
		GroupBy:  NewGroupByNode(groupBy, having),
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
	self.Metadata.ClearKeys()
	for _, e := range self.GroupBy.GroupingElements {
		t, err := e.GetType(self.Metadata)
		if err != nil {
			return err
		}
		self.Metadata.AppendKeyByType(t)
	}
	return nil
}

func (self *PlanGroupByNode) GetMetadata() *Util.Metadata {
	return self.Metadata
}

func (self *PlanGroupByNode) String() string {
	res := "PlanGroupByNode {\n"
	res += "Input: " + self.Input.String() + "\n"
	res += "GroupBy: " + fmt.Sprint(self.GroupBy) + "\n"
	res += "}/n"
	return res
}
