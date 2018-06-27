package Plan

import (
	"fmt"

	"github.com/xitongsys/guery/Config"
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Row"
	"github.com/xitongsys/guery/Type"
	"github.com/xitongsys/guery/parser"
)

type CaseNode struct {
	Whens []*WhenClauseNode
	Else  *ExpressionNode
}

func NewCaseNode(runtime *Config.ConfigRuntime, whens []parser.IWhenClauseContext, el parser.IExpressionContext) *CaseNode {
	res := &CaseNode{
		Whens: []*WhenClauseNode{},
		Else:  NewExpressionNode(runtime, el),
	}
	for _, w := range whens {
		res.Whens = append(res.Whens, NewWhenClauseNode(runtime, w))
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

func (self *CaseNode) GetType(md *Metadata.Metadata) (Type.Type, error) {
	for _, w := range self.Whens {
		return w.GetType(md)
	}
	return Type.UNKNOWNTYPE, fmt.Errorf("unknown type")
}

func (self *CaseNode) Result(input *Row.RowsGroup) (interface{}, error) {
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

func NewWhenClauseNode(runtime *Config.ConfigRuntime, wh parser.IWhenClauseContext) *WhenClauseNode {
	tt := wh.(*parser.WhenClauseContext)
	ct, rt := tt.GetCondition(), tt.GetResult()
	res := &WhenClauseNode{
		Condition: NewExpressionNode(runtime, ct),
		Res:       NewExpressionNode(runtime, rt),
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

func (self *WhenClauseNode) GetType(md *Metadata.Metadata) (Type.Type, error) {
	return self.Res.GetType(md)
}

func (self *WhenClauseNode) Result(input *Row.RowsGroup) (interface{}, error) {
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
