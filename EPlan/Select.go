package EPlan

import (
	"github.com/xitongsys/guery/Logger"
	. "github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/pb"
)

type EPlanSelectNode struct {
	Location        *pb.Location
	Inputs, Outputs []*pb.Location
	SelectItems     []*SelectItemNode
}

func (self *EPlanSelectNode) GetNodeType() EPlanNodeType {
	return ESELECTNODE
}

func CreateEPlanSelectNodes(node *PlanSelectNode, pn int32, inputs, outputs []*pb.Location) []*EPlanSelectNode {
	res := []*EPlanSelectNode{}
	if len(inputs) != pn || len(outputs) != pn {
		Logger.Errorf("parallel number doesn't match inputs/output number")
		return nil
	}

	for i := 0; i < pn; i++ {
		enode := &EPlanSelectNode{
			Location:    outputs[i],
			Inputs:      []*pb.Location{inputs[i]},
			Outputs:     []*pb.Location{output[i]},
			SelectItems: node.SelectItems,
		}
		res = append(res, enode)
	}
	return res
}
