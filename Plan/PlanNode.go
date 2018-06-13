package Plan

import (
	"github.com/xitongsys/guery/Util"
)

type PlanNodeType int32

const (
	_ PlanNodeType = iota
	SCANNODE
	JOINNODE
	HASHJOINNODE
	FILTERNODE
	ORDERBYNODE
	LIMITNODE
	SELECTNODE
	UNIONNODE
	HAVINGNODE
	RENAMENODE
	COMBINENODE
	GROUPBYNODE
	AGGREGATENODE
)

type PlanNode interface {
	GetNodeType() PlanNodeType
	SetMetadata() error
	GetMetadata() *Util.Metadata

	GetOutput() PlanNode
	SetOutput(output PlanNode)

	GetInputs() []PlanNode
	SetInputs(input []PlanNode)

	String() string
}
