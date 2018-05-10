package EPlan

import (
	"github.com/xitongsys/guery/Logger"
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

func NewEPlanScanNode(node *PlanScanNode, outputs []pb.Location) *EPlanScanNode {
	if len(outputs) <= 0 {
		Logger.Errorf("outputs number <= 0")
		return nil
	}
	return &EPlanScanNode{
		Location:   outputs[0],
		SourceName: node.Name,
		Outputs:    outputs,
		Metadata:   node.GetMetadata(),
	}
}
