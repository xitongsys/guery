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

func CreateEPlan(node PlanNode, ePlanNodes []ENode, freeExecutors []pb.Loacation, pn int) []ENode {
	res := []ENode{}
	switch node.(type) {
	case *PlanScanNode:
		nodea := node.(*PlanScanNode)
		ln := len(freeExecutors)
		outputs := make([]Loacation, pn)
		for i := 0; i < pn; i++ {
			outputs[i] = freeExecutors[ln-1]
			outputs[i].ChannelIndex = i
		}
		freeExecutors = freeExecutors[:ln-1]
		res = append(res, NewEPlanScanNode(nodea, outputs))
		ePlanNodes = append(ePlanNodes, res)
		return res

	case *PlanSelectNode:
		nodea := node.(*PlanSelectNode)
		inputNodes := CreateEPlan(nodea.Input, ePlanNodes, freeExecutors, pn)
		for _, in := range inputNodes {
			for _, input := range n.GetOutputs() {
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
			output.ChannelIndex = i
			outputs = append(output)
		}
		freeExecutors = freeExecutors[:ln-1]

		res = append(res, NewEPlanGroupByNode(nodea, inputs, outputs))
		ePlanNodes = append(ePlanNodes, res...)
		return res

	case *PlanJoinNode:
		nodea := node.(*PlanJoinNode)
		leftInputNodes := CreateEPlan(nodea.LeftInput, ePlanNodes, freeExecutors, pn)
		rightInputNodes := CreateEPlan(nodea.RightInput, ePlanNodes, freeExecutors, pn)

		duplicateNode := NewEPlanDuplicateNode()
	}
}
