package optimizer

import (
	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/plan"
)

func ExtractAggFunc(runtime *config.ConfigRuntime, node plan.PlanNode) error {
	if node == nil {
		return nil
	}
	switch node.(type) {
	case *plan.PlanSelectNode:
		nodea := node.(*plan.PlanSelectNode)
		if nodea.IsAggregate {
			funcs := []*plan.FuncCallNode{}
			for _, item := range nodea.SelectItems {
				item.ExtractAggFunc(&funcs)
			}
			if nodea.Having != nil {
				nodea.Having.ExtractAggFunc(&funcs)
			}

			nodeLocal := plan.NewPlanAggregateFuncLocalNode(runtime, funcs, nodea.Input)
			funcsGlobal := make([]*plan.FuncCallNode, len(funcs))
			for i, f := range funcs {
				funcsGlobal[i] = &plan.FuncCallNode{
					FuncName:   f.FuncName + "GLOBAL",
					ResColName: f.ResColName,
					Expressions: []*plan.ExpressionNode{
						&plan.ExpressionNode{
							Name: f.ResColName,
							BooleanExpression: &plan.BooleanExpressionNode{
								Name: f.ResColName,
								Predicated: &plan.PredicatedNode{
									Name: f.ResColName,
									ValueExpression: &plan.ValueExpressionNode{
										Name: f.ResColName,
										PrimaryExpression: &plan.PrimaryExpressionNode{
											Name: f.ResColName,
											Identifier: &plan.IdentifierNode{
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
			nodeGlobal := plan.NewPlanAggregateFuncGlobalNode(runtime, funcsGlobal, nodeLocal)
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
