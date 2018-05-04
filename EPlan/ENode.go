package EPlan

import (
	"github.com/xitongsys/guery/Logger"
	. "github.com/xitongsys/guery/Plan"
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
	EDUPLICATE
)

type ENode interface {
	GetNodeType() EPlanNodeType
}
