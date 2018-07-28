package optimizer

import (
	"github.com/xitongsys/guery/plan"
)

func DeleteRenameNode(node plan.PlanNode) error {
	if node == nil {
		return nil
	}
	switch node.(type) {
	case *plan.PlanRenameNode:
		nodea := node.(*plan.PlanRenameNode)
		if err := DeleteRenameNode(nodea.Input); err != nil {
			return err
		}
		parent := nodea.Output
		md := nodea.Input.GetMetadata()
		md.Rename(nodea.Rename)

		nodea.Input.SetOutput(parent)

		parInputs := parent.GetInputs()
		i := 0
		for i = 0; i < len(parInputs); i++ {
			if parInputs[i] == node {
				break
			}
		}
		parInputs[i] = nodea.Input
		parent.SetInputs(parInputs)

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
