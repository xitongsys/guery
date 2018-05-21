package EPlan

import (
	. "github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/pb"
)

type EPlanScanNode struct {
	Location   pb.Location
	SourceName string
	Metadata   *Util.Metadata
	Outputs    []pb.Location
}

func (self *EPlanScanNode) GetNodeType() EPlanNodeType {
	return ESCANNODE
}

func (self *EPlanScanNode) GetOutputs() []pb.Location {
	return self.Outputs
}

func (self *EPlanScanNode) GetLocation() pb.Location {
	return self.Location
}

func NewEPlanScanNode(node *PlanScanNode, loc pb.Location, outputs []pb.Location) *EPlanScanNode {
	return &EPlanScanNode{
		Location:   loc,
		SourceName: node.Table,
		Outputs:    outputs,
		Metadata:   node.GetMetadata(),
	}
}
