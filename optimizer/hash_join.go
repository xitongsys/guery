package optimizer

import (
	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/gtype"
	"github.com/xitongsys/guery/plan"
)

func GetJoinKeys(leftInput, rightInput plan.PlanNode, e *plan.BooleanExpressionNode) (*plan.ValueExpressionNode, *plan.ValueExpressionNode, bool) {
	if e.Predicated == nil || e.Predicated.Predicate == nil || *(e.Predicated.Predicate.ComparisonOperator) != gtype.EQ {
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

func NewValueExpressionFromIdentifier(runtime *config.ConfigRuntime, id *plan.IdentifierNode) *plan.ValueExpressionNode {
	return &plan.ValueExpressionNode{
		PrimaryExpression: &plan.PrimaryExpressionNode{
			Identifier: id,
		},
	}
}

func HashJoin(runtime *config.ConfigRuntime, node plan.PlanNode) error {
	if node == nil {
		return nil
	}

	switch node.(type) {
	case *plan.PlanJoinNode:
		nodea := node.(*plan.PlanJoinNode)
		inputs := nodea.GetInputs()
		leftInput, rightInput := inputs[0], inputs[1]
		leftKeys, rightKeys := []*plan.ValueExpressionNode{}, []*plan.ValueExpressionNode{}

		if nodea.JoinCriteria.BooleanExpression != nil { //JOIN ON...
			es := ExtractPredicates(nodea.JoinCriteria.BooleanExpression, gtype.AND)
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
						leftKeys = append(leftKeys, NewValueExpressionFromIdentifier(runtime, id))
						rightKeys = append(rightKeys, NewValueExpressionFromIdentifier(runtime, id))
					}
				}
			}
		}

		if len(leftKeys) > 0 && len(leftKeys) == len(rightKeys) {
			hashJoinNode := plan.NewPlanHashJoinNodeFromJoinNode(runtime, nodea, leftKeys, rightKeys)
			nodea.LeftInput.SetOutput(hashJoinNode)
			nodea.RightInput.SetOutput(hashJoinNode)
			parent := nodea.Output

			parInputs := parent.GetInputs()
			i := 0
			for i = 0; i < len(parInputs); i++ {
				if nodeb, ok := parInputs[i].(*plan.PlanJoinNode); ok && nodeb == nodea {
					break
				}
			}
			parInputs[i] = hashJoinNode
			parent.SetInputs(parInputs)
			node = nodea
		}
	}

	for _, input := range node.GetInputs() {
		if err := HashJoin(runtime, input); err != nil {
			return err
		}
	}

	return nil
}
