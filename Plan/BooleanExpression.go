package Plan

import (
	"github.com/xitongsys/guery/Common"
	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/DataSource"
	"github.com/xitongsys/guery/parser"
)

type BooleanExpressionNode struct {
	Name                    string
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
		res.Predicated = NewPredicatedNode(ctx, tt.Predicated())
		res.Name = res.Predicated.Name

	case 2: //NOT
		res.NotBooleanExpression = NewNotBooleanExpressionNode(ctx, tt.BooleanExpression(0))
		res.Name = res.NotBooleanExpression.Name

	case 3: //Binary
		var o Common.Operator
		if tt.AND() != nil {
			o = Common.AND
		} else if tt.OR() != nil {
			o = Common.OR
		}
		res.BinaryBooleanExpression = NewBinaryBooleanExpressionNode(ctx, tt.GetLeft(), tt.GetRight(), o)
		res.Name = res.BinaryBooleanExpression.Name

	}
	return res
}

func (self *BooleanExpressionNode) Result(input *DataSource.DataSource) interface{} {
	if self.Predicated != nil {
		return self.Predicated.Result(input)
	} else if self.NotBooleanExpression != nil {
		return self.NotBooleanExpression.Result(input)
	} else if self.BinaryBooleanExpression != nil {
		return self.BinaryBooleanExpression.Result(input)
	}
	return nil
}

func (self *BooleanExpressionNode) IsAggregate() bool {
	if self.Predicated != nil {
		return self.Predicated.IsAggregate()
	} else if self.NotBooleanExpression != nil {
		return self.NotBooleanExpression.IsAggregate()
	} else if self.BinaryBooleanExpression != nil {
		return self.BinaryBooleanExpression.IsAggregate()
	}
	return false
}

////////////////////////
type NotBooleanExpressionNode struct {
	Name              string
	BooleanExpression *BooleanExpressionNode
}

func NewNotBooleanExpressionNode(ctx *Context.Context, t parser.IBooleanExpressionContext) *NotBooleanExpressionNode {
	res := &NotBooleanExpressionNode{
		BooleanExpression: NewBooleanExpressionNode(ctx, t),
	}
	res.Name = "NOT_" + res.BooleanExpression.Name
	return res
}

func (self *NotBooleanExpressionNode) Result(input *DataSource.DataSource) bool {
	return !self.BooleanExpression.Result(input).(bool)
}

func (self *NotBooleanExpressionNode) IsAggregate() bool {
	return self.BooleanExpression.IsAggregate()
}

////////////////////////
type BinaryBooleanExpressionNode struct {
	Name                   string
	LeftBooleanExpression  *BooleanExpressionNode
	RightBooleanExpression *BooleanExpressionNode
	Operator               *Common.Operator
}

func NewBinaryBooleanExpressionNode(ctx *Context.Context,
	left parser.IBooleanExpressionContext,
	right parser.IBooleanExpressionContext,
	op Common.Operator) *BinaryBooleanExpressionNode {

	res := &BinaryBooleanExpressionNode{
		LeftBooleanExpression:  NewBooleanExpressionNode(ctx, left),
		RightBooleanExpression: NewBooleanExpressionNode(ctx, right),
		Operator:               &op,
	}
	res.Name = res.LeftBooleanExpression.Name + "_" + res.RightBooleanExpression.Name
	return res
}

func (self *BinaryBooleanExpressionNode) Result(input *DataSource.DataSource) bool {
	if *self.Operator == Common.AND {
		if leftRes := self.LeftBooleanExpression.Result(input).(bool); !leftRes {
			return false
		} else {
			return self.RightBooleanExpression.Result(input).(bool)
		}

	} else if *self.Operator == Common.OR {
		if leftRes := self.LeftBooleanExpression.Result(input).(bool); leftRes {
			return true
		} else {
			return self.RightBooleanExpression.Result(input).(bool)
		}
	}
	return false
}

func (self *BinaryBooleanExpressionNode) IsAggregate() bool {
	return self.LeftBooleanExpression.IsAggregate() || self.RightBooleanExpression.IsAggregate()
}
