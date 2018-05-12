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
	Index      int32
	TotalNum   int32
	Metadata   *Util.Metadata
	Output     pb.Location
}

func (self *EPlanScanNode) GetNodeType() EPlanNodeType {
	return ESCANNODE
}

func (self *EPlanScanNode) GetOutputs() []pb.Location {
	return []pb.Location{self.Output}
}

func (self *EPlanScanNode) GetLocation() pb.Location {
	return self.Location
}

func NewEPlanScanNode(node *PlanScanNode, index, totalNum int32, output pb.Location) *EPlanScanNode {
	return &EPlanScanNode{
		Location:   output,
		SourceName: node.Name,
		Output:     output,
		Index:      index,
		TotalNum:   totalNum,
		Metadata:   node.GetMetadata(),
	}
}
