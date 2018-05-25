package Plan

import (
	"fmt"

	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/parser"
)

type PlanFiliterNode struct {
	Input              PlanNode
	Output             PlanNode
	Metadata           *Util.Metadata
	BooleanExpressions []*BooleanExpressionNode
}

func NewPlanFiliterNode(input PlanNode, t parser.IBooleanExpressionContext) *PlanFiliterNode {
	res := &PlanFiliterNode{
		Input:              input,
		Metadata:           Util.NewMetadata(),
		BooleanExpressions: []*BooleanExpressionNode{NewBooleanExpressionNode(t)},
	}
	return res
}
func (self *PlanFiliterNode) GetInputs() []PlanNode {
	return []PlanNode{self.Input}
}

func (self *PlanFiliterNode) SetInputs(inputs []PlanNode) {
	self.Input = inputs[0]
}

func (self *PlanFiliterNode) GetOutput() PlanNode {
	return self.Output
}

func (self *PlanFiliterNode) SetOutput(output PlanNode) {
	self.Output = output
}

func (self *PlanFiliterNode) GetNodeType() PlanNodeType {
	return FILTERNODE
}

func (self *PlanFiliterNode) SetMetadata() (err error) {
	if err = self.Input.SetMetadata(); err != nil {
		return err
	}
	self.Metadata = self.Input.GetMetadata().Copy()
	return nil
}

func (self *PlanFiliterNode) GetMetadata() *Util.Metadata {
	return self.Metadata
}

func (self *PlanFiliterNode) String() string {
	res := "PlanFiliterNode {\n"
	res += "Input: " + self.Input.String() + "\n"
	res += "BooleanExpressions: " + fmt.Sprint(self.BooleanExpressions) + "\n"
	res += "}\n"
	return res
}

func (self *PlanFiliterNode) AddBooleanExpressions(nodes ...*BooleanExpressionNode) {
	self.BooleanExpressions = append(self.BooleanExpressions, nodes...)
}
