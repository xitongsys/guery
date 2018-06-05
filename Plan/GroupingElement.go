package Plan

import (
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/parser"
)

type GroupingElementNode struct {
	Expression *ExpressionNode
}

func NewGroupingElementNode(t parser.IGroupingElementContext) *GroupingElementNode {
	res := &GroupingElementNode{}
	tt := t.(*parser.GroupingElementContext).Expression()
	res.Expression = NewExpressionNode(tt)
	return res
}

func (self *GroupingElementNode) Result(input *Util.RowsGroup) (interface{}, error) {
	return self.Expression.Result(input)
}

func (self *GroupingElementNode) GetColumns() ([]string, error) {
	return self.Expression.GetColumns()
}

func (self *GroupingElementNode) GetType(md *Util.Metadata) (Util.Type, error) {
	return self.Expression.GetType(md)
}

func (self *GroupingElementNode) IsAggregate() bool {
	return self.Expression.IsAggregate()
}
