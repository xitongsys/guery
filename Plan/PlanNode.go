package Plan

import (
	"github.com/xitongsys/guery/Util"
)

type PlanNodeType int32

const (
	_ PlanNodeType = iota
	SCANNODE
	JOINNODE
	FILTERNODE
	ORDERBYNODE
	LIMITNODE
	SELECTNODE
	UNIONNODE
	HAVINGNODE
	RENAMENODE
	COMBINENODE
	GROUPBYNODE
)

type PlanNode interface {
	GetNodeType() PlanNodeType
	SetMetadata() error
	GetMetadata() *Util.Metadata
	SetOutput(output PlanNode)
	SetInputs(input []PlanNode)
	String() string
}
