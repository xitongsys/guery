package Optimizer

import (
	"github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/Util"
)

func GetJoinKeys(e *BooleanExpressionNode) (*BooleanExpressionNode, *BooleanExpressionNode, bool) {
	if e.Predicated == nil || e.Predicated.Predicate == nil || e.Predicated.Predicate.ComparsionOperator != Util.EQ {
		return nil, nil, false
	}
	leftExp, rightExp := e.Predicated.ValueExpression, e.Predicated.Predicate.RightValueExpression
	leftCols, err := leftExp.GetColumns()
	if err != nil {
		return nil, nil, false
	}
	rightCols, err := rightExp.GetColumns()
	if err != nil {
		return nil, nil, false
	}

}

func HashJoin(node Plan.PlanNode) error {
	if node == nil {
		return nil
	}

	switch node.(type) {
	case *Plan.PlanJoinNode:
		nodea := node.(*Plan.PlanJoinNode)
		if nodea.JoinCriteria.BooleanExpression != nil { //JOIN ON...
			es := ExtractPredicates(nodea.JoinCriteria.BooleanExpression)

		} else { //JOIN USING (...)
		}

	default:
		for _, input := range node.GetInputs() {
			if err := HashJoin(input); err != nil {
				return err
			}
		}
	}
	return nil
}
