package Plan

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/Metadata"
)

type UnionType int32

const (
	_ UnionType = iota
	INTERSECT
	UNION
	EXCEPT
)

type PlanUnionNode struct {
	LeftInput  PlanNode
	RightInput PlanNode
	Output     PlanNode
	Operator   UnionType
	Metadata   *Metadata.Metadata
}

func NewPlanUnionNode(left, right PlanNode, op antlr.Token) *PlanUnionNode {
	var operator UnionType
	switch op.GetText() {
	case "INTERSECT":
		operator = INTERSECT
	case "UNION":
		operator = UNION
	case "EXCEPT":
		operator = EXCEPT
	}

	res := &PlanUnionNode{
		LeftInput:  left,
		RightInput: right,
		Operator:   operator,
		Metadata:   Metadata.NewMetadata(),
	}
	return res
}

func (self *PlanUnionNode) GetInputs() []PlanNode {
	return []PlanNode{self.LeftInput, self.RightInput}
}

func (self *PlanUnionNode) SetInputs(inputs []PlanNode) {
	self.LeftInput, self.RightInput = inputs[0], inputs[1]
}

func (self *PlanUnionNode) GetOutput() PlanNode {
	return self.Output
}

func (self *PlanUnionNode) SetOutput(output PlanNode) {
	self.Output = output
}

func (self *PlanUnionNode) GetNodeType() PlanNodeType {
	return UNIONNODE
}

func (self *PlanUnionNode) GetMetadata() *Metadata.Metadata {
	return self.Metadata
}

func (self *PlanUnionNode) SetMetadata() (err error) {
	if err = self.LeftInput.SetMetadata(); err != nil {
		return err
	}
	if err = self.RightInput.SetMetadata(); err != nil {
		return err
	}
	self.Metadata = self.LeftInput.GetMetadata().Copy()
	return nil
}

func (self *PlanUnionNode) String() string {
	res := "PlanUnionNode {\n"
	res += "LeftInput: " + self.LeftInput.String() + "\n"
	res += "RightInput: " + self.RightInput.String() + "\n"
	res += "Operator: " + fmt.Sprint(self.Operator) + "\n"
	res += "}\n"
	return res
}
