package Plan

import (
	"github.com/xitongsys/guery/DataSource"
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

func (self *GroupingElementNode) Result(input *DataSource.DataSource) interface{} {
	return self.Expression.Result(input)
}
