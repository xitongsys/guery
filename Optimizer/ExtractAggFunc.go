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
			nodeb := Plan.NewPlanAggregateFuncLocalNode(runtime, funcs, nodea.Input)
			nodeb.SetMetadata()
			nodea.Input = nodeb
			nodea.SetMetadata()
		}
	}

	for _, input := range node.GetInputs() {
		if err := ExtractAggFunc(input); err != nil {
			return err
		}
	}
	return nil
}
