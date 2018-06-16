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

	ESHOWTABLESNODE
)

func (self EPlanNodeType) String() string {
	switch self {
	case ESCANNODE:
		return "SCAN"
	case ESELECTNODE:
		return "SELECT"
	case EGROUPBYNODE:
		return "GROUP BY"
	case EGROUPBYLOCALNODE:
		return "GROUP BY LOCAL"
	case EFILITERNODE:
		return "FILITER"
	case EUNIONNODE:
		return "UNION"
	case ELIMITNODE:
		return "LIMIT"
	case EORDERBYNODE:
		return "ORDER BY"
	case EORDERBYLOCALNODE:
		return "ORDER BY LOCAL"
	case EJOINNODE:
		return "JOIN"
	case EHASHJOINNODE:
		return "HASH JOIN"
	case EHAVINGNODE:
		return "HAVING"
	case ECOMBINENODE:
		return "COMBINE"
	case EDUPLICATENODE:
		return "DUPLICATE"
	case EAGGREGATENODE:
		return "AGGREGATE"
	default:
		return "UNKNOWN"
	}
}

type ENode interface {
	GetNodeType() EPlanNodeType
	GetInputs() []pb.Location
	GetOutputs() []pb.Location
	GetLocation() pb.Location
}
