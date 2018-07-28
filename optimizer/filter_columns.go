package optimizer

import (
	"fmt"
	"sort"

	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/plan"
)

func FilterColumns(node Plan.PlanNode, columns []string) error {
	if node == nil {
		return nil
	}
	switch node.(type) {
	case *Plan.PlanJoinNode, *Plan.PlanLimitNode, *Plan.PlanUnionNode, *Plan.PlanCombineNode, *Plan.PlanAggregateNode, *Plan.PlanAggregateFuncLocalNode:
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
		mdis := []*Metadata.Metadata{}
		for _, input := range inputs {
			mdis = append(mdis, input.GetMetadata())
		}

		columnsForInput := make([][]string, len(inputs))

		i, indexNum := 0, mdis[0].GetColumnNumber()
		for j := 0; j < len(indexes); j++ {
			index := indexes[j]
			if index < indexNum {
				indexForInput := index - (indexNum - mdis[i].GetColumnNumber())
				cname := mdis[i].Columns[indexForInput].GetName()
				columnsForInput[i] = append(columnsForInput[i], cname)
			} else {
				i++
				indexNum += mdis[i].GetColumnNumber()
				j--
			}
		}
		for i, input := range inputs {
			err := FilterColumns(input, columnsForInput[i])
			if err != nil {
				return err
			}
		}

	case *Plan.PlanFilterNode:
		nodea := node.(*Plan.PlanFilterNode)
		columnsForInput := []string{}
		for _, be := range nodea.BooleanExpressions {
			cols, err := be.GetColumns()
			if err != nil {
				return err
			}
			columnsForInput = append(columnsForInput, cols...)
		}
		columnsForInput = append(columnsForInput, columns...)
		return FilterColumns(nodea.Input, columnsForInput)

	case *Plan.PlanGroupByNode:
		nodea := node.(*Plan.PlanGroupByNode)
		columnsForInput, err := nodea.GroupBy.GetColumns()
		if err != nil {
			return err
		}
		columnsForInput = append(columnsForInput, columns...)
		return FilterColumns(nodea.Input, columnsForInput)

	case *Plan.PlanOrderByNode:
		nodea := node.(*Plan.PlanOrderByNode)
		columnsForInput := columns
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
		columnsForInput := columns
		for _, item := range nodea.SelectItems {
			cs, err := item.GetColumns(nodea.Input.GetMetadata())
			if err != nil {
				return err
			}
			columnsForInput = append(columnsForInput, cs...)
		}
		if nodea.Having != nil {
			cs, err := nodea.Having.GetColumns()
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

	case *Plan.PlanShowNode:
		return nil

	case *Plan.PlanRenameNode: //already use deleteRenameNode
		return nil

	default:
		return fmt.Errorf("[Optimizer:FilterColumns]Unknown PlanNode type")
	}

	return nil
}
