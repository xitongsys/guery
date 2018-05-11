package Optimizer

import (
	"fmt"

	"github.com/xitongsys/guery/Plan"
)

func DeleteRenameNode(node Plan.PlanNode) error {
	switch node.(type) {
	case *PlanScanNode:
		return nil

	case *PlanSelectNode:
		nodea := node.(*PlanSelectNode)
		return DeleteRenameNode(nodea.Input)

	case *PlanGroupByNode:
		nodea := node.(*PlanGroupByNode)
		return DeleteRenameNode(nodea.Input)

	case *PlanJoinNode:
		nodea := node.(*PlanJoinNode)
		if err := DeleteRenameNode(nodea.LeftInput); err != nil {
			return err
		}
		if err := DeleteRenameNode(nodea.RightInput); err != nil {
			return err
		}
		return nil

	case *PlanRenameNode:
		nodea := node.(*PlanRenameNode)
		if err := DeleteRenameNode(nodea.Input); err != nil {
			return err
		}
		parent := nodea.Output
		md := nodea.Input.GetMetadata()
		md.Rename(nodea.Rename)
		nodea.Input.SetOutput(parent)

	default:
		return fmt.Errorf("unknown PlanNode type")
	}
}
