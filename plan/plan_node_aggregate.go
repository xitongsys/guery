package plan

import (
	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/metadata"
)

type PlanAggregateNode struct {
	Input    PlanNode
	Output   PlanNode
	Metadata *metadata.Metadata
}

func NewPlanAggregateNode(runtime *config.ConfigRuntime, input PlanNode) *PlanAggregateNode {
	return &PlanAggregateNode{
		Input:    input,
		Metadata: metadata.NewMetadata(),
	}
}

func (self *PlanAggregateNode) GetInputs() []PlanNode {
	return []PlanNode{self.Input}
}

func (self *PlanAggregateNode) SetInputs(inputs []PlanNode) {
	self.Input = inputs[0]
}

func (self *PlanAggregateNode) GetOutput() PlanNode {
	return self.Output
}

func (self *PlanAggregateNode) SetOutput(output PlanNode) {
	self.Output = output
}

func (self *PlanAggregateNode) GetNodeType() PlanNodeType {
	return AGGREGATENODE
}

func (self *PlanAggregateNode) GetMetadata() *metadata.Metadata {
	return self.Metadata
}

func (self *PlanAggregateNode) SetMetadata() (err error) {
	if err = self.Input.SetMetadata(); err != nil {
		return err
	}
	self.Metadata = self.Input.GetMetadata().Copy()

	return nil
}

func (self *PlanAggregateNode) String() string {
	res := "PlanAggregateNode {\n"
	res += self.Input.String()
	res += "}\n"
	return res
}
