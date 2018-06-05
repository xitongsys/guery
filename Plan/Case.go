package Plan

import (
	"fmt"

	"github.com/xitongsys/guery/Util"
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

func (self *CaseNode) GetColumns() ([]string, error) {
	res, resmp := []string{}, map[string]int{}
	for _, w := range self.Whens {
		cs, err := w.GetColumns()
		if err != nil {
			return res, err
		}
		for _, c := range cs {
			resmp[c] = 1
		}
	}
	cs, err := self.Else.GetColumns()
	if err != nil {
		return res, err
	}
	for _, c := range cs {
		resmp[c] = 1
	}
	for c, _ := range resmp {
		res = append(res, c)
	}
	return res, nil
}

func (self *CaseNode) GetType(md *Util.Metadata) (Util.Type, error) {
	for _, w := range self.Whens {
		return w.GetType(md)
	}
	return Util.UNKNOWNTYPE, fmt.Errorf("unknown type")
}

func (self *CaseNode) Result(input *Util.RowsGroup) (interface{}, error) {
	var res interface{}
	var err error
	for _, w := range self.Whens {
		res, err = w.Result(input)
		if err != nil {
			return nil, err
		}
		if res != nil {
			return res, nil
		}
	}
	return self.Else.Result(input)
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

func (self *WhenClauseNode) GetColumns() ([]string, error) {
	res, resmp := []string{}, map[string]int{}
	cs, err := self.Condition.GetColumns()
	if err != nil {
		return res, err
	}
	for _, c := range cs {
		resmp[c] = 1
	}
	cs, err = self.Res.GetColumns()
	if err != nil {
		return res, err
	}
	for _, c := range cs {
		resmp[c] = 1
	}
	for c, _ := range resmp {
		res = append(res, c)
	}
	return res, nil
}

func (self *WhenClauseNode) GetType(md *Util.Metadata) (Util.Type, error) {
	return self.Res.GetType(md)
}

func (self *WhenClauseNode) Result(input *Util.RowsGroup) (interface{}, error) {
	var res, cd interface{}
	var err error

	cd, err = self.Condition.Result(input)
	if err != nil {
		return nil, err
	}
	if cd.(bool) {
		input.Reset()
		res, err = self.Res.Result(input)
	}
	return res, err
}

func (self *WhenClauseNode) IsAggregate() bool {
	if self.Condition.IsAggregate() || self.Res.IsAggregate() {
		return true
	}
	return false
}
