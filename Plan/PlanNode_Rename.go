package Plan

import (
	"github.com/xitongsys/guery/Config"
	"github.com/xitongsys/guery/Metadata"
)

type PlanRenameNode struct {
	Rename   string
	Metadata *Metadata.Metadata
	Input    PlanNode
	Output   PlanNode
}

func NewPlanRenameNode(runtime *Config.ConfigRuntime, input PlanNode, tname string) *PlanRenameNode {
	return &PlanRenameNode{
		Rename:   tname,
		Metadata: Metadata.NewMetadata(),
		Input:    input,
	}
}

func (self *PlanRenameNode) GetInputs() []PlanNode {
	return []PlanNode{self.Input}
}

func (self *PlanRenameNode) SetInputs(inputs []PlanNode) {
	self.Input = inputs[0]
}

func (self *PlanRenameNode) GetOutput() PlanNode {
	return self.Output
}

func (self *PlanRenameNode) SetOutput(output PlanNode) {
	self.Output = output
}

func (self *PlanRenameNode) GetNodeType() PlanNodeType {
	return RENAMENODE
}

func (self *PlanRenameNode) GetMetadata() *Metadata.Metadata {
	return self.Metadata
}

func (self *PlanRenameNode) SetMetadata() (err error) {
	if err = self.Input.SetMetadata(); err != nil {
		return err
	}
	self.Metadata = self.Input.GetMetadata().Copy()
	self.Metadata.Rename(self.Rename)
	return nil
}

func (self *PlanRenameNode) String() string {
	res := "PlanRenameNode {\n"
	res += "Rename: " + self.Rename + "\n"
	res += "Input: " + self.Input.String() + "\n"
	res += "}\n"
	return res
}
