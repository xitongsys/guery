package plan

import (
	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/metadata"
)

type PlanAggregateFuncGlobalNode struct {
	Input     PlanNode
	Output    PlanNode
	FuncNodes []*FuncCallNode
	Metadata  *Metadata.Metadata
}

func NewPlanAggregateFuncGlobalNode(runtime *Config.ConfigRuntime, funcs []*FuncCallNode, input PlanNode) *PlanAggregateFuncGlobalNode {
	return &PlanAggregateFuncGlobalNode{
		Input:     input,
		FuncNodes: funcs,
		Metadata:  Metadata.NewMetadata(),
	}
}

func (self *PlanAggregateFuncGlobalNode) GetInputs() []PlanNode {
	return []PlanNode{self.Input}
}

func (self *PlanAggregateFuncGlobalNode) SetInputs(inputs []PlanNode) {
	self.Input = inputs[0]
}

func (self *PlanAggregateFuncGlobalNode) GetOutput() PlanNode {
	return self.Output
}

func (self *PlanAggregateFuncGlobalNode) SetOutput(output PlanNode) {
	self.Output = output
}

func (self *PlanAggregateFuncGlobalNode) GetNodeType() PlanNodeType {
	return AGGREGATEFUNCGLOBALNODE
}

func (self *PlanAggregateFuncGlobalNode) GetMetadata() *Metadata.Metadata {
	return self.Metadata
}

func (self *PlanAggregateFuncGlobalNode) SetMetadata() (err error) {
	if err = self.Input.SetMetadata(); err != nil {
		return err
	}
	self.Metadata = self.Input.GetMetadata().Copy()
	return nil
}

func (self *PlanAggregateFuncGlobalNode) String() string {
	res := "PlanAggregateFuncGlobalNode {\n"
	res += self.Input.String()
	res += "}\n"
	return res
}
