package plan

import (
	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/parser"
	"github.com/xitongsys/guery/row"
)

type JoinCriteriaNode struct {
	BooleanExpression *BooleanExpressionNode
	Identifiers       []*IdentifierNode
}

func NewJoinCriteriaNode(runtime *config.ConfigRuntime, t parser.IJoinCriteriaContext) *JoinCriteriaNode {
	res := &JoinCriteriaNode{}
	tt := t.(*parser.JoinCriteriaContext)
	if be := tt.BooleanExpression(); be != nil {
		res.BooleanExpression = NewBooleanExpressionNode(runtime, be)

	} else {
		ids := tt.AllIdentifier()
		res.Identifiers = []*IdentifierNode{}
		for _, id := range ids {
			res.Identifiers = append(res.Identifiers, NewIdentifierNode(runtime, id))
		}
	}
	return res
}

func (self *JoinCriteriaNode) GetColumns() ([]string, error) {
	if self.BooleanExpression != nil {
		return self.BooleanExpression.GetColumns()
	} else {
		res := []string{}
		for _, id := range self.Identifiers {
			r, err := id.GetColumns()
			if err != nil {
				return res, err
			}
			res = append(res, r...)
		}
		return res, nil
	}
}

func (self *JoinCriteriaNode) Init(md *metadata.Metadata) error {
	if self.BooleanExpression != nil {
		return self.BooleanExpression.Init(md)
	}
	return nil
}

func (self *JoinCriteriaNode) Result(input *row.RowsGroup) (bool, error) {
	if self.BooleanExpression != nil {
		res, err := self.BooleanExpression.Result(input)
		return res.([]interface{})[0].(bool), err
	} else {
		return true, nil
	}
}
