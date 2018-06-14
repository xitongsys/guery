package Optimizer

import (
	"github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/Type"
)

func GetJoinKeys(leftInput, rightInput Plan.PlanNode, e *Plan.BooleanExpressionNode) (*Plan.ValueExpressionNode, *Plan.ValueExpressionNode, bool) {
	if e.Predicated == nil || e.Predicated.Predicate == nil || *(e.Predicated.Predicate.ComparisonOperator) != Type.EQ {
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
	leftMd, rightMd := leftInput.GetMetadata(), rightInput.GetMetadata()

	if leftMd.Contains(leftCols) && !leftMd.Contains(rightCols) &&
		rightMd.Contains(rightCols) && !rightMd.Contains(leftCols) {
		return leftExp, rightExp, true
	}

	if leftMd.Contains(rightCols) && !leftMd.Contains(leftCols) &&
		rightMd.Contains(leftCols) && !rightMd.Contains(rightCols) {
		return rightExp, leftExp, true
	}

	return nil, nil, false
}

func NewValueExpressionFromIdentifier(id *Plan.IdentifierNode) *Plan.ValueExpressionNode {
	return &Plan.ValueExpressionNode{
		PrimaryExpression: &Plan.PrimaryExpressionNode{
			Identifier: id,
		},
	}
}

func HashJoin(node Plan.PlanNode) error {
	if node == nil {
		return nil
	}

	switch node.(type) {
	case *Plan.PlanJoinNode:
		nodea := node.(*Plan.PlanJoinNode)
		inputs := nodea.GetInputs()
		leftInput, rightInput := inputs[0], inputs[1]
		leftKeys, rightKeys := []*Plan.ValueExpressionNode{}, []*Plan.ValueExpressionNode{}

		if nodea.JoinCriteria.BooleanExpression != nil { //JOIN ON...
			es := ExtractPredicates(nodea.JoinCriteria.BooleanExpression, Type.AND)
			for _, e := range es {
				leftExp, rightExp, ok := GetJoinKeys(leftInput, rightInput, e)
				if ok {
					leftKeys = append(leftKeys, leftExp)
					rightKeys = append(rightKeys, rightExp)
				}
			}

		} else { //JOIN USING (...)
			leftMd, rightMd := leftInput.GetMetadata(), rightInput.GetMetadata()
			for _, id := range nodea.JoinCriteria.Identifiers {
				if cols, err := id.GetColumns(); err != nil {
					return err
				} else {
					if leftMd.Contains(cols) && rightMd.Contains(cols) {
						leftKeys = append(leftKeys, NewValueExpressionFromIdentifier(id))
						rightKeys = append(rightKeys, NewValueExpressionFromIdentifier(id))
					}
				}
			}
		}

		if len(leftKeys) > 0 && len(leftKeys) == len(rightKeys) {
			hashJoinNode := Plan.NewPlanHashJoinNodeFromJoinNode(nodea, leftKeys, rightKeys)
			nodea.LeftInput.SetOutput(hashJoinNode)
			nodea.RightInput.SetOutput(hashJoinNode)
			parent := nodea.Output

			parInputs := parent.GetInputs()
			i := 0
			for i = 0; i < len(parInputs); i++ {
				if nodeb, ok := parInputs[i].(*Plan.PlanJoinNode); ok && nodeb == nodea {
					break
				}
			}
			parInputs[i] = hashJoinNode
			parent.SetInputs(parInputs)
			node = nodea
		}
	}

	for _, input := range node.GetInputs() {
		if err := HashJoin(input); err != nil {
			return err
		}
	}

	return nil
}
