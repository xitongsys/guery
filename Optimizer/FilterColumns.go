package Optimizer

import (
	"fmt"
	"sort"

	"github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/Util"
)

func FilterColumns(node Plan.PlanNode, columns []string) error {
	if node == nil {
		return nil
	}
	switch node.(type) {
	case *Plan.PlanJoinNode, *Plan.PlanLimitNode, *Plan.PlanUnionNode, *Plan.PlanCombineNode:
		indexes := []int{}
		md := node.GetMetadata()
		//for join node
		if nodea, ok := node.(*Plan.PlanJoinNode); ok {
			cs, err := nodea.JoinCriteria.GetColumns()
			if err != nil {
				return err
			}
			columns = append(columns, cs...)
		}
		for _, c := range columns {
			if index, err := md.GetIndexByName(c); err == nil {
				indexes = append(indexes, index)
			}
		}
		sort.Ints(indexes)

		inputs := node.GetInputs()
		mdis := []*Util.Metadata{}
		for _, input := range inputs {
			mdis = append(mdis, input.GetMetadata())
		}

		columnsForInput := make([][]string, len(inputs))

		i, indexNum := 0, mdis[0].GetColumnNumber()
		for _, index := range indexes {
			if index < indexNum {
				indexForInput := index - (indexNum - mdis[i].GetColumnNumber())
				cname := mdis[i].Columns[indexForInput].GetName()
				columnsForInput[i] = append(columnsForInput[i], cname)
			} else {
				i++
				indexNum += mdis[i].GetColumnNumber()
			}
		}
		for i, input := range inputs {
			err := FilterColumns(input, columnsForInput[i])
			if err != nil {
				return err
			}
		}

	case *Plan.PlanFiliterNode:
		nodea := node.(*Plan.PlanFiliterNode)
		columnsForInput, err := nodea.BooleanExpression.GetColumns()
		if err != nil {
			return err
		}
		return FilterColumns(nodea.Input, columnsForInput)

	case *Plan.PlanGroupByNode:
		nodea := node.(*Plan.PlanGroupByNode)
		columnsForInput := []string{}
		cs, err := nodea.GroupBy.GetColumns()
		if err != nil {
			return err
		}
		columnsForInput = append(columnsForInput, cs...)
		return FilterColumns(nodea.Input, columnsForInput)

	case *Plan.PlanOrderByNode:
		nodea := node.(*Plan.PlanOrderByNode)
		columnsForInput := []string{}
		for _, item := range nodea.SortItems {
			cs, err := item.GetColumns()
			if err != nil {
				return err
			}
			columnsForInput = append(columnsForInput, cs...)
		}
		return FilterColumns(nodea.Input, columnsForInput)

	case *Plan.PlanSelectNode:
		nodea := node.(*Plan.PlanSelectNode)
		columnsForInput := []string{}
		for _, item := range nodea.SelectItems {
			cs, err := item.GetColumns(nodea.Input.GetMetadata())
			if err != nil {
				return err
			}
			columnsForInput = append(columnsForInput, cs...)
		}
		return FilterColumns(nodea.Input, columnsForInput)

	case *Plan.PlanScanNode:
		nodea := node.(*Plan.PlanScanNode)
		nodea.Metadata = nodea.Metadata.SelectColumns(columns)
		parent := nodea.GetOutput()
		for parent != nil {
			parent.SetMetadata()
			parent = parent.GetOutput()
		}
		return nil

	case *Plan.PlanRenameNode: //already use deleteRenameNode
		return nil
	default:
		return fmt.Errorf("Unknown PlanNode type")
	}

	return nil
}
