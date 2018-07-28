package plan

import (
	"fmt"

	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/parser"
)

type PlanFilterNode struct {
	Input              PlanNode
	Output             PlanNode
	Metadata           *Metadata.Metadata
	BooleanExpressions []*BooleanExpressionNode
}

func NewPlanFilterNode(runtime *Config.ConfigRuntime, input PlanNode, t parser.IBooleanExpressionContext) *PlanFilterNode {
	res := &PlanFilterNode{
		Input:              input,
		Metadata:           Metadata.NewMetadata(),
		BooleanExpressions: []*BooleanExpressionNode{NewBooleanExpressionNode(runtime, t)},
	}
	return res
}
func (self *PlanFilterNode) GetInputs() []PlanNode {
	return []PlanNode{self.Input}
}

func (self *PlanFilterNode) SetInputs(inputs []PlanNode) {
	self.Input = inputs[0]
}

func (self *PlanFilterNode) GetOutput() PlanNode {
	return self.Output
}

func (self *PlanFilterNode) SetOutput(output PlanNode) {
	self.Output = output
}

func (self *PlanFilterNode) GetNodeType() PlanNodeType {
	return FILTERNODE
}

func (self *PlanFilterNode) SetMetadata() (err error) {
	if err = self.Input.SetMetadata(); err != nil {
		return err
	}
	self.Metadata = self.Input.GetMetadata().Copy()
	return nil
}

func (self *PlanFilterNode) GetMetadata() *Metadata.Metadata {
	return self.Metadata
}

func (self *PlanFilterNode) String() string {
	res := "PlanFilterNode {\n"
	res += "Input: " + self.Input.String() + "\n"
	res += "BooleanExpressions: " + fmt.Sprint(self.BooleanExpressions) + "\n"
	res += "}\n"
	return res
}

func (self *PlanFilterNode) AddBooleanExpressions(nodes ...*BooleanExpressionNode) {
	self.BooleanExpressions = append(self.BooleanExpressions, nodes...)
}
