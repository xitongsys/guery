package EPlan

import (
	"fmt"

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

func NewStack(items *[]pb.Location) *Stack {
	return &Stack{
		Items: items,
	}
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
	Logger.Infof("======inputs=%v", inputs)
	aggNode := NewEPlanAggregateNode(inputs, output)
	*ePlanNodes = append(*ePlanNodes, aggNode)
	return aggNode, err
}

func createEPlan(node PlanNode, ePlanNodes *[]ENode, freeExecutors *Stack, pn int) ([]ENode, error) {
	res := []ENode{}
	switch node.(type) {
	case *PlanScanNode:
		nodea := node.(*PlanScanNode)
		output, err := freeExecutors.Pop()
		if err != nil {
			return res, err
		}
		outputs := []pb.Location{}
		for i := 0; i < pn; i++ {
			output.ChannelIndex = int32(i)
			outputs = append(outputs, output)
		}
		res = append(res, NewEPlanScanNode(nodea, output, outputs))
		*ePlanNodes = append(*ePlanNodes, res...)
		return res, nil

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
		*ePlanNodes = append(*ePlanNodes, res...)
		return res, nil

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
		duplicateNode := NewEPlanDuplicateNode(inputs, outputs)

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

	case *PlanFiliterNode:
		nodea := node.(*PlanFiliterNode)
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
				res = append(res, NewEPlanFiliterNode(nodea, input, output))
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
