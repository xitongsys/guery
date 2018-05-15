package Optimizer

import (
	"github.com/xitongsys/guery/Plan"
)

func DeleteRenameNode(node Plan.PlanNode) error {
	if node == nil {
		return nil
	}
	switch node.(type) {
	case *Plan.PlanRenameNode:
		nodea := node.(*Plan.PlanRenameNode)
		if err := DeleteRenameNode(nodea.Input); err != nil {
			return err
		}
		parent := nodea.Output
		md := nodea.Input.GetMetadata()
		md.Rename(nodea.Rename)

		nodea.Input.SetOutput(parent)
		parent.SetInputs([]Plan.PlanNode{nodea.Input})

		return nil

	default:
		for _, input := range node.GetInputs() {
			if err := DeleteRenameNode(input); err != nil {
				return err
			}
		}
	}
	return nil
}
