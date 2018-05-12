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
	case *PlanScanNode:
		return pn, nil

	case *PlanSelectNode:
		nodea := node.(*PlanSelectNode)
		res, err := getEPlanExecutorNumber(nodea.Input, pn)
		if err != nil {
			return -1, err
		}
		res += pn
		return res, nil

	case *PlanGroupByNode:
		nodea := node.(*PlanGroupByNode)
		res, err := getEPlanExecutorNumber(nodea.Input, pn)
		if err != nil {
			return -1, err
		}
		return res + 1, nil

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

	default:
		Logger.Errorf("Unknown node type")
		return -1, fmt.Errorf("Unknown node type")
	}
}
