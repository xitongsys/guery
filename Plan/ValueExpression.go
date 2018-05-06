package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/Common"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/parser"
)

type ValueExpressionNode struct {
	Name                  string
	PrimaryExpression     *PrimaryExpressionNode
	Operator              *Common.Operator
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
			res.Operator = Common.NewOperator("-")
			ops = "-"
		} else {
			res.Operator = Common.NewOperator("+")
			ops = "+"
		}
		res.ValueExpression = NewValueExpressionNode(children[1].(parser.IValueExpressionContext))
		res.Name = ops + res.ValueExpression.Name

	case 3: //BinaryValueExpression
		op := Common.NewOperator(children[1].(*antlr.TerminalNodeImpl).GetText())
		res.BinaryVauleExpression = NewBinaryValueExpressionNode(tt.GetLeft(), tt.GetRight(), op)
		res.Name = res.BinaryVauleExpression.Name
	}
	return res
}

func (self *ValueExpressionNode) Result(input *Util.RowsBuffer) (interface{}, error) {
	if self.PrimaryExpression != nil {
		return self.PrimaryExpression.Result(input)

	} else if self.ValueExpression != nil {
		if *self.Operator == Common.MINUS {
			res, err := self.ValueExpression.Result(input)
			if err != nil {
				return nil, err
			}
			return Common.Arithmetic(-1, res, Common.ASTERISK)
		}
		return self.ValueExpression.Result(input)

	} else if self.BinaryVauleExpression != nil {
		return self.BinaryVauleExpression.Result(input)
	}
	return nil
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
	Operator             *Common.Operator
}

func NewBinaryValueExpressionNode(
	left parser.IValueExpressionContext,
	right parser.IValueExpressionContext,
	op *Common.Operator) *BinaryValueExpressionNode {

	res := &BinaryValueExpressionNode{
		LeftValueExpression:  NewValueExpressionNode(left),
		RightValueExpression: NewValueExpressionNode(right),
		Operator:             op,
	}
	res.Name = res.LeftValueExpression.Name + "_" + res.RightValueExpression.Name
	return res
}

func (self *BinaryValueExpressionNode) Result(input *Util.RowsBuffer) (interface{}, error) {
	leftVal, errL := self.LeftValueExpression.Result(input)
	if errL != nil {
		return nil, errL
	}
	rightVal, errR := self.RightValueExpression.Result(input)
	if errR != nil {
		return nil, errR
	}
	return Common.Arithmetic(leftVal, rightVal, *self.Operator), nil
}

func (self *BinaryValueExpressionNode) IsAggregate() bool {
	return self.LeftValueExpression.IsAggregate() || self.RightValueExpression.IsAggregate()
}
