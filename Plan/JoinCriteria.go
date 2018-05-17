package Plan

import (
	"github.com/xitongsys/guery/Util"
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

func (self *JoinCriteriaNode) GetColumns(md *Util.Metadata) ([]string, error) {
	if self.BooleanExpression != nil {
		return self.BooleanExpression.GetColumns(md)
	} else {
		res := []string{}
		for _, id := range self.Identifier {
			r, err := id.GetColumns(md)
			if err != nil {
				return res, err
			}
			res = append(res, r)
		}
		return res, nil
	}
}

func (self *JoinCriteriaNode) Result(input *Util.RowsBuffer) (bool, error) {
	if self.BooleanExpression != nil {
		res, err := self.BooleanExpression.Result(input)
		return res.(bool), err
	} else {
		return true, nil
	}
}
