package Plan

import (
	"fmt"

	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/DataSource"
	"github.com/xitongsys/guery/parser"
)

type GroupByNode struct {
	GroupingElements []*GroupingElementNode
}

func NewGroupByNode(ctx *Context.Context, t parser.IGroupByContext) *GroupByNode {
	if t == nil {
		return nil
	}
	res := &GroupByNode{
		GroupingElements: []*GroupingElementNode{},
	}
	tt := t.(*parser.GroupByContext)
	elements := tt.AllGroupingElement()
	for _, element := range elements {
		res.GroupingElements = append(res.GroupingElements, NewGroupingElementNode(ctx, element))
	}
	return res
}

func (self *GroupByNode) Result(intput *DataSource.DataSource) string {
	res := ""
	for _, element := range self.GroupingElements {
		res += fmt.Sprintf("%s", element.Result(intput))
	}
	return res
}
