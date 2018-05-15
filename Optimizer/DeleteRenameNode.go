package Optimizer

import (
	"fmt"

	"github.com/xitongsys/guery/Plan"
)

func DeleteRenameNode(node Plan.PlanNode) error {
	switch node.(type) {
	case *Plan.PlanScanNode:
		return nil

	case *Plan.PlanSelectNode:
		nodea := node.(*Plan.PlanSelectNode)
		return DeleteRenameNode(nodea.Input)

	case *Plan.PlanGroupByNode:
		nodea := node.(*Plan.PlanGroupByNode)
		return DeleteRenameNode(nodea.Input)

	case *Plan.PlanJoinNode:
		nodea := node.(*Plan.PlanJoinNode)
		if err := DeleteRenameNode(nodea.LeftInput); err != nil {
			return err
		}
		if err := DeleteRenameNode(nodea.RightInput); err != nil {
			return err
		}
		return nil

	case *Plan.PlanUnionNode:
		nodea := node.(*Plan.PlanUnionNode)
		if err := DeleteRenameNode(nodea.LeftInput); err != nil {
			return err
		}
		if err := DeleteRenameNode(nodea.RightInput); err != nil {
			return err
		}

	case *Plan.PlanOrderByNode:
		nodea := node.(*Plan.PlanOrderByNode)
		return DeleteRenameNode(nodea.Input)

	case *Plan.PlanLimitNode:
		nodea := node.(*Plan.PlanLimitNode)
		return DeleteRenameNode(nodea.Input)

	case *Plan.PlanFiliterNode:
		nodea := node.(*Plan.PlanFiliterNode)
		return DeleteRenameNode(nodea.Input)

	case *Plan.PlanRenameNode:
		nodea := node.(*Plan.PlanRenameNode)
		if err := DeleteRenameNode(nodea.Input); err != nil {
			return err
		}
		parent := nodea.Output
		md := nodea.Input.GetMetadata()
		md.Rename(nodea.Rename)
		nodea.Input.SetOutput(parent)
		return nil

	default:
		return fmt.Errorf("unknown PlanNode type")
	}
	return nil
}
