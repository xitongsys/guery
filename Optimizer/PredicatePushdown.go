package Optimizer

import (
	"github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/Util"
)

func ExtractPredicates(node Plan.PlanNode, t Util.Operator) []Plan.PlanNode {
	res := []PlanNode{}
	if nodea, ok := node.(*Plan.BinaryBooleanExpressionNode); ok && *(nodea.Operator) == t {
		leftRes := ExtractPredicates(nodea.LeftBooleanExpression)
		rightRes := ExtractPredicates(nodea.RightBooleanExpression)
		res = append(res, leftRes...)
		res = append(res, rightRes...)
		return res
	}
	return []Plan.PlanNode{node}
}

func PrecidatePushDown(node Plan.PlanNode) error {
	if node == nil {
		return nil
	}

	switch node.(type) {
	case *PlanFiliterNode:
		nodea := node.(*PlanFiliterNode)
		predicates := ExtractPredicates(nodea.BooleanExpressionNode, Util.AND)
	default:

	}
}
