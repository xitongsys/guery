package Plan

import (
	"fmt"

	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/parser"
)

type GroupByNode struct {
	GroupingElements []*GroupingElementNode
	Having           *BooleanExpressionNode
}

func NewGroupByNode(t parser.IGroupByContext, having parser.IBooleanExpressionContext) *GroupByNode {
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
	if having != nil {
		res.Having = NewBooleanExpressionNode(having)
	}
	return res
}

func (self *GroupByNode) Result(input *Util.RowsGroup) (string, error) {
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

func (self *GroupByNode) GetColumns() ([]string, error) {
	res := []string{}
	for _, ele := range self.GroupingElements {
		cs, err := ele.GetColumns()
		if err != nil {
			return res, err
		}
		res = append(res, cs...)
	}
	return res, nil
}
