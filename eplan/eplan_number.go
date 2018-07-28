package eplan

import (
	"fmt"

	"github.com/xitongsys/guery/logger"
	. "github.com/xitongsys/guery/plan"
)

func GetEPlanExecutorNumber(node PlanNode, pn int32) (int32, error) {
	res, _, err := getEPlanExecutorNumber(node, pn)
	if err != nil {
		return -1, err
	}
	return res + 1, nil
}

func getEPlanExecutorNumber(node PlanNode, pn int32) (int32, int32, error) {
	switch node.(type) {
	case *PlanShowNode:
		return 1, 1, nil

	case *PlanScanNode:
		return pn, pn, nil

	case *PlanSelectNode:
		nodea := node.(*PlanSelectNode)
		res, cur, err := getEPlanExecutorNumber(nodea.Input, pn)
		if err != nil {
			return -1, -1, err
		}
		return res + cur, cur, nil

	case *PlanGroupByNode:
		nodea := node.(*PlanGroupByNode)
		res, cur, err := getEPlanExecutorNumber(nodea.Input, pn)
		if err != nil {
			return -1, -1, err
		}
		return res + cur, cur, nil

	case *PlanJoinNode:
		nodea := node.(*PlanJoinNode)
		res1, _, err1 := getEPlanExecutorNumber(nodea.LeftInput, pn)
		if err1 != nil {
			return -1, -1, err1
		}
		res2, _, err2 := getEPlanExecutorNumber(nodea.RightInput, pn)
		if err2 != nil {
			return -1, -1, err2
		}
		return res1 + res2 + 1 + pn, pn, nil

	case *PlanHashJoinNode:
		nodea := node.(*PlanHashJoinNode)
		res1, cur1, err1 := getEPlanExecutorNumber(nodea.LeftInput, pn)
		if err1 != nil {
			return -1, -1, err1
		}
		res2, cur2, err2 := getEPlanExecutorNumber(nodea.RightInput, pn)
		if err2 != nil {
			return -1, -1, err2
		}
		return res1 + res2 + cur1 + cur2 + pn, pn, nil

	case *PlanLimitNode:
		nodea := node.(*PlanLimitNode)
		res, cur, err := getEPlanExecutorNumber(nodea.Input, pn)
		if err != nil {
			return -1, -1, err
		}
		return res + cur, cur, nil

	case *PlanAggregateNode:
		nodea := node.(*PlanAggregateNode)
		res, _, err := getEPlanExecutorNumber(nodea.Input, pn)
		if err != nil {
			return -1, -1, err
		}
		return res + 1, 1, nil

	case *PlanAggregateFuncGlobalNode:
		nodea := node.(*PlanAggregateFuncGlobalNode)
		res, _, err := getEPlanExecutorNumber(nodea.Input, pn)
		if err != nil {
			return -1, -1, err
		}
		return res + 1, 1, nil

	case *PlanAggregateFuncLocalNode:
		nodea := node.(*PlanAggregateFuncLocalNode)
		res, cur, err := getEPlanExecutorNumber(nodea.Input, pn)
		if err != nil {
			return -1, -1, err
		}
		return res + cur, cur, nil

	case *PlanOrderByNode:
		nodea := node.(*PlanOrderByNode)
		res, cur, err := getEPlanExecutorNumber(nodea.Input, pn)
		if err != nil {
			return -1, -1, err
		}
		return res + cur + 1, 1, nil

	case *PlanUnionNode:
		nodea := node.(*PlanUnionNode)
		leftRes, curl, err := getEPlanExecutorNumber(nodea.LeftInput, pn)
		if err != nil {
			return -1, -1, err
		}
		rightRes, curr, err := getEPlanExecutorNumber(nodea.RightInput, pn)
		if err != nil {
			return -1, -1, err
		}
		return leftRes + rightRes + curl, curr, nil

	case *PlanFilterNode:
		nodea := node.(*PlanFilterNode)
		res, cur, err := getEPlanExecutorNumber(nodea.Input, pn)
		if err != nil {
			return -1, -1, err
		}
		return res + cur, cur, nil

	default:
		Logger.Errorf("getEPlanExecutorNumber: unknown node type")
		return -1, -1, fmt.Errorf("getEPlanExecutorNumber: unknown node type")
	}
}
