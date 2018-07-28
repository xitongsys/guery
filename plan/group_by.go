package plan

import (
	"fmt"

	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/parser"
	"github.com/xitongsys/guery/row"
)

type GroupByNode struct {
	GroupingElements []*GroupingElementNode
}

func NewGroupByNode(runtime *config.ConfigRuntime, t parser.IGroupByContext) *GroupByNode {
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
	return res
}

func (self *GroupByNode) Init(md *metadata.Metadata) error {
	for _, element := range self.GroupingElements {
		if err := element.Init(md); err != nil {
			return err
		}
	}
	return nil
}

func (self *GroupByNode) Result(input *row.RowsGroup) ([]interface{}, error) {
	rn := input.GetRowsNumber()
	res := make([]interface{}, rn)
	for i := 0; i < rn; i++ {
		res[i] = ""
	}

	for _, element := range self.GroupingElements {
		esi, err := element.Result(input)
		if err != nil {
			return nil, err
		}
		es := esi.([]interface{})
		for i, e := range es {
			res[i] = res[i].(string) + fmt.Sprintf("%v:", e)
		}
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
