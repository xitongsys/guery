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
			funcsGlobal := make([]*Plan.FuncCallNode, len(funcs))
			for i, f := range funcs {
				funcsGlobal[i] = &Plan.FuncCallNode{
					FuncName:   f.FuncName + "GLOBAL",
					ResColName: f.ResColName,
					Expressions: []*Plan.ExpressionNode{
						&Plan.ExpressionNode{
							Name: f.ResColName,
							BooleanExpression: &Plan.BooleanExpressionNode{
								Name: f.ResColName,
								Predicated: &Plan.PredicatedNode{
									Name: f.ResColName,
									ValueExpression: &Plan.ValueExpressionNode{
										Name: f.ResColName,
										PrimaryExpression: &Plan.PrimaryExpressionNode{
											Name: f.ResColName,
											Identifier: &Plan.IdentifierNode{
												Str: &f.ResColName,
											},
										},
									},
								},
							},
						},
					},
				}
			}
			nodeGlobal := Plan.NewPlanAggregateFuncGlobalNode(runtime, funcsGlobal, nodeLocal)
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
