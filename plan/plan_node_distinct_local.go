package plan

import (
	"fmt"

	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/metadata"
)

type PlanDistinctLocalNode struct {
	Input       PlanNode
	Output      PlanNode
	Metadata    *metadata.Metadata
	Expressions []*ExpressionNode
}

func NewPlanDistinctLocalNode(runtime *config.ConfigRuntime, eps []*ExpressionNode, input PlanNode) *PlanDistinctLocalNode {
	res := &PlanDistinctLocalNode{
		Input:       input,
		Metadata:    metadata.NewMetadata(),
		Expressions: eps,
	}
	return res
}
func (self *PlanDistinctLocalNode) GetInputs() []PlanNode {
	return []PlanNode{self.Input}
}

func (self *PlanDistinctLocalNode) SetInputs(inputs []PlanNode) {
	self.Input = inputs[0]
}

func (self *PlanDistinctLocalNode) GetOutput() PlanNode {
	return self.Output
}

func (self *PlanDistinctLocalNode) SetOutput(output PlanNode) {
	self.Output = output
}

func (self *PlanDistinctLocalNode) GetNodeType() PlanNodeType {
	return DISTINCTLOCALNODE
}

func (self *PlanDistinctLocalNode) SetMetadata() (err error) {
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

func (self *PlanDistinctLocalNode) GetMetadata() *metadata.Metadata {
	return self.Metadata
}

func (self *PlanDistinctLocalNode) String() string {
	res := "PlanDistinctLocalNode {\n"
	res += "Input: " + self.Input.String() + "\n"
	res += "Expressions: " + fmt.Sprint(self.Expressions) + "\n"
	res += "}\n"
	return res
}

func (self *PlanDistinctLocalNode) AddExpressions(nodes ...*ExpressionNode) {
	self.Expressions = append(self.Expressions, nodes...)
}
