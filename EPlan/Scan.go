package EPlan

import (
	"github.com/xitongsys/guery/Logger"
	. "github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/pb"
)

type EPlanScanNode struct {
	Location   *pb.Location
	SourceName string
	Outputs    []*pb.Location
}

func (self *EPlanScanNode) GetNodeType() EPlanNodeType {
	return ESCANNODE
}

func CreateEPlanScanNode(node *PlanScanNode, pn int32, outputs []*pb.Location) *EPlanScanNode {
	if len(outputs) <= 0 {
		Logger.Errorf("outputs number <= 0")
		return nil
	}
	return &EPlanScanNode{
		Location:   outputs[0],
		SourceName: node.Name,
		Outputs:    outputs,
	}
}
