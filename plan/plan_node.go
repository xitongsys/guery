package plan

import (
	"github.com/xitongsys/guery/metadata"
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
	AGGREGATEFUNCLOCALNODE
	AGGREGATEFUNCGLOBALNODE
	SHUFFLENODE
	DISTINCTLOCALNODE
	DISTINCTGLOBALNODE

	USENODE
	SHOWNODE
)

type PlanNode interface {
	GetNodeType() PlanNodeType
	SetMetadata() error
	GetMetadata() *metadata.Metadata

	GetOutput() PlanNode
	SetOutput(output PlanNode)

	GetInputs() []PlanNode
	SetInputs(input []PlanNode)

	String() string
}
