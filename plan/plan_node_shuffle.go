package plan

import (
	"fmt"

	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/parser"
)

type PlanShuffleNode struct {
	Input              PlanNode
	Output             PlanNode
	Metadata           *metadata.Metadata
	BooleanExpressions []*BooleanExpressionNode
}

func NewPlanShuffleNode(runtime *config.ConfigRuntime, input PlanNode, t parser.IBooleanExpressionContext) *PlanShuffleNode {
	res := &PlanShuffleNode{
		Input:              input,
		Metadata:           metadata.NewMetadata(),
		BooleanExpressions: []*BooleanExpressionNode{NewBooleanExpressionNode(runtime, t)},
	}
	return res
}
func (self *PlanShuffleNode) GetInputs() []PlanNode {
	return []PlanNode{self.Input}
}

func (self *PlanShuffleNode) SetInputs(inputs []PlanNode) {
	self.Input = inputs[0]
}

func (self *PlanShuffleNode) GetOutput() PlanNode {
	return self.Output
}

func (self *PlanShuffleNode) SetOutput(output PlanNode) {
	self.Output = output
}

func (self *PlanShuffleNode) GetNodeType() PlanNodeType {
	return SHUFFLENODE
}

func (self *PlanShuffleNode) SetMetadata() (err error) {
	if err = self.Input.SetMetadata(); err != nil {
		return err
	}
	self.Metadata = self.Input.GetMetadata().Copy()
	return nil
}

func (self *PlanShuffleNode) GetMetadata() *metadata.Metadata {
	return self.Metadata
}

func (self *PlanShuffleNode) String() string {
	res := "PlanShuffleNode {\n"
	res += "Input: " + self.Input.String() + "\n"
	res += "BooleanExpressions: " + fmt.Sprint(self.BooleanExpressions) + "\n"
	res += "}\n"
	return res
}

func (self *PlanShuffleNode) AddBooleanExpressions(nodes ...*BooleanExpressionNode) {
	self.BooleanExpressions = append(self.BooleanExpressions, nodes...)
}
