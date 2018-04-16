package Plan

import (
	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/parser"
)

type JoinCriteriaNode struct {
	BooleanExpression *BooleanExpressionNode
	Identifiers       []*IdentifierNode
}

func NewJoinCriteriaNode(ctx *Context.Context, t parser.IJoinCriteriaContext) *JoinCriteriaNode {
	res := &JoinCriteriaNode{}
	tt := t.(*parser.JoinCriteriaContext)
	if be := tt.BooleanExpression(); be != nil {
		res.BooleanExpression = NewBooleanExpressionNode(ctx, be)

	} else {
		ids := tt.AllIdentifier()
		res.Identifiers = []*IdentifierNode{}
		for _, id := range ids {
			res.Identifiers = append(res.Identifiers, NewIdentifierNode(ctx, id))
		}
	}
	return res
}
