package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/Common"
	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/DataSource"
	"github.com/xitongsys/guery/parser"
)

type ValueExpressionNode struct {
	PrimaryExpression     *PrimaryExpressionNode
	Operator              *Common.Operator
	ValueExpression       *ValueExpressionNode
	BinaryVauleExpression *BinaryValueExpressionNode
}

func NewValueExpressionNode(ctx *Context.Context, t parser.IValueExpressionContext) *ValueExpressionNode {
	tt := t.(*parser.ValueExpressionContext)
	res := &ValueExpressionNode{}
	children := t.GetChildren()
	switch len(children) {
	case 1: //PrimaryExpression
		res.PrimaryExpression = NewPrimaryExpressionNode(ctx, tt.PrimaryExpression())

	case 2: //ValueExpression
		if tt.MINUS() != nil {
			res.Operator = Common.NewOperator("-")
		} else {
			res.Operator = Common.NewOperator("+")
		}
		res.ValueExpression = NewValueExpressionNode(ctx, children[1].(parser.IValueExpressionContext))

	case 3: //BinaryValueExpression
		op := Common.NewOperator(children[1].(*antlr.TerminalNodeImpl).GetText())
		res.BinaryVauleExpression = NewBinaryValueExpressionNode(ctx, tt.GetLeft(), tt.GetRight(), op)
	}
	return res
}

func (self *ValueExpressionNode) Result(input DataSource.DataSource) interface{} {
	if self.PrimaryExpression != nil {
		return self.PrimaryExpression.Result(input)

	} else if self.ValueExpression != nil {
		if *self.Operator == Common.MINUS {
			return Common.Arithmetic(-1, self.ValueExpression.Result(input), Common.ASTERISK)
		}
		return self.ValueExpression.Result(input)

	} else if self.BinaryVauleExpression != nil {
		return self.BinaryVauleExpression.Result(input)
	}
	return nil
}

/////////////////
type BinaryValueExpressionNode struct {
	LeftValueExpression  *ValueExpressionNode
	RightValueExpression *ValueExpressionNode
	Operator             *Common.Operator
}

func NewBinaryValueExpressionNode(ctx *Context.Context,
	left parser.IValueExpressionContext,
	right parser.IValueExpressionContext,
	op *Common.Operator) *BinaryValueExpressionNode {

	res := &BinaryValueExpressionNode{
		LeftValueExpression:  NewValueExpressionNode(ctx, left),
		RightValueExpression: NewValueExpressionNode(ctx, right),
		Operator:             op,
	}
	return res
}

func (self *BinaryValueExpressionNode) Result(input DataSource.DataSource) interface{} {
	leftVal, rightVal := self.LeftValueExpression.Result(input), self.RightValueExpression.Result(input)
	return Common.Arithmetic(leftVal, rightVal, *self.Operator)
}
