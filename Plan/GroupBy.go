package Plan

import (
	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/parser"
)

type GroupByNode struct {
	SetQuantifier    *SetQuantifierNode
	GroupingElements []*GroupingElementNode
}

func NewGroupByNode(ctx *Context.Context, t parser.IGroupByContext) *GroupByNode {
	res := &GroupByNode{
		GroupingElements: []*GroupingElementNode{},
	}
	tt := t.(*parser.GroupByContext)
	if sq := tt.SetQuantifier(); sq != nil {
		res.SetQuantifier = sq
	}

	elements := tt.AllGroupingElement()
	for _, element := range elements {
		res.GroupingElements = append(res.GroupingElements, NewGroupingElementNode(element))
	}

	return res
}
