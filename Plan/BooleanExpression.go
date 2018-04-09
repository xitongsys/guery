package Plan

import (
	"github.com/xitongsys/guery/Common"
	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/parser"
)

type BooleanExpressionNode struct {
	Predicated              *PredicatedNode
	NotBooleanExpression    *NotBooleanExpressionNode
	BinaryBooleanExpression *BinaryBooleanExpressionNode
}

func NewBooleanExpressionNode(ctx *Context.Context, t parser.IBooleanExpressionContext) *BooleanExpressionNode {
	tt := t.(*parser.BooleanExpressionContext)
	res := &BooleanExpressionNode{}
	children := tt.GetChildren()
	switch len(children) {
	case 1: //Predicated
		res.Predicated = NewPredicatedNode(tt.Predicated())

	case 2: //NOT
		res.NotBooleanExpression = NewNotBooleanExpressionNode(tt.BooleanExpression())

	case 3: //Binary
		var o Common.Operator
		if t.AND() != nil {
			o = Common.AND
		} else if t.OR() != nil {
			o = Common.OR
		}
		res.BinaryBooleanExpression = NewBinaryBooleanExpressionNode(tt.GetLeft(), tt.GetRight(), &o)

	}
	return res
}

func (self *BooleanExpressionNode) Result(input DataSource.DataSource) interface{} {
	if self.Predicated != nil {
		return self.Predicated.Result(input)
	} else if self.NotBooleanExpression != nil {
		return self.NotBooleanExpression.Result(input)
	} else if self.BinaryBooleanExpression != nil {
		return self.BinaryBooleanExpression.Result(input)
	}
	return nil
}

////////////////////////
type NotBooleanExpressionNode struct {
	BooleanExpression *BooleanExpressionNode
}

func NewNotBooleanExpressionNode(t parser.IBooleanExpressionContext) *NotBooleanExpressionNode {
	res := &NotBooleanExpressionNode{
		BooleanExpression: NewBooleanExpressionNode(t),
	}
	return res
}

func (self *NotBooleanExpressionNode) Result(input DataSource.DataSource) bool {
	return !self.BooleanExpression.Result(input).(bool)
}

////////////////////////
type BinaryBooleanExpressionNode struct {
	LeftBooleanExpression  *BooleanExpressionNode
	RightBooleanExpression *BooleanExpressionNode
	Operator               *Common.Operator
}

func NewBinaryBooleanExpressionNode(left parser.IBooleanExpressionContext,
	right parser.IBooleanExpressionContext,
	op Common.Operator) *BinaryBooleanExpressionNode {

	res := &BinaryBooleanExpressionNode{
		LeftBooleanExpression:  NewBooleanExpressionNode(left),
		RightBooleanExpression: NewBooleanExpressionNode(right),
		Operator:               op,
	}
	return res
}

func (self *BinaryBooleanExpressionNode) Result(input DataSource.DataSource) bool {
	if self.Operator == Common.AND {
		if leftRes := self.LeftBooleanExpression.Result(input).(bool); !leftRes {
			return false
		} else {
			return self.RightBooleanExpression.Result(input).(bool)
		}

	} else if self.Operator == Common.OR {
		if leftRes := self.LeftBooleanExpression.Result(input).(bool); leftRes {
			return true
		} else {
			return self.RightBooleanExpression.Result(input).(bool)
		}
	}
	return false
}
