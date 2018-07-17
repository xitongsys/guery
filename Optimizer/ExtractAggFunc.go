package Optimizer

import (
	"github.com/xitongsys/guery/Config"
	"github.com/xitongsys/guery/Plan"
)

func ExtractAggFunc(runtime *Config.ConfigRuntime, node Plan.PlanNode) error {
	if node == nil {
		return nil
	}
	switch node.(type) {
	case *Plan.PlanSelectNode:
		nodea := node.(*Plan.PlanSelectNode)
		if nodea.IsAggregate {
			funcs := []*Plan.FuncCallNode{}
			for _, item := range nodea.SelectItems {
				item.ExtractAggFunc(&funcs)
			}
			nodeLocal := Plan.NewPlanAggregateFuncLocalNode(runtime, funcs, nodea.Input)
			nodeGlobal := Plan.NewPlanAggregateFuncGlobalNode(runtime, funcs, nodeLocal)
			nodea.Input = nodeGlobal
			if err := nodea.SetMetadata(); err != nil {
				return err
			}
		}
	}

	for _, input := range node.GetInputs() {
		if err := ExtractAggFunc(runtime, input); err != nil {
			return err
		}
	}
	return nil
}
