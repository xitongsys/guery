package Plan

import (
	"github.com/xitongsys/guery/Config"
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Row"
	"github.com/xitongsys/guery/Type"
	"github.com/xitongsys/guery/parser"
)

type GroupingElementNode struct {
	Expression *ExpressionNode
}

func NewGroupingElementNode(runtime *Config.ConfigRuntime, t parser.IGroupingElementContext) *GroupingElementNode {
	res := &GroupingElementNode{}
	tt := t.(*parser.GroupingElementContext).Expression()
	res.Expression = NewExpressionNode(runtime, tt)
	return res
}

func (self *GroupingElementNode) Result(input *Row.RowsGroup) (interface{}, error) {
	return self.Expression.Result(input)
}

func (self *GroupingElementNode) GetColumns() ([]string, error) {
	return self.Expression.GetColumns()
}

func (self *GroupingElementNode) GetType(md *Metadata.Metadata) (Type.Type, error) {
	return self.Expression.GetType(md)
}

func (self *GroupingElementNode) IsAggregate() bool {
	return self.Expression.IsAggregate()
}
