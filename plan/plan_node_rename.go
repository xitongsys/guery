package plan

import (
	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/metadata"
)

type PlanRenameNode struct {
	Rename   string
	Metadata *metadata.Metadata
	Input    PlanNode
	Output   PlanNode
}

func NewPlanRenameNode(runtime *config.ConfigRuntime, input PlanNode, tname string) *PlanRenameNode {
	return &PlanRenameNode{
		Rename:   tname,
		Metadata: metadata.NewMetadata(),
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

func (self *PlanRenameNode) GetMetadata() *metadata.Metadata {
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
