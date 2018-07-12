package EPlan

import (
	"fmt"

	"github.com/xitongsys/guery/Logger"
	. "github.com/xitongsys/guery/Plan"
)

func GetEPlanExecutorNumber(node PlanNode, pn int32) (int32, error) {
	res, err := getEPlanExecutorNumber(node, pn)
	if err != nil {
		return -1, err
	}
	return res + 1, nil
}

func getEPlanExecutorNumber(node PlanNode, pn int32) (int32, error) {
	switch node.(type) {
	case *PlanShowNode:
		return 1, nil

	case *PlanScanNode:
		return pn, nil

	case *PlanSelectNode:
		nodea := node.(*PlanSelectNode)
		res, err := getEPlanExecutorNumber(nodea.Input, pn)
		if err != nil {
			return -1, err
		}
		return res + pn, nil

	case *PlanGroupByNode:
		nodea := node.(*PlanGroupByNode)
		res, err := getEPlanExecutorNumber(nodea.Input, pn)
		if err != nil {
			return -1, err
		}
		return res + 1 + pn, nil

	case *PlanJoinNode:
		nodea := node.(*PlanJoinNode)
		res1, err1 := getEPlanExecutorNumber(nodea.LeftInput, pn)
		if err1 != nil {
			return -1, err1
		}
		res2, err2 := getEPlanExecutorNumber(nodea.RightInput, pn)
		if err2 != nil {
			return -1, err2
		}
		return res1 + res2 + 1 + pn, nil

	case *PlanHashJoinNode:
		nodea := node.(*PlanHashJoinNode)
		res1, err1 := getEPlanExecutorNumber(nodea.LeftInput, pn)
		if err1 != nil {
			return -1, err1
		}
		res2, err2 := getEPlanExecutorNumber(nodea.RightInput, pn)
		if err2 != nil {
			return -1, err2
		}
		return res1 + res2 + 2 + pn, nil

	case *PlanLimitNode:
		nodea := node.(*PlanLimitNode)
		res, err := getEPlanExecutorNumber(nodea.Input, pn)
		if err != nil {
			return -1, err
		}
		return res + pn, nil

	case *PlanAggregateNode:
		nodea := node.(*PlanAggregateNode)
		res, err := getEPlanExecutorNumber(nodea.Input, pn)
		if err != nil {
			return -1, err
		}
		return res + 1, nil

	case *PlanAggregateFuncLocalNode:
		nodea := node.(*PlanAggregateFuncLocalNode)
		res, err := getEPlanExecutorNumber(nodea.Input, pn)
		if err != nil {
			return -1, err
		}
		return res + 1, nil

	case *PlanOrderByNode:
		nodea := node.(*PlanOrderByNode)
		res, err := getEPlanExecutorNumber(nodea.Input, pn)
		if err != nil {
			return -1, err
		}
		return res + pn + 1, nil

	case *PlanUnionNode:
		nodea := node.(*PlanUnionNode)
		leftRes, err := getEPlanExecutorNumber(nodea.LeftInput, pn)
		if err != nil {
			return -1, err
		}
		rightRes, err := getEPlanExecutorNumber(nodea.RightInput, pn)
		if err != nil {
			return -1, err
		}
		return leftRes + rightRes + pn, nil

	case *PlanFilterNode:
		nodea := node.(*PlanFilterNode)
		res, err := getEPlanExecutorNumber(nodea.Input, pn)
		if err != nil {
			return -1, err
		}
		return res + pn, nil

	default:
		Logger.Errorf("getEPlanExecutorNumber: unknown node type")
		return -1, fmt.Errorf("getEPlanExecutorNumber: unknown node type")
	}
}
