package optimizer

import (
	"fmt"
	"sort"

	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/plan"
)

func FilterColumns(node plan.PlanNode, columns []string) error {
	if node == nil {
		return nil
	}
	switch node.(type) {
	case *plan.PlanJoinNode, *plan.PlanLimitNode, *plan.PlanUnionNode, *plan.PlanCombineNode, *plan.PlanAggregateNode, *plan.PlanAggregateFuncLocalNode:
		indexes := []int{}
		md := node.GetMetadata()
		//for join node
		if nodea, ok := node.(*plan.PlanJoinNode); ok {
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
		mdis := []*metadata.Metadata{}
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

	case *plan.PlanFilterNode:
		nodea := node.(*plan.PlanFilterNode)
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

	case *plan.PlanGroupByNode:
		nodea := node.(*plan.PlanGroupByNode)
		columnsForInput, err := nodea.GroupBy.GetColumns()
		if err != nil {
			return err
		}
		columnsForInput = append(columnsForInput, columns...)
		return FilterColumns(nodea.Input, columnsForInput)

	case *plan.PlanOrderByNode:
		nodea := node.(*plan.PlanOrderByNode)
		columnsForInput := columns
		for _, item := range nodea.SortItems {
			cs, err := item.GetColumns()
			if err != nil {
				return err
			}
			columnsForInput = append(columnsForInput, cs...)
		}
		return FilterColumns(nodea.Input, columnsForInput)

	case *plan.PlanSelectNode:
		nodea := node.(*plan.PlanSelectNode)
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

	case *plan.PlanScanNode:
		nodea := node.(*plan.PlanScanNode)
		nodea.Metadata = nodea.Metadata.SelectColumns(columns)
		parent := nodea.GetOutput()

		for parent != nil {
			parent.SetMetadata()
			parent = parent.GetOutput()
		}
		return nil

	case *plan.PlanShowNode:
		return nil

	case *plan.PlanRenameNode: //already use deleteRenameNode
		return nil

	default:
		return fmt.Errorf("[Optimizer:FilterColumns]Unknown PlanNode type")
	}

	return nil
}
