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

func (self EPlanNodeType) String() string {
	switch self {
	case ESCANNODE:
		return "scan"
	case ESELECTNODE:
		return "select"
	case EGROUPBYNODE:
		return "group by"
	case EGROUPBYLOCALNODE:
		return "group by local"
	case EFILITERNODE:
		return "filiter"
	case EUNIONNODE:
		return "union"
	case ELIMITNODE:
		return "limit"
	case EORDERBYNODE:
		return "order by"
	case EORDERBYLOCALNODE:
		return "order by local"
	case EJOINNODE:
		return "join"
	case EHASHJOINNODE:
		return "hash join"
	case EHAVINGNODE:
		return "having"
	case ECOMBINENODE:
		return "combine"
	case EDUPLICATENODE:
		return "duplicate"
	case EAGGREGATENODE:
		return "aggregate"
	default:
		return "unknown"
	}
}

type ENode interface {
	GetNodeType() EPlanNodeType
	GetInputs() []pb.Location
	GetOutputs() []pb.Location
	GetLocation() pb.Location
}
