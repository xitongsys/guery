package Plan

import (
	"github.com/xitongsys/guery/Metadata"
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

	USENODE
	SHOWNODE
)

type PlanNode interface {
	GetNodeType() PlanNodeType
	SetMetadata() error
	GetMetadata() *Metadata.Metadata

	GetOutput() PlanNode
	SetOutput(output PlanNode)

	GetInputs() []PlanNode
	SetInputs(input []PlanNode)

	String() string
}
