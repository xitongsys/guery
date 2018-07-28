package eplan

import (
	"github.com/xitongsys/guery/pb"
)

type EPlanNodeType int32

const (
	_ EPlanNodeType = iota
	ESCANNODE
	ESELECTNODE
	EGROUPBYNODE
	EFILTERNODE
	EUNIONNODE
	ELIMITNODE
	EORDERBYNODE
	EORDERBYLOCALNODE
	EJOINNODE
	EHASHJOINNODE
	EHASHJOINSHUFFLENODE
	EHAVINGNODE
	ECOMBINENODE
	EDUPLICATENODE
	EAGGREGATENODE
	EAGGREGATEFUNCLOCALNODE
	EAGGREGATEFUNCGLOBALNODE
	EBALANCENODE

	ESHOWNODE
)

func (self EPlanNodeType) String() string {
	switch self {
	case ESCANNODE:
		return "SCAN"
	case ESELECTNODE:
		return "SELECT"
	case EGROUPBYNODE:
		return "GROUP BY"
	case EFILTERNODE:
		return "FILTER"
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
	case EHASHJOINSHUFFLENODE:
		return "HASH JOIN SHUFFLE"

	case EHAVINGNODE:
		return "HAVING"
	case ECOMBINENODE:
		return "COMBINE"
	case EDUPLICATENODE:
		return "DUPLICATE"
	case EAGGREGATENODE:
		return "AGGREGATE"
	case EAGGREGATEFUNCGLOBALNODE:
		return "AGGREGATE FUNC GLOBAL"
	case EAGGREGATEFUNCLOCALNODE:
		return "AGGREGATE FUNC LOCAL"
	case ESHOWNODE:
		return "SHOW"
	case EBALANCENODE:
		return "BALANCE"
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
