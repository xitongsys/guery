package optimizer

import (
	"github.com/xitongsys/guery/gtype"
	"github.com/xitongsys/guery/plan"
)

func ExtractPredicates(node *plan.BooleanExpressionNode, t gtype.Operator) []*plan.BooleanExpressionNode {
	res := []*plan.BooleanExpressionNode{}
	if node.Predicated != nil {
		res = append(res, node)

	} else if node.NotBooleanExpression != nil {
		res = append(res, node)

	} else if node.BinaryBooleanExpression != nil {
		leftNode := node.BinaryBooleanExpression.LeftBooleanExpression
		rightNode := node.BinaryBooleanExpression.RightBooleanExpression

		if *(node.BinaryBooleanExpression.Operator) == t {
			leftRes := ExtractPredicates(leftNode, t)
			rightRes := ExtractPredicates(rightNode, t)
			res = append(res, leftRes...)
			res = append(res, rightRes...)

		} else {
			res = append(res, node)
		}
	}
	return res
}

func PredicatePushDown(node plan.PlanNode, predicates []*plan.BooleanExpressionNode) error {
	if node == nil {
		return nil
	}

	switch node.(type) {
	case *plan.PlanFilterNode:
		nodea := node.(*plan.PlanFilterNode)
		for _, be := range nodea.BooleanExpressions {
			predicates = append(predicates, ExtractPredicates(be, gtype.AND)...)
		}

		inputs := nodea.GetInputs()
		for _, input := range inputs {
			predicatesForInput := []*plan.BooleanExpressionNode{}
			for _, predicate := range predicates {
				md := input.GetMetadata()
				cols, err := predicate.GetColumns()
				if err != nil {
					return err
				}

				if md.Contains(cols) {
					predicatesForInput = append(predicatesForInput, predicate)
				}
			}
			if len(predicatesForInput) > 0 {
				if err := PredicatePushDown(input, predicatesForInput); err != nil {
					return err
				}
			}
		}

	case *plan.PlanSelectNode:
		nodea := node.(*plan.PlanSelectNode)
		md := nodea.GetMetadata()

		res := []*plan.BooleanExpressionNode{}
		for _, predicate := range predicates {
			cols, err := predicate.GetColumns()
			if err != nil {
				return err
			}
			if md.Contains(cols) {
				res = append(res, predicate)
			}
		}
		if len(res) > 0 {
			output := nodea.GetOutput()
			if _, ok := output.(*plan.PlanFilterNode); !ok {
				newFilterNode := &plan.PlanFilterNode{
					Input:              node,
					Output:             output,
					Metadata:           node.GetMetadata().Copy(),
					BooleanExpressions: []*plan.BooleanExpressionNode{},
				}
				output.SetInputs([]plan.PlanNode{newFilterNode})
				node.SetOutput(newFilterNode)
			}
			outputNode := nodea.GetOutput().(*plan.PlanFilterNode)
			outputNode.AddBooleanExpressions(res...)
		}

		for _, input := range node.GetInputs() {
			PredicatePushDown(input, []*plan.BooleanExpressionNode{})
		}

		return nil

	case *plan.PlanScanNode:
		nodea := node.(*plan.PlanScanNode)
		md := node.GetMetadata()
		for _, predicate := range predicates {
			cols, err := predicate.GetColumns()
			if err != nil {
				return err
			}
			if md.Contains(cols) {
				nodea.Filters = append(nodea.Filters, predicate)
			}
		}
	case *plan.PlanShowNode:
		return nil

	default:
		inputs := node.GetInputs()
		for _, input := range inputs {
			if len(predicates) <= 0 {
				if err := PredicatePushDown(input, predicates); err != nil {
					return err
				}
				continue
			}
			predicatesForInput := []*plan.BooleanExpressionNode{}
			for _, predicate := range predicates {
				md := input.GetMetadata()
				cols, err := predicate.GetColumns()
				if err != nil {
					return err
				}

				if md.Contains(cols) {
					predicatesForInput = append(predicatesForInput, predicate)
				}
			}

			if err := PredicatePushDown(input, predicatesForInput); err != nil {
				return err
			}
		}
	}
	return nil
}
