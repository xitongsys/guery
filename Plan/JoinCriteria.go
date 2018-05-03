package Plan

import (
	"github.com/xitongsys/guery/DataSource"
	"github.com/xitongsys/guery/parser"
)

type JoinCriteriaNode struct {
	BooleanExpression *BooleanExpressionNode
	Identifiers       []*IdentifierNode
}

func NewJoinCriteriaNode(t parser.IJoinCriteriaContext) *JoinCriteriaNode {
	res := &JoinCriteriaNode{}
	tt := t.(*parser.JoinCriteriaContext)
	if be := tt.BooleanExpression(); be != nil {
		res.BooleanExpression = NewBooleanExpressionNode(be)

	} else {
		ids := tt.AllIdentifier()
		res.Identifiers = []*IdentifierNode{}
		for _, id := range ids {
			res.Identifiers = append(res.Identifiers, NewIdentifierNode(id))
		}
	}
	return res
}

func (self *JoinCriteriaNode) Result(input *DataSource.DataSource) bool {
	if self.BooleanExpression != nil {
		return self.BooleanExpression.Result(input).(bool)
	} else {
		return true
	}
}
