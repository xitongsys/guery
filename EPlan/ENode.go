package EPlan

import (
	"github.com/xitongsys/guery/pb"
)

type EPlanNodeType int32

const (
	_ EPlanNodeType = iota
	ESCANNODE
	ESELECTNODE
	EGROUPBYNODE
	EGROUPBYLOCALNODE
	EFILITERNODE
	EUNIONNODE
	ELIMITNODE
	EORDERBYNODE
	EORDERBYLOCALNODE
	EJOINNODE
	EHASHJOINNODE
	EHAVINGNODE
	ECOMBINENODE
	EDUPLICATENODE
	EAGGREGATENODE
)

type ENode interface {
	GetNodeType() EPlanNodeType
	GetOutputs() []pb.Location
	GetLocation() pb.Location
}
