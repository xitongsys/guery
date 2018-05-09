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

func (self *GroupingElementNode) Result(input *Util.RowsBuffer) (interface{}, error) {
	return self.Expression.Result(input)
}
