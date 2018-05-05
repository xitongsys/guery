package EPlan

import (
	"fmt"

	"github.com/xitongsys/guery/Logger"
	. "github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/pb"
)

type EPlanNodeType int32

const (
	_ EPlanNodeType = iota
	ESCANNODE
	ESELECTNODE
	EGROUPBYNODE
	EFILITERNODE
	EUNIONNODE
	ELIMITNODE
	EORDERBYNODE
	EJOINNODE
	EHAVINGNODE
	ECOMBINENODE
	EDUPLICATENODE
)

type ENode interface {
	GetNodeType() EPlanNodeType
	GetOutputs() []pb.Location
	GetLocation() pb.Location
}

func CreateEPlan(node PlanNode, ePlanNodes *[]ENode, freeExecutors *[]pb.Location, pn int) ([]ENode, error) {
	res := []ENode{}
	switch node.(type) {
	case *PlanScanNode:
		nodea := node.(*PlanScanNode)
		ln := len(*freeExecutors)
		if ln <= 0 {
			return nil, fmt.Errorf("No executor available")
		}
		outputs := make([]pb.Location, pn)
		for i := 0; i < pn; i++ {
			outputs[i] = (*freeExecutors)[ln-1]
			outputs[i].ChannelIndex = int32(i)
		}
		*freeExecutors = (*freeExecutors)[:ln-1]
		res = append(res, NewEPlanScanNode(nodea, outputs))
		*ePlanNodes = append(*ePlanNodes, res...)
		return res, nil

	case *PlanSelectNode:
		nodea := node.(*PlanSelectNode)
		inputNodes, err := CreateEPlan(nodea.Input, ePlanNodes, freeExecutors, pn)
		if err != nil {
			return res, err
		}
		for _, inputNode := range inputNodes {
			for _, input := range inputNode.GetOutputs() {
				ln := len(*freeExecutors)
				if ln <= 0 {
					return nil, fmt.Errorf("No executor available")
				}
				output := (*freeExecutors)[ln-1]
				output.ChannelIndex = 0
				*freeExecutors = (*freeExecutors)[:ln-1]
				res = append(res, NewEPlanSelectNode(nodea, input, output))
			}
		}
		*ePlanNodes = append(*ePlanNodes, res...)
		return res, nil

	case *PlanGroupByNode:
		nodea := node.(*PlanGroupByNode)
		inputNodes, err := CreateEPlan(nodea.Input, ePlanNodes, freeExecutors, pn)
		if err != nil {
			return nil, err
		}
		inputs, outputs := []pb.Location{}, []pb.Location{}
		for _, in := range inputNodes {
			inputs = append(inputs, in.GetOutputs()...)
		}

		ln := len(*freeExecutors)
		if ln <= 0 {
			return nil, fmt.Errorf("No executor available")
		}
		for i := 0; i < pn; i++ {
			output := (*freeExecutors)[ln-1]
			output.ChannelIndex = int32(i)
			outputs = append(outputs, output)
		}
		*freeExecutors = (*freeExecutors)[:ln-1]

		res = append(res, NewEPlanGroupByNode(nodea, inputs, outputs))
		*ePlanNodes = append(*ePlanNodes, res...)
		return res, nil

	case *PlanJoinNode:
		nodea := node.(*PlanJoinNode)
		leftInputNodes, err1 := CreateEPlan(nodea.LeftInput, ePlanNodes, freeExecutors, pn)
		if err1 != nil {
			return nil, err1
		}
		rightInputNodes, err2 := CreateEPlan(nodea.RightInput, ePlanNodes, freeExecutors, pn)
		if err2 != nil {
			return nil, err2
		}

		inputs, outputs := []pb.Location{}, []pb.Location{}
		for _, inputNode := range rightInputNodes {
			inputs = append(inputs, inputNode.GetOutputs()...)
		}
		ln := len(*freeExecutors)
		if ln <= 0 {
			return nil, fmt.Errorf("No executor available")
		}
		for i := 0; i < pn; i++ {
			output := (*freeExecutors)[ln-1]
			output.ChannelIndex = int32(i)
			outputs = append(outputs, output)
		}
		*freeExecutors = (*freeExecutors)[:ln-1]
		duplicateNode := NewEPlanDuplicateNode(inputs, outputs)

		rightInputs := duplicateNode.GetOutputs()
		leftInputs := []pb.Location{}
		for _, leftInputNode := range leftInputNodes {
			leftInputs = append(leftInputs, leftInputNode.GetOutputs()...)
		}
		if len(leftInputs) != len(rightInputs) {
			return nil, fmt.Errorf("JoinNode leftInputs number <> rightInputs number")
		}

		for i := 0; i < len(leftInputs); i++ {
			ln := len(*freeExecutors)
			if ln <= 0 {
				return nil, fmt.Errorf("No executor available")
			}
			output := (*freeExecutors)[ln-1]
			output.ChannelIndex = 0
			joinNode := NewEPlanJoinNode(nodea, leftInputs[i], rightInputs[i], output)
			res = append(res, joinNode)
			*freeExecutors = (*freeExecutors)[:ln-1]
		}
		*ePlanNodes = append(*ePlanNodes, res...)
		return res, nil

	default:
		Logger.Errorf("Unknown node type")
		return nil, fmt.Errorf("Unknown node type")

	}
}
