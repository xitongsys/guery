package plan

import (
	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/metadata"
)

type PlanAggregateFuncLocalNode struct {
	Input     PlanNode
	Output    PlanNode
	FuncNodes []*FuncCallNode
	Metadata  *metadata.Metadata
}

func NewPlanAggregateFuncLocalNode(runtime *config.ConfigRuntime, funcs []*FuncCallNode, input PlanNode) *PlanAggregateFuncLocalNode {
	return &PlanAggregateFuncLocalNode{
		Input:     input,
		FuncNodes: funcs,
		Metadata:  metadata.NewMetadata(),
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

func (self *PlanAggregateFuncLocalNode) GetMetadata() *metadata.Metadata {
	return self.Metadata
}

func (self *PlanAggregateFuncLocalNode) SetMetadata() (err error) {
	if err = self.Input.SetMetadata(); err != nil {
		return err
	}
	self.Metadata = self.Input.GetMetadata().Copy()
	for _, f := range self.FuncNodes {
		t, err := f.GetType(self.Input.GetMetadata())
		if err != nil {
			return err
		}
		col := metadata.NewColumnMetadata(t, f.ResColName)
		self.Metadata.AppendColumn(col)
	}

	return nil
}

func (self *PlanAggregateFuncLocalNode) String() string {
	res := "PlanAggregateFuncLocalNode {\n"
	res += self.Input.String()
	res += "}\n"
	return res
}
