package Plan

import (
	"fmt"

	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/parser"
)

type GroupByNode struct {
	GroupingElements []*GroupingElementNode
}

func NewGroupByNode(t parser.IGroupByContext) *GroupByNode {
	if t == nil {
		return nil
	}
	res := &GroupByNode{
		GroupingElements: []*GroupingElementNode{},
	}
	tt := t.(*parser.GroupByContext)
	elements := tt.AllGroupingElement()
	for _, element := range elements {
		res.GroupingElements = append(res.GroupingElements, NewGroupingElementNode(element))
	}
	return res
}

func (self *GroupByNode) Result(intput *Util.RowsBuffer) (string, error) {
	res := ""
	for _, element := range self.GroupingElements {
		res += fmt.Sprintf("%v", element.Result(intput))
	}
	return res, nil
}
