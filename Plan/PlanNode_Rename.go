package Plan

import (
	"github.com/xitongsys/guery/Util"
)

type PlanRenameNode struct {
	Rename   string
	Metadata *Util.Metadata
	Input    PlanNode
	Output   PlanNode
}

func NewPlanRenameNode(input PlanNode, tname string) *PlanRenameNode {
	return &PlanRenameNode{
		Rename:   tname,
		Metadata: Util.NewDefaultMetadata(),
		Input:    input,
	}
}

func (self *PlanRenameNode) SetOutput(output PlanNode) {
	self.Output = output
}

func (self *PlanRenameNode) GetNodeType() PlanNodeType {
	return RENAMENODE
}

func (self *PlanRenameNode) GetMetadata() *Util.Metadata {
	return self.Metadata
}

func (self *PlanRenameNode) SetMetadata() (err error) {
	if err = self.Input.SetMetadata(); err != nil {
		return err
	}
	self.Metadata.Copy(self.Input.GetMetadata())
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
