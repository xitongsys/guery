package plan

import (
	"fmt"

	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/metadata"
)

type PlanDistinctNode struct {
	Input       PlanNode
	Output      PlanNode
	Metadata    *metadata.Metadata
	Expressions []*ExpressionNode
	Names       []string
}

func NewPlanDistinctNode(runtime *config.ConfigRuntime, input PlanNode, eps []*ExpressionNode) *PlanDistinctNode {
	res := &PlanDistinctNode{
		Input:       input,
		Metadata:    metadata.NewMetadata(),
		Expressions: eps,
	}
	return res
}
func (self *PlanDistinctNode) GetInputs() []PlanNode {
	return []PlanNode{self.Input}
}

func (self *PlanDistinctNode) SetInputs(inputs []PlanNode) {
	self.Input = inputs[0]
}

func (self *PlanDistinctNode) GetOutput() PlanNode {
	return self.Output
}

func (self *PlanDistinctNode) SetOutput(output PlanNode) {
	self.Output = output
}

func (self *PlanDistinctNode) GetNodeType() PlanNodeType {
	return DISTINCTNODE
}

func (self *PlanDistinctNode) SetMetadata() (err error) {
	if err = self.Input.SetMetadata(); err != nil {
		return err
	}
	self.Metadata = self.Input.GetMetadata().Copy()
	for _, e := range self.Expressions {
		t, err := e.GetType(self.Input.GetMetadata())
		if err != nil {
			return err
		}
		col := metadata.NewColumnMetadata(t, e.Name)
		self.Metadata.AppendColumn(col)
	}
	return nil
}

func (self *PlanDistinctNode) GetMetadata() *metadata.Metadata {
	return self.Metadata
}

func (self *PlanDistinctNode) String() string {
	res := "PlanDistinctNode {\n"
	res += "Input: " + self.Input.String() + "\n"
	res += "Expressions: " + fmt.Sprint(self.Expressions) + "\n"
	res += "}\n"
	return res
}

func (self *PlanDistinctNode) AddExpressions(nodes ...*ExpressionNode) {
	self.Expressions = append(self.Expressions, nodes...)
}
