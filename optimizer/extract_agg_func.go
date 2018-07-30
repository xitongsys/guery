package optimizer

import (
	"fmt"
	"math/rand"

	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/gtype"
	"github.com/xitongsys/guery/plan"
)

func ExtractDistinctExpressions(funcs []*plan.FuncCallNode) []*plan.ExpressionNode {
	res := []*plan.ExpressionNode{}
	for _, f := range funcs {
		if f.SetQuantifier != nil && (*f.SetQuantifier) == gtype.DISTINCT {
			res = append(res, f.Expressions...)
			colName := fmt.Sprintf("DIST_%v_%v", len(res), rand.Int())
			f.Expressions[0].Name = colName

			f.Expressions = []*plan.ExpressionNode{
				&plan.ExpressionNode{
					Name: colName,
					BooleanExpression: &plan.BooleanExpressionNode{
						Name: colName,
						Predicated: &plan.PredicatedNode{
							Name: colName,
							ValueExpression: &plan.ValueExpressionNode{
								Name: colName,
								PrimaryExpression: &plan.PrimaryExpressionNode{
									Name: colName,
									Identifier: &plan.IdentifierNode{
										Str: &colName,
									},
								},
							},
						},
					},
				},
			}
		}
	}
	return res
}

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

			var nodeLocal *plan.PlanAggregateFuncLocalNode

			//for distinct
			distEps := ExtractDistinctExpressions(funcs)
			if len(distEps) > 0 {
				distLocalNode := plan.NewPlanDistinctLocalNode(runtime, distEps, nodea.Input)
				distGlobalNode := plan.NewPlanDistinctGlobalNode(runtime, distEps, distLocalNode)
				nodeLocal = plan.NewPlanAggregateFuncLocalNode(runtime, funcs, distGlobalNode)
			} else {
				nodeLocal = plan.NewPlanAggregateFuncLocalNode(runtime, funcs, nodea.Input)
			}

			nodeLocal.SetMetadata()

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
