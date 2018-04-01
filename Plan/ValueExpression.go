package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/Common"
	"github.com/xitongsys/guery/parser"
)

type ValueExpressionNode struct {
	Tree                  *parser.ValueExpressionContext
	PrimaryExpression     *PrimaryExpressionNode
	Operator              *Common.Operator
	ValueExpression       *ValueExpressionNode
	BinaryVauleExpression *BinaryValueExpressionNode
}

func NewValueExpressionNode(ctx *Context, t *parser.ValueExpressionContext) *ValueExpressionNode {
	res := &ValueExpressionNode{
		Tree: t,
	}
	children := t.GetChildren()
	switch len(children) {
	case 1: //PrimaryExpression
		res.PrimaryExpression = NewPrimaryExpressionNode(ctx, children[0].(*parser.PrimaryExpressionContext))

	case 2: //ValueExpression
		if t.MINUS() != nil {
			res.Operator = Common.NewOperator("MINUS")
		} else {
			res.Operator = Common.NewOperator("PLUS")
		}
		res.ValueExpression = NewValueExpressionNode(ctx, children[1].(*parser.ValueExpressionContext))

	case 3: //BinaryValueExpression
		op := Common.NewOperator(children[1].(*antlr.TerminalNodeImpl).GetText())
		res.BinaryVauleExpression = NewBinaryValueExpressionNode(ctx,
			children[0].(*parser.ValueExpressionContext),
			children[2].(*parser.ValueExpressionContext),
			op)
	}
	return res
}

func (self *ValueExpressionNode) Result(ctx *Context) interface{} {
	if self.PrimaryExpression != nil {
		return self.PrimaryExpression.Result(ctx)

	} else if self.ValueExpression != nil {
		if *self.Operator == Common.MINUS {
			return Common.Arithmetic(-1, self.ValueExpression.Result(ctx), Common.ASTERISK)
		}
		return self.ValueExpression.Result(ctx)

	} else if self.BinaryVauleExpression != nil {
		return self.BinaryVauleExpression.Result(ctx)
	}
	return nil
}

/////////////////
type BinaryValueExpressionNode struct {
	LeftValueExpression  *ValueExpressionNode
	RightValueExpression *ValueExpressionNode
	Operator             *Common.Operator
}

func NewBinaryValueExpressionNode(ctx *Context,
	left *parser.ValueExpressionContext, right *parser.ValueExpressionContext,
	op *Common.Operator) *BinaryValueExpressionNode {

	res := &BinaryValueExpressionNode{
		LeftValueExpression:  NewValueExpressionNode(ctx, left),
		RightValueExpression: NewValueExpressionNode(ctx, right),
		Operator:             op,
	}
	return res
}

func (self *BinaryValueExpressionNode) Result(ctx *Context) interface{} {
	leftVal, rightVal := self.LeftValueExpression.Result(ctx), self.RightValueExpression.Result(ctx)
	return Common.Arithmetic(leftVal, rightVal, *self.Operator)
}
