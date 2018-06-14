package Plan

import (
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Row"
	"github.com/xitongsys/guery/Type"
	"github.com/xitongsys/guery/parser"
)

type PredicatedNode struct {
	Name            string
	ValueExpression *ValueExpressionNode
	Predicate       *PredicateNode
}

func NewPredicatedNode(t parser.IPredicatedContext) *PredicatedNode {
	tt := t.(*parser.PredicatedContext)
	res := &PredicatedNode{}
	res.ValueExpression = NewValueExpressionNode(tt.ValueExpression())
	if tp := tt.Predicate(); tp != nil {
		res.Predicate = NewPredicateNode(tp)
	}
	res.Name = res.ValueExpression.Name
	return res
}

func (self *PredicatedNode) GetType(md *Metadata.Metadata) (Type.Type, error) {
	t, err := self.ValueExpression.GetType(md)
	if err != nil {
		return t, err
	}
	if self.Predicate != nil {
		return self.Predicate.GetType(md)
	}
	return t, nil
}

func (self *PredicatedNode) GetColumns() ([]string, error) {
	var (
		err    error
		res    = []string{}
		resmp  = map[string]int{}
		rp, rv = []string{}, []string{}
	)

	rv, err = self.ValueExpression.GetColumns()
	if err != nil {
		return res, err
	}
	if self.Predicate != nil {
		rp, err = self.Predicate.GetColumns()
		if err != nil {
			return res, err
		}
	}
	for _, c := range rv {
		resmp[c] = 1
	}
	for _, c := range rp {
		resmp[c] = 1
	}
	for c, _ := range resmp {
		res = append(res, c)
	}
	return res, nil
}

func (self *PredicatedNode) Result(input *Row.RowsGroup) (interface{}, error) {
	res, err := self.ValueExpression.Result(input)
	if err != nil {
		return nil, err
	}
	if self.Predicate == nil {
		return res, nil
	}
	input.Reset()
	return self.Predicate.Result(res, input)
}

func (self *PredicatedNode) IsAggregate() bool {
	return self.ValueExpression.IsAggregate()
}
