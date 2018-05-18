package Plan

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/Util"
)

type PlanLimitNode struct {
	Input       PlanNode
	Output      PlanNode
	Metadata    *Util.Metadata
	LimitNumber *int64
}

func NewPlanLimitNode(input PlanNode, t antlr.TerminalNode) *PlanLimitNode {
	res := &PlanLimitNode{
		Input:    input,
		Metadata: Util.NewMetadata(),
	}
	if ns := t.GetText(); ns != "ALL" {
		var num int64
		fmt.Sscanf(ns, "%d", &num)
		res.LimitNumber = &num
	}
	return res
}

func (self *PlanLimitNode) GetInputs() []PlanNode {
	return []PlanNode{self.Input}
}

func (self *PlanLimitNode) SetInputs(inputs []PlanNode) {
	self.Input = inputs[0]
}

func (self *PlanLimitNode) GetOutput() PlanNode {
	return self.Output
}

func (self *PlanLimitNode) SetOutput(output PlanNode) {
	self.Output = output
}

func (self *PlanLimitNode) GetNodeType() PlanNodeType {
	return LIMITNODE
}

func (self *PlanLimitNode) GetMetadata() *Util.Metadata {
	return self.Metadata
}

func (self *PlanLimitNode) SetMetadata() error {
	err := self.Input.SetMetadata()
	if err != nil {
		return err
	}
	self.Metadata = self.Input.GetMetadata().Copy()
	return nil
}

func (self *PlanLimitNode) String() string {
	res := "PlanLimitNode {\n"
	res += "Input: " + self.Input.String() + "\n"
	res += "LimitNubmer: " + fmt.Sprint(*self.LimitNumber) + "\n"
	res += "}\n"
	return res
}
