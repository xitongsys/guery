package Plan

import (
	"github.com/xitongsys/guery/Config"
	"github.com/xitongsys/guery/Metadata"
)

type PlanAggregateFuncLocalNode struct {
	Input    PlanNode
	Output   PlanNode
	Funcs    map[string][]*ExpressionNode
	Metadata *Metadata.Metadata
}

func NewPlanAggregateFuncLocalNode(runtime *Config.ConfigRuntime, funcs map[string][]*ExpressionNode, input PlanNode) *PlanAggregateNode {
	return &PlanAggregateFuncLocalNode{
		Input:    input,
		Funcs:    funcs,
		Metadata: Metadata.NewMetadata(),
	}
}

func (self *PlanAggregateFuncLocalNode) GetInputs() []PlanNode {
	return []PlanNode{self.Input}
}

func (self *PlanAggregateFuncLocalNode) SetInputs(inputs []PlanNode) {
	self.Input = inputs[0]
}

func (self *PlanAggregateFuncLocalNode) GetOutput() PlanNode {
	return self.Output
}

func (self *PlanAggregateFuncLocalNode) SetOutput(output PlanNode) {
	self.Output = output
}

func (self *PlanAggregateFuncLocalNode) GetNodeType() PlanNodeType {
	return AGGREGATEFUNCLOCALNODE
}

func (self *PlanAggregateFuncLocalNode) GetMetadata() *Metadata.Metadata {
	return self.Metadata
}

func (self *PlanAggregateFuncLocalNode) SetMetadata() (err error) {
	if err = self.Input.SetMetadata(); err != nil {
		return err
	}
	self.Metadata = self.Input.GetMetadata().Copy()

	return nil
}

func (self *PlanAggregateFuncLocalNode) String() string {
	res := "PlanAggregateFuncLocalNode {\n"
	res += self.Input.String()
	res += "}\n"
	return res
}
