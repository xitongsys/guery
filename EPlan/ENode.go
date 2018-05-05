package EPlan

import (
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
	EDUPLICATE
)

type ENode interface {
	GetNodeType() EPlanNodeType
	GetOutputs() []pb.Location
}

func CreateEPlan(node PlanNode, ePlanNodes []ENode, freeExecutors []pb.Location, pn int) []ENode {
	res := []ENode{}
	switch node.(type) {
	case *PlanScanNode:
		nodea := node.(*PlanScanNode)
		ln := len(freeExecutors)
		outputs := make([]pb.Location, pn)
		for i := 0; i < pn; i++ {
			outputs[i] = freeExecutors[ln-1]
			outputs[i].ChannelIndex = int32(i)
		}
		freeExecutors = freeExecutors[:ln-1]
		res = append(res, NewEPlanScanNode(nodea, outputs))
		ePlanNodes = append(ePlanNodes, res...)
		return res

	case *PlanSelectNode:
		nodea := node.(*PlanSelectNode)
		inputNodes := CreateEPlan(nodea.Input, ePlanNodes, freeExecutors, pn)
		for _, inputNode := range inputNodes {
			for _, input := range inputNode.GetOutputs() {
				ln := len(freeExecutors)
				output := freeExecutors[ln-1]
				output.ChannelIndex = 0
				freeExecutors = freeExecutors[:ln-1]
				res = append(res, NewEPlanSelectNode(nodea, input, output))
			}
		}
		ePlanNodes = append(ePlanNodes, res...)
		return res

	case *PlanGroupByNode:
		nodea := node.(*PlanGroupByNode)
		inputNodes := CreateEPlan(nodea.Input, ePlanNodes, freeExecutors, pn)
		inputs, outputs := []pb.Location{}, []pb.Location{}
		for _, in := range inputNodes {
			inputs = append(inputs, in.GetOutputs()...)
		}

		ln := len(freeExecutors)
		for i := 0; i < pn; i++ {
			output := freeExecutors[ln-1]
			output.ChannelIndex = int32(i)
			outputs = append(outputs, output)
		}
		freeExecutors = freeExecutors[:ln-1]

		res = append(res, NewEPlanGroupByNode(nodea, inputs, outputs))
		ePlanNodes = append(ePlanNodes, res...)
		return res

	case *PlanJoinNode:
		nodea := node.(*PlanJoinNode)
		leftInputNodes := CreateEPlan(nodea.LeftInput, ePlanNodes, freeExecutors, pn)
		rightInputNodes := CreateEPlan(nodea.RightInput, ePlanNodes, freeExecutors, pn)

		inputs, outputs := []pb.Location{}, []pb.Location{}
		for _, inputNode := range rightInputNodes {
			inputs = append(inputs, inputNode.GetOutputs()...)
		}
		ln := len(freeExecutors)
		for i := 0; i < pn; i++ {
			output := freeExecutors[ln-1]
			output.ChannelIndex = int32(i)
			outputs = append(outputs, output)
		}
		freeExecutors = freeExecutors[:ln-1]
		duplicateNode := NewEPlanDuplicateNode(inputs, outputs)

		rightInputs := duplicateNode.GetOutputs()
		leftInputs := []pb.Location{}
		for _, leftInputNode := range leftInputNodes {
			leftInputs = append(leftInputs, leftInputNode.GetOutputs()...)
		}
		if len(leftInputs) != len(rightInputs) {
			Logger.Errorf("JoinNode leftInputs number <> rightInputs number")
			return nil
		}

		for i := 0; i < len(leftInputs); i++ {
			ln := len(freeExecutors)
			output := freeExecutors[ln-1]
			output.ChannelIndex = 0
			joinNode := NewEPlanJoinNode(nodea, leftInputs[i], rightInputs[i], output)
			res = append(res, joinNode)
			freeExecutors = freeExecutors[:ln-1]
		}
		ePlanNodes = append(ePlanNodes, res...)
		return res

	default:
		Logger.Errorf("Unknown node type")
		return nil

	}
}
