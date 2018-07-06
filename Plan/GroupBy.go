package Plan

import (
	"fmt"

	"github.com/xitongsys/guery/Config"
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Row"
	"github.com/xitongsys/guery/parser"
)

type GroupByNode struct {
	GroupingElements []*GroupingElementNode
	Having           *BooleanExpressionNode
}

func NewGroupByNode(runtime *Config.ConfigRuntime, t parser.IGroupByContext, having parser.IBooleanExpressionContext) *GroupByNode {
	if t == nil {
		return nil
	}
	res := &GroupByNode{
		GroupingElements: []*GroupingElementNode{},
	}
	tt := t.(*parser.GroupByContext)
	elements := tt.AllGroupingElement()
	for _, element := range elements {
		res.GroupingElements = append(res.GroupingElements, NewGroupingElementNode(runtime, element))
	}
	if having != nil {
		res.Having = NewBooleanExpressionNode(runtime, having)
	}
	return res
}

func (self *GroupByNode) Init(md *Metadata.Metadata) error {
	for _, element := range self.GroupingElements {
		if err := element.Init(md); err != nil {
			return err
		}
	}
	return nil
}

func (self *GroupByNode) Result(input *Row.RowsGroup) (string, error) {
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
