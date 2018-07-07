package Optimizer

import (
	"github.com/xitongsys/guery/Plan"
)

func ExtractAggFunc(node Plan.PlanNode) error {
	if node == nil {
		return nil
	}
	switch node.(type) {
	case *Plan.PlanSelectNode:
		nodea := node.(*Plan.PlanSelectNode)
		if nodea.IsAggregate {
		}
	}

	for _, input := range node.GetInputs() {
		if err := DeleteRenameNode(input); err != nil {
			return err
		}
	}
	return nil
}
