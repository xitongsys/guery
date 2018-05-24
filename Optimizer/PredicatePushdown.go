package Optimizer

import (
	"github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/Util"
)

func ExtractPredicates(node ParserNode, t Util.Operator) []ParserNode {
	res := []ParserNode{}
	switch node.(type) {
	case *Plan.BooleanExpressionNode:
		nodea := node.(*Plan.BooleanExpressionNode)
		if nodea.Predicated != nil {
			res = append(res, node)

		} else if nodea.NotBooleanExpression != nil {
			res = append(res, node)

		} else if self.BinaryBooleanExpressionNode != nil {
			leftNode := nodea.BinaryBooleanExpressionNode.LeftBooleanExpression
			rightNode := nodea.BinaryBooleanExpressionNode.RightBooleanExpression

			if (*nodea.Operator) == t {
				leftRes := ExtractPredicates(leftNode, t)
				rightRes := ExtractPredicates(rightNode, t)
				res = append(res, leftRes...)
				res = append(res, rightRes...)

			} else {
				res = append(res, node)
			}
		}
	}
	return res
}

func PredicatePushDown(node Plan.PlanNode, predicates []ParserNode) error {
	if node == nil {
		return nil
	}

	switch node.(type) {
	case *PlanFiliterNode:
		nodea := node.(*PlanFiliterNode)
		predicates = append(predicates, ExtractPredicates(nodea.BooleanExpressionNode, Util.AND)...)
		inputs := nodea.GetInputs()
		for _, input := range inputs {
			predicatesForInput := []ParserNode{}
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

	case *PlanSelectNode:
		nodea := node.(*PlanSelectNode)
		md := nodea.GetMetadata()
		output := nodea.GetOutput()
		if outputNode, ok := output.(*PlanFiliterNode); ok {
			for _, predicate := range predicates {
				cols, err := predicate.GetColumns()
				if err != nil {
					return err
				}
				if md.Contains(cols) {
					outputNode.AddBooleanExpressions(predicate)
				}
			}

		} else {

		}
		return nil

	case *PlanJoinNode:

	default:

	}
}
