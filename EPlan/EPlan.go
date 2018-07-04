package EPlan

import (
	"fmt"
	"math/rand"

	"github.com/xitongsys/guery/FileSystem"
	"github.com/xitongsys/guery/FileSystem/Partition"
	"github.com/xitongsys/guery/Logger"
	. "github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/pb"
)

/////////////////////////////////////////
type Stack struct {
	Items *[]pb.Location
}

func (self *Stack) Pop() (pb.Location, error) {
	res := pb.Location{}
	ln := len(*self.Items)
	if ln <= 0 {
		return res, fmt.Errorf("no item available")
	}
	res = (*self.Items)[ln-1]
	*self.Items = (*self.Items)[:ln-1]
	return res, nil
}

func (self *Stack) Random() {
	ln := len(*self.Items)
	for i := 0; i < ln-1; i++ {
		r := rand.Int() % (ln - i - 1)
		ri := i + r + 1
		(*self.Items)[i], (*self.Items)[ri] = (*self.Items)[ri], (*self.Items)[i]
	}
}

func NewStack(items *[]pb.Location) *Stack {
	res := &Stack{
		Items: items,
	}
	res.Random()
	return res
}

/////////////////////////////////////////
func CreateEPlan(node PlanNode, ePlanNodes *[]ENode, freeExecutors *[]pb.Location, pn int) (ENode, error) {
	exeStack := NewStack(freeExecutors)
	inputNodes, err := createEPlan(node, ePlanNodes, exeStack, pn)
	if err != nil {
		return nil, err
	}
	ln := len(*freeExecutors)
	if ln <= 0 {
		return nil, fmt.Errorf("no executor available")
	}
	output, err := exeStack.Pop()
	if err != nil {
		return nil, err
	}
	inputs := []pb.Location{}
	for _, inputNode := range inputNodes {
		inputs = append(inputs, inputNode.GetOutputs()...)
	}
	aggNode := NewEPlanAggregateNode(inputs, output)
	*ePlanNodes = append(*ePlanNodes, aggNode)
	return aggNode, err
}

func createEPlan(node PlanNode, ePlanNodes *[]ENode, freeExecutors *Stack, pn int) ([]ENode, error) {
	res := []ENode{}
	switch node.(type) {

	case *PlanShowNode:
		nodea := node.(*PlanShowNode)
		output, err := freeExecutors.Pop()
		if err != nil {
			return res, err
		}
		output.ChannelIndex = int32(0)
		res = append(res, NewEPlanShowNode(nodea, output))
		*ePlanNodes = append(*ePlanNodes, res...)
		return res, nil

	case *PlanScanNode:
		nodea := node.(*PlanScanNode)
		outputs := []pb.Location{}
		for i := 0; i < pn; i++ {
			output, err := freeExecutors.Pop()
			if err != nil {
				return res, err
			}
			output.ChannelIndex = int32(0)
			outputs = append(outputs, output)
		}

		parInfos := make([]*Partition.PartitionInfo, pn)
		recMap := make([]map[int]int, pn)
		for i := 0; i < pn; i++ {
			parInfos[i] = Partition.NewPartitionInfo(nodea.PartitionInfo.Metadata)
			recMap[i] = map[int]int{}
		}

		k := 0
		if nodea.PartitionInfo.IsPartition() {
			partitionNum := nodea.PartitionInfo.GetPartitionNum()
			parFilters := []*BooleanExpressionNode{}
			for _, f := range nodea.Filters {
				cols, err := f.GetColumns()
				if err != nil {
					return res, err
				}
				if nodea.PartitionInfo.Metadata.Contains(cols) {
					parFilters = append(parFilters, f)
				}
			}

			for i := 0; i < partitionNum; i++ {
				prg := nodea.PartitionInfo.GetPartitionRowGroup(i)
				flag := true
				for _, exp := range parFilters {
					if r, err := exp.Result(prg); err != nil {
						return res, err
					} else if !r.(bool) {
						flag = false
						break
					}
				}
				if !flag {
					continue
				}

				row := prg.Rows[0]
				location := nodea.PartitionInfo.GetLocation(i)
				fileType := nodea.PartitionInfo.GetFileType(i)
				files := nodea.PartitionInfo.GetPartitionFiles(i)
				for _, file := range files {
					if _, ok := recMap[k][i]; !ok {
						recMap[k][i] = parInfos[k].GetPartitionNum()
						parInfos[k].Write(row)
						parInfos[k].Locations = append(parInfos[k].Locations, location)
						parInfos[k].FileTypes = append(parInfos[k].FileTypes, fileType)
						parInfos[k].FileLists = append(parInfos[k].FileLists, []*FileSystem.FileLocation{})
					}
					j := recMap[k][i]
					parInfos[k].FileLists[j] = append(parInfos[k].FileLists[j], file)

					k++
					k = k % pn
				}
			}

		} else {
			for i, file := range nodea.PartitionInfo.GetNoPartititonFiles() {
				parInfos[i%pn].FileList = append(parInfos[i%pn].FileList, file)
			}
		}

		resScan := []ENode{}
		for i := 0; i < pn; i++ {
			resScan = append(resScan, NewEPlanScanNode(nodea, parInfos[i], outputs[i], []pb.Location{outputs[i]}))
		}

		/*

			balanceLoc, err := freeExecutors.Pop()
			if err != nil {
				return res, err
			}
			balanceLoc.ChannelIndex = 0
			balanceOutputs, balanceInputs := make([]pb.Location, pn), make([]pb.Location, pn)

			for i := 0; i < pn; i++ {
				balanceInputs[i] = resScan[i].GetOutputs()[0]
				balanceOutputs[i] = balanceLoc
				balanceOutputs[i].ChannelIndex = int32(i)
			}
			res = append(res, NewEPlanBalanceNode(balanceInputs, balanceOutputs))

			*ePlanNodes = append(*ePlanNodes, resScan...)
			*ePlanNodes = append(*ePlanNodes, res...)

			return res, nil
		*/

		*ePlanNodes = append(*ePlanNodes, resScan...)
		return resScan, nil

	case *PlanSelectNode:
		nodea := node.(*PlanSelectNode)
		inputNodes, err := createEPlan(nodea.Input, ePlanNodes, freeExecutors, pn)
		if err != nil {
			return res, err
		}
		for _, inputNode := range inputNodes {
			for _, input := range inputNode.GetOutputs() {
				output, err := freeExecutors.Pop()
				if err != nil {
					return res, err
				}
				output.ChannelIndex = 0
				res = append(res, NewEPlanSelectNode(nodea, input, output))
			}
		}
		*ePlanNodes = append(*ePlanNodes, res...)
		return res, nil

	case *PlanGroupByNode:
		nodea := node.(*PlanGroupByNode)
		inputNodes, err := createEPlan(nodea.Input, ePlanNodes, freeExecutors, pn)
		if err != nil {
			return nil, err
		}
		inputs, outputs := []pb.Location{}, []pb.Location{}
		for _, in := range inputNodes {
			inputs = append(inputs, in.GetOutputs()...)
		}
		output, err := freeExecutors.Pop()
		if err != nil {
			return res, err
		}
		for i := 0; i < pn; i++ {
			output.ChannelIndex = int32(i)
			outputs = append(outputs, output)
		}
		res = append(res, NewEPlanGroupByNode(nodea, inputs, outputs))

		for i := 0; i < pn; i++ {
			loc, err := freeExecutors.Pop()
			if err != nil {
				return res, err
			}
			loc.ChannelIndex = 0
			res = append(res, NewEPlanGroupByLocalNode(nodea, outputs[i], loc))
		}

		*ePlanNodes = append(*ePlanNodes, res...)
		return res[1:], nil

	case *PlanJoinNode:
		nodea := node.(*PlanJoinNode)
		leftInputNodes, err1 := createEPlan(nodea.LeftInput, ePlanNodes, freeExecutors, pn)
		if err1 != nil {
			return nil, err1
		}
		rightInputNodes, err2 := createEPlan(nodea.RightInput, ePlanNodes, freeExecutors, pn)
		if err2 != nil {
			return nil, err2
		}

		//duplicate right inputs
		inputs, outputs := []pb.Location{}, []pb.Location{}
		for _, inputNode := range rightInputNodes {
			inputs = append(inputs, inputNode.GetOutputs()...)
		}
		output, err := freeExecutors.Pop()
		if err != nil {
			return res, err
		}
		for i := 0; i < pn; i++ {
			output.ChannelIndex = int32(i)
			outputs = append(outputs, output)
		}
		duplicateNode := NewEPlanDuplicateNode(inputs, outputs, nil)

		//join
		rightInputs := duplicateNode.GetOutputs()
		leftInputs := []pb.Location{}
		for _, leftInputNode := range leftInputNodes {
			leftInputs = append(leftInputs, leftInputNode.GetOutputs()...)
		}
		if len(leftInputs) != len(rightInputs) {
			return nil, fmt.Errorf("JoinNode leftInputs number <> rightInputs number")
		}

		for i := 0; i < len(leftInputs); i++ {
			output, err := freeExecutors.Pop()
			if err != nil {
				return res, err
			}
			output.ChannelIndex = 0
			joinNode := NewEPlanJoinNode(nodea, leftInputs[i], rightInputs[i], output)
			res = append(res, joinNode)
		}
		*ePlanNodes = append(*ePlanNodes, duplicateNode)
		*ePlanNodes = append(*ePlanNodes, res...)
		return res, nil

	case *PlanHashJoinNode:
		nodea := node.(*PlanHashJoinNode)
		leftInputNodes, err1 := createEPlan(nodea.LeftInput, ePlanNodes, freeExecutors, pn)
		if err1 != nil {
			return nil, err1
		}
		rightInputNodes, err2 := createEPlan(nodea.RightInput, ePlanNodes, freeExecutors, pn)
		if err2 != nil {
			return nil, err2
		}

		//duplicate right inputs
		inputs, outputs := []pb.Location{}, []pb.Location{}
		for _, inputNode := range rightInputNodes {
			inputs = append(inputs, inputNode.GetOutputs()...)
		}
		output, err := freeExecutors.Pop()
		if err != nil {
			return res, err
		}
		for i := 0; i < pn; i++ {
			output.ChannelIndex = int32(i)
			outputs = append(outputs, output)
		}
		duplicateNode := NewEPlanDuplicateNode(inputs, outputs, nodea.RightKeys)

		//hash join
		rightInputs := duplicateNode.GetOutputs()
		leftInputs := []pb.Location{}
		for _, leftInputNode := range leftInputNodes {
			leftInputs = append(leftInputs, leftInputNode.GetOutputs()...)
		}
		if len(leftInputs) != len(rightInputs) {
			return nil, fmt.Errorf("JoinNode leftInputs number <> rightInputs number")
		}

		for i := 0; i < len(leftInputs); i++ {
			output, err := freeExecutors.Pop()
			if err != nil {
				return res, err
			}
			output.ChannelIndex = 0
			joinNode := NewEPlanHashJoinNode(nodea, leftInputs[i], rightInputs[i], output)
			res = append(res, joinNode)
		}
		*ePlanNodes = append(*ePlanNodes, duplicateNode)
		*ePlanNodes = append(*ePlanNodes, res...)
		return res, nil

	case *PlanLimitNode:
		nodea := node.(*PlanLimitNode)
		inputNodes, err := createEPlan(nodea.Input, ePlanNodes, freeExecutors, pn)
		if err != nil {
			return res, err
		}
		for _, inputNode := range inputNodes {
			for _, input := range inputNode.GetOutputs() {
				output, err := freeExecutors.Pop()
				if err != nil {
					return res, err
				}
				output.ChannelIndex = 0
				res = append(res, NewEPlanLimitNode(nodea, input, output))
			}
		}
		*ePlanNodes = append(*ePlanNodes, res...)
		return res, nil

	case *PlanAggregateNode:
		nodea := node.(*PlanAggregateNode)
		inputNodes, err := createEPlan(nodea.Input, ePlanNodes, freeExecutors, pn)
		if err != nil {
			return res, err
		}
		output, err := freeExecutors.Pop()
		if err != nil {
			return res, err
		}
		inputs := []pb.Location{}
		for _, inputNode := range inputNodes {
			inputs = append(inputs, inputNode.GetOutputs()...)
		}
		res = append(res, NewEPlanAggregateNode(inputs, output))
		*ePlanNodes = append(*ePlanNodes, res...)
		return res, nil

	case *PlanFilterNode:
		nodea := node.(*PlanFilterNode)
		inputNodes, err := createEPlan(nodea.Input, ePlanNodes, freeExecutors, pn)
		if err != nil {
			return res, err
		}
		for _, inputNode := range inputNodes {
			for _, input := range inputNode.GetOutputs() {
				output, err := freeExecutors.Pop()
				if err != nil {
					return res, err
				}
				output.ChannelIndex = 0
				res = append(res, NewEPlanFilterNode(nodea, input, output))
			}
		}
		*ePlanNodes = append(*ePlanNodes, res...)
		return res, nil

	case *PlanUnionNode:
		nodea := node.(*PlanUnionNode)
		leftInputNodes, err1 := createEPlan(nodea.LeftInput, ePlanNodes, freeExecutors, pn)
		if err1 != nil {
			return nil, err1
		}
		rightInputNodes, err2 := createEPlan(nodea.RightInput, ePlanNodes, freeExecutors, pn)
		if err2 != nil {
			return nil, err2
		}

		//union
		leftInputs := []pb.Location{}
		for _, leftInputNode := range leftInputNodes {
			leftInputs = append(leftInputs, leftInputNode.GetOutputs()...)
		}
		rightInputs := []pb.Location{}
		for _, rightInputNode := range rightInputNodes {
			rightInputs = append(rightInputs, rightInputNode.GetOutputs()...)
		}

		if len(leftInputs) != len(rightInputs) {
			return nil, fmt.Errorf("JoinNode leftInputs number <> rightInputs number")
		}

		for i := 0; i < len(leftInputs); i++ {
			output, err := freeExecutors.Pop()
			if err != nil {
				return res, err
			}
			output.ChannelIndex = 0
			joinNode := NewEPlanUnionNode(nodea, leftInputs[i], rightInputs[i], output)
			res = append(res, joinNode)
		}
		*ePlanNodes = append(*ePlanNodes, res...)
		return res, nil

	case *PlanOrderByNode:
		nodea := node.(*PlanOrderByNode)
		inputNodes, err := createEPlan(nodea.Input, ePlanNodes, freeExecutors, pn)
		if err != nil {
			return nil, err
		}

		inputs := []pb.Location{}
		for _, inputNode := range inputNodes {
			inputs = append(inputs, inputNode.GetOutputs()...)
		}

		localRes := []ENode{}
		for _, input := range inputs {
			output, err := freeExecutors.Pop()
			if err != nil {
				return res, err
			}
			output.ChannelIndex = 0
			orderByNodeLocal := NewEPlanOrderByLocalNode(nodea, input, output)
			localRes = append(localRes, orderByNodeLocal)
		}

		inputs = []pb.Location{}
		for _, inputNode := range localRes {
			inputs = append(inputs, inputNode.GetOutputs()...)
		}
		output, err := freeExecutors.Pop()
		if err != nil {
			return res, err
		}
		orderByNode := NewEPlanOrderByNode(nodea, inputs, output)
		res = append(res, orderByNode)

		*ePlanNodes = append(*ePlanNodes, localRes...)
		*ePlanNodes = append(*ePlanNodes, orderByNode)
		return res, nil

	default:
		Logger.Errorf("createEPlan: unknown node type")
		return nil, fmt.Errorf("createEPlan: unknown node type")

	}
}
