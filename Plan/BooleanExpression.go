package Plan

import (
	"github.com/xitongsys/guery/Common"
	"github.com/xitongsys/guery/parser"
)

type BooleanExpressionNode struct {
	Tree                    *parser.BooleanExpressionContext
	Predicated              *PredicatedNode
	NotBooleanExpression    *NotBooleanExpressionNode
	BinaryBooleanExpression *BinaryBooleanExpressionNode
}

func NewBooleanExpressionNode(ctx *Context, t *parser.BooleanExpressionContext) *BooleanExpressionNode {
	res := &BooleanExpressionNode{
		Tree: t,
	}
	children := t.GetChildren()
	switch len(children) {
	case 1: //Predicated
		res.Predicated = NewPredicatedNode(ctx,
			t.Predicated().(*parser.PredicatedContext))

	case 2: //NOT
		res.NotBooleanExpression = NewNotBooleanExpressionNode(ctx,
			children[1].(*parser.BooleanExpressionContext))

	case 3: //Binary
		var o Common.Operator
		if t.AND() != nil {
			o = Common.AND
		} else if t.OR() != nil {
			o = Common.OR
		}

		res.BinaryBooleanExpression = NewBinaryBooleanExpressionNode(ctx,
			children[0].(*parser.BooleanExpressionContext),
			children[2].(*parser.BooleanExpressionContext),
			&o)

	}
	return res
}

func (self *BooleanExpressionNode) Result(ctx *Context) interface{} {
	if self.Predicated != nil {
		return self.Predicated.Result(ctx)
	}
	return nil
}

////////////////////////
type NotBooleanExpressionNode struct {
	BooleanExpression *BooleanExpressionNode
}

func NewNotBooleanExpressionNode(ctx *Context, t *parser.BooleanExpressionContext) *NotBooleanExpressionNode {
	res := &NotBooleanExpressionNode{
		BooleanExpression: NewBooleanExpressionNode(ctx, t),
	}
	return res
}

func (self *NotBooleanExpressionNode) Result(ctx *Context) bool {
	return !self.BooleanExpression.Result(ctx).(bool)
}

////////////////////////
type BinaryBooleanExpressionNode struct {
	LeftBooleanExpression  *BooleanExpressionNode
	RightBooleanExpression *BooleanExpressionNode
	Operator               *Common.Operator
}

func NewBinaryBooleanExpressionNode(ctx *Context,
	left *parser.BooleanExpressionContext, right *parser.BooleanExpressionContext, o *Common.Operator) *BinaryBooleanExpressionNode {
	res := &BinaryBooleanExpressionNode{
		LeftBooleanExpression:  NewBooleanExpressionNode(ctx, left),
		RightBooleanExpression: NewBooleanExpressionNode(ctx, right),
		Operator:               o,
	}
	return res
}

func (self *BinaryBooleanExpressionNode) Result(ctx *Context) bool {
	if *self.Operator == Common.AND {
		if leftRes := self.LeftBooleanExpression.Result(ctx).(bool); !leftRes {
			return false
		} else {
			return self.RightBooleanExpression.Result(ctx).(bool)
		}

	} else if *self.Operator == Common.OR {
		if leftRes := self.LeftBooleanExpression.Result(ctx).(bool); leftRes {
			return true
		} else {
			return self.RightBooleanExpression.Result(ctx).(bool)
		}
	}
	return false
}
