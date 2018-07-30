package plan

import (
	"fmt"

	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/metadata"
)

type PlanDistinctGlobalNode struct {
	Input       PlanNode
	Output      PlanNode
	Metadata    *metadata.Metadata
	Expressions []*ExpressionNode
}

func NewPlanDistinctGlobalNode(runtime *config.ConfigRuntime, eps []*ExpressionNode, input PlanNode) *PlanDistinctGlobalNode {
	res := &PlanDistinctGlobalNode{
		Input:       input,
		Metadata:    metadata.NewMetadata(),
		Expressions: eps,
	}
	return res
}
func (self *PlanDistinctGlobalNode) GetInputs() []PlanNode {
	return []PlanNode{self.Input}
}

func (self *PlanDistinctGlobalNode) SetInputs(inputs []PlanNode) {
	self.Input = inputs[0]
}

func (self *PlanDistinctGlobalNode) GetOutput() PlanNode {
	return self.Output
}

func (self *PlanDistinctGlobalNode) SetOutput(output PlanNode) {
	self.Output = output
}

func (self *PlanDistinctGlobalNode) GetNodeType() PlanNodeType {
	return DISTINCTGLOBALNODE
}

func (self *PlanDistinctGlobalNode) SetMetadata() (err error) {
	if err = self.Input.SetMetadata(); err != nil {
		return err
	}
	self.Metadata = self.Input.GetMetadata().Copy()
	return nil
}

func (self *PlanDistinctGlobalNode) GetMetadata() *metadata.Metadata {
	return self.Metadata
}

func (self *PlanDistinctGlobalNode) String() string {
	res := "PlanDistinctGlobalNode {\n"
	res += "Input: " + self.Input.String() + "\n"
	res += "Expressions: " + fmt.Sprint(self.Expressions) + "\n"
	res += "}\n"
	return res
}

func (self *PlanDistinctGlobalNode) AddExpressions(nodes ...*ExpressionNode) {
	self.Expressions = append(self.Expressions, nodes...)
}
