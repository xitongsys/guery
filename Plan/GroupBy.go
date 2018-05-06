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

func (self *GroupByNode) Result(input *Util.RowsBuffer) (string, error) {
	res := ""
	for _, element := range self.GroupingElements {
		er, err := element.Result(input)
		if err != nil {
			return "", err
		}
		res += fmt.Sprintf("%v", er)
	}
	return res, nil
}
