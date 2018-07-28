package plan

import (
	"fmt"

	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/gtype"
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/parser"
	"github.com/xitongsys/guery/row"
)

type CaseNode struct {
	Whens []*WhenClauseNode
	Else  *ExpressionNode
}

func NewCaseNode(runtime *config.ConfigRuntime, whens []parser.IWhenClauseContext, el parser.IExpressionContext) *CaseNode {
	res := &CaseNode{
		Whens: []*WhenClauseNode{},
		Else:  NewExpressionNode(runtime, el),
	}
	for _, w := range whens {
		res.Whens = append(res.Whens, NewWhenClauseNode(runtime, w))
	}
	return res
}

func (self *CaseNode) ExtractAggFunc(res *[]*FuncCallNode) {
	for _, w := range self.Whens {
		w.ExtractAggFunc(res)
	}
	self.Else.ExtractAggFunc(res)
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

func (self *CaseNode) GetType(md *metadata.Metadata) (gtype.Type, error) {
	for _, w := range self.Whens {
		return w.GetType(md)
	}
	return gtype.UNKNOWNTYPE, fmt.Errorf("unknown type")
}

func (self *CaseNode) Init(md *metadata.Metadata) error {
	for _, w := range self.Whens {
		if err := w.Init(md); err != nil {
			return err
		}
	}
	return nil
}

func (self *CaseNode) Result(input *row.RowsGroup) (interface{}, error) {
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

func NewWhenClauseNode(runtime *config.ConfigRuntime, wh parser.IWhenClauseContext) *WhenClauseNode {
	tt := wh.(*parser.WhenClauseContext)
	ct, rt := tt.GetCondition(), tt.GetResult()
	res := &WhenClauseNode{
		Condition: NewExpressionNode(runtime, ct),
		Res:       NewExpressionNode(runtime, rt),
	}
	return res
}

func (self *WhenClauseNode) ExtractAggFunc(res *[]*FuncCallNode) {
	self.Condition.ExtractAggFunc(res)
	self.Res.ExtractAggFunc(res)
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

func (self *WhenClauseNode) GetType(md *metadata.Metadata) (gtype.Type, error) {
	return self.Res.GetType(md)
}

func (self *WhenClauseNode) Init(md *metadata.Metadata) error {
	if err := self.Condition.Init(md); err != nil {
		return err
	}
	if err := self.Res.Init(md); err != nil {
		return err
	}
	return nil
}

func (self *WhenClauseNode) Result(input *row.RowsGroup) (interface{}, error) {
	var res, cd interface{}
	var err error

	cd, err = self.Condition.Result(input)
	if err != nil {
		return nil, err
	}
	if cd.(bool) {
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
