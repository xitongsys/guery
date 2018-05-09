package Plan

import (
	"fmt"

	"github.com/xitongsys/guery/Common"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/parser"
)

type BooleanExpressionNode struct {
	Name                    string
	Predicated              *PredicatedNode
	NotBooleanExpression    *NotBooleanExpressionNode
	BinaryBooleanExpression *BinaryBooleanExpressionNode
}

func NewBooleanExpressionNode(t parser.IBooleanExpressionContext) *BooleanExpressionNode {
	tt := t.(*parser.BooleanExpressionContext)
	res := &BooleanExpressionNode{}
	children := tt.GetChildren()
	switch len(children) {
	case 1: //Predicated
		res.Predicated = NewPredicatedNode(tt.Predicated())
		res.Name = res.Predicated.Name

	case 2: //NOT
		res.NotBooleanExpression = NewNotBooleanExpressionNode(tt.BooleanExpression(0))
		res.Name = res.NotBooleanExpression.Name

	case 3: //Binary
		var o Common.Operator
		if tt.AND() != nil {
			o = Common.AND
		} else if tt.OR() != nil {
			o = Common.OR
		}
		res.BinaryBooleanExpression = NewBinaryBooleanExpressionNode(tt.GetLeft(), tt.GetRight(), o)
		res.Name = res.BinaryBooleanExpression.Name

	}
	return res
}

func (self *BooleanExpressionNode) Result(input *Util.RowsBuffer) (interface{}, error) {
	if self.Predicated != nil {
		return self.Predicated.Result(input)
	} else if self.NotBooleanExpression != nil {
		return self.NotBooleanExpression.Result(input)
	} else if self.BinaryBooleanExpression != nil {
		return self.BinaryBooleanExpression.Result(input)
	}
	return nil, fmt.Errorf("wrong BooleanExpressionNode")
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

func NewNotBooleanExpressionNode(t parser.IBooleanExpressionContext) *NotBooleanExpressionNode {
	res := &NotBooleanExpressionNode{
		BooleanExpression: NewBooleanExpressionNode(t),
	}
	res.Name = "NOT_" + res.BooleanExpression.Name
	return res
}

func (self *NotBooleanExpressionNode) Result(input *Util.RowsBuffer) (bool, error) {
	res, err := self.BooleanExpression.Result(input)
	if err != nil {
		return false, err
	}
	return !res.(bool), nil
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

func NewBinaryBooleanExpressionNode(
	left parser.IBooleanExpressionContext,
	right parser.IBooleanExpressionContext,
	op Common.Operator) *BinaryBooleanExpressionNode {

	res := &BinaryBooleanExpressionNode{
		LeftBooleanExpression:  NewBooleanExpressionNode(left),
		RightBooleanExpression: NewBooleanExpressionNode(right),
		Operator:               &op,
	}
	res.Name = res.LeftBooleanExpression.Name + "_" + res.RightBooleanExpression.Name
	return res
}

func (self *BinaryBooleanExpressionNode) Result(input *Util.RowsBuffer) (bool, error) {
	if *self.Operator == Common.AND {
		leftRes, err := self.LeftBooleanExpression.Result(input)
		if err != nil {
			return false, err
		}
		if !(leftRes.(bool)) {
			return false, nil

		} else {
			rightRes, err := self.RightBooleanExpression.Result(input)
			if err != nil {
				return false, err
			}
			return rightRes.(bool), nil
		}

	} else if *self.Operator == Common.OR {
		leftRes, err := self.LeftBooleanExpression.Result(input)
		if err != nil {
			return false, err
		}
		if leftRes.(bool) {
			return true, nil

		} else {
			rightRes, err := self.RightBooleanExpression.Result(input)
			if err != nil {
				return false, err
			}
			return rightRes.(bool), nil
		}
	}
	return false, fmt.Errorf("wrong BinaryBooleanExpressionNode")
}

func (self *BinaryBooleanExpressionNode) IsAggregate() bool {
	return self.LeftBooleanExpression.IsAggregate() || self.RightBooleanExpression.IsAggregate()
}
