package Plan

import (
	"github.com/xitongsys/guery/DataSource"
	"github.com/xitongsys/guery/parser"
)

type CaseNode struct {
	Whens []*WhenClauseNode
	Else  *ExpressionNode
}

func NewCaseNode(whens []parser.IWhenClauseContext, el parser.IExpressionContext) *CaseNode {
	res := &CaseNode{
		Whens: []*WhenClauseNode{},
		Else:  NewExpressionNode(el),
	}
	for _, w := range whens {
		res.Whens = append(res.Whens, NewWhenClauseNode(w))
	}
	return res
}

func (self *CaseNode) Result(input *DataSource.DataSource) interface{} {
	input.Reset()
	var res interface{}
	for _, w := range self.Whens {
		res = w.Result(input)
		if res != nil {
			return res
		}
	}
	res = self.Else.Result(input)
	return res
}

func (self *CaseNode) IsAggregate() bool {
	for _, w := range self.Whens {
		if w.IsAggregate() {
			return true
		}
	}
	if self.Else != nil && self.Else.IsAggregate() {
		return true
	}
	return false
}

////////
type WhenClauseNode struct {
	Condition *ExpressionNode
	Res       *ExpressionNode
}

func NewWhenClauseNode(wh parser.IWhenClauseContext) *WhenClauseNode {
	tt := wh.(*parser.WhenClauseContext)
	ct, rt := tt.GetCondition(), tt.GetResult()
	res := &WhenClauseNode{
		Condition: NewExpressionNode(ct),
		Res:       NewExpressionNode(rt),
	}
	return res
}

func (self *WhenClauseNode) Result(input *DataSource.DataSource) interface{} {
	input.Reset()
	var res interface{} = nil
	if self.Condition.Result(input).(bool) {
		res = self.Res.Result(input)
	}
	return res
}

func (self *WhenClauseNode) IsAggregate() bool {
	if self.Condition.IsAggregate() || self.Res.IsAggregate() {
		return true
	}
	return false
}
