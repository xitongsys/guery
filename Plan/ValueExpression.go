package Plan

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/parser"
)

type ValueExpressionNode struct {
	Name                  string
	PrimaryExpression     *PrimaryExpressionNode
	Operator              *Util.Operator
	ValueExpression       *ValueExpressionNode
	BinaryVauleExpression *BinaryValueExpressionNode
}

func NewValueExpressionNode(t parser.IValueExpressionContext) *ValueExpressionNode {
	tt := t.(*parser.ValueExpressionContext)
	res := &ValueExpressionNode{}
	children := t.GetChildren()
	switch len(children) {
	case 1: //PrimaryExpression
		res.PrimaryExpression = NewPrimaryExpressionNode(tt.PrimaryExpression())
		res.Name = res.PrimaryExpression.Name

	case 2: //ValueExpression
		ops := "+"
		if tt.MINUS() != nil {
			res.Operator = Util.NewOperatorFromString("-")
			ops = "-"
		} else {
			res.Operator = Util.NewOperatorFromString("+")
			ops = "+"
		}
		res.ValueExpression = NewValueExpressionNode(children[1].(parser.IValueExpressionContext))
		res.Name = ops + res.ValueExpression.Name

	case 3: //BinaryValueExpression
		op := Util.NewOperatorFromString(children[1].(*antlr.TerminalNodeImpl).GetText())
		res.BinaryVauleExpression = NewBinaryValueExpressionNode(tt.GetLeft(), tt.GetRight(), op)
		res.Name = res.BinaryVauleExpression.Name
	}
	return res
}

func (self *ValueExpressionNode) GetType(md *Util.Metadata) (Util.Type, error) {
	if self.PrimaryExpression != nil {
		return self.PrimaryExpression.GetType(md)
	} else if self.ValueExpression != nil {
		return self.ValueExpression.GetType(md)
	} else if self.BinaryVauleExpression != nil {
		return self.BinaryVauleExpression.GetType(md)
	}
	return Util.UNKNOWNTYPE, fmt.Errorf("ValueExpressionNode type error")
}

func (self *ValueExpressionNode) GetColumns() ([]string, error) {
	if self.PrimaryExpression != nil {
		return self.PrimaryExpression.GetColumns()
	} else if self.ValueExpression != nil {
		return self.PrimaryExpression.GetColumns()
	} else if self.BinaryVauleExpression != nil {
		return self.BinaryVauleExpression.GetColumns()
	}
	return []string{}, fmt.Errorf("ValueExpression node error")
}

func (self *ValueExpressionNode) Result(input *Util.RowsGroup) (interface{}, error) {
	if self.PrimaryExpression != nil {
		return self.PrimaryExpression.Result(input)

	} else if self.ValueExpression != nil {
		if *self.Operator == Util.MINUS {
			res, err := self.ValueExpression.Result(input)
			if err != nil {
				return nil, err
			}
			return Util.OperatorFunc(-1, res, Util.ASTERISK), nil
		}
		return self.ValueExpression.Result(input)

	} else if self.BinaryVauleExpression != nil {
		return self.BinaryVauleExpression.Result(input)
	}
	return nil, fmt.Errorf("wrong ValueExpressionNode")
}

func (self *ValueExpressionNode) IsAggregate() bool {
	if self.PrimaryExpression != nil {
		return self.PrimaryExpression.IsAggregate()

	} else if self.ValueExpression != nil {
		return self.ValueExpression.IsAggregate()

	} else if self.BinaryVauleExpression != nil {
		return self.BinaryVauleExpression.IsAggregate()
	}
	return false
}

/////////////////
type BinaryValueExpressionNode struct {
	Name                 string
	LeftValueExpression  *ValueExpressionNode
	RightValueExpression *ValueExpressionNode
	Operator             *Util.Operator
}

func NewBinaryValueExpressionNode(
	left parser.IValueExpressionContext,
	right parser.IValueExpressionContext,
	op *Util.Operator) *BinaryValueExpressionNode {

	res := &BinaryValueExpressionNode{
		LeftValueExpression:  NewValueExpressionNode(left),
		RightValueExpression: NewValueExpressionNode(right),
		Operator:             op,
	}
	res.Name = res.LeftValueExpression.Name + "_" + res.RightValueExpression.Name
	return res
}

func (self *BinaryValueExpressionNode) GetType(md *Util.Metadata) (Util.Type, error) {
	lt, errL := self.LeftValueExpression.GetType(md)
	if errL != nil {
		return lt, errL
	}
	rt, errR := self.RightValueExpression.GetType(md)
	if errR != nil {
		return rt, errR
	}
	return Util.CheckType(lt, rt, *self.Operator)
}

func (self *BinaryValueExpressionNode) GetColumns() ([]string, error) {
	res, resmp := []string{}, map[string]int{}
	rl, err := self.LeftValueExpression.GetColumns()
	if err != nil {
		return res, err
	}
	rr, err := self.RightValueExpression.GetColumns()
	if err != nil {
		return res, err
	}
	for _, c := range rl {
		resmp[c] = 1
	}
	for _, c := range rr {
		resmp[c] = 1
	}
	for c, _ := range resmp {
		res = append(res, c)
	}
	return res, nil
}

func (self *BinaryValueExpressionNode) Result(input *Util.RowsGroup) (interface{}, error) {
	leftVal, errL := self.LeftValueExpression.Result(input)
	if errL != nil {
		return nil, errL
	}
	rightVal, errR := self.RightValueExpression.Result(input)
	if errR != nil {
		return nil, errR
	}
	return Util.OperatorFunc(leftVal, rightVal, *self.Operator), nil
}

func (self *BinaryValueExpressionNode) IsAggregate() bool {
	return self.LeftValueExpression.IsAggregate() || self.RightValueExpression.IsAggregate()
}
