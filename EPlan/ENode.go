package EPlan

import (
	"fmt"

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
	EAGGREGATENODE
)

type ENode interface {
	GetNodeType() EPlanNodeType
	GetOutputs() []pb.Location
	GetLocation() pb.Location
}
