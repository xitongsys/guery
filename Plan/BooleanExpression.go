package Plan

import (
	"fmt"

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
		var o Util.Operator
		if tt.AND() != nil {
			o = Util.AND
		} else if tt.OR() != nil {
			o = Util.OR
		}
		res.BinaryBooleanExpression = NewBinaryBooleanExpressionNode(tt.GetLeft(), tt.GetRight(), o)
		res.Name = res.BinaryBooleanExpression.Name

	}
	return res
}

func (self *BooleanExpressionNode) GetType(md *Util.Metadata) (Util.Type, error) {
	if self.Predicated != nil {
		return self.Predicated.GetType(md)
	} else if self.NotBooleanExpression != nil {
		return self.NotBooleanExpression.GetType(md)
	} else if self.BinaryBooleanExpression != nil {
		return self.BinaryBooleanExpression.GetType(md)
	}
	return Util.UNKNOWNTYPE, fmt.Errorf("GetType: wrong BooleanExpressionNode")
}

func (self *BooleanExpressionNode) GetColumns(md *Util.Metadata) ([]string, error) {
	if self.Predicated != nil {
		return self.Predicated.GetColumns(md)
	} else if self.NotBooleanExpression != nil {
		return self.NotBooleanExpression.GetColumns(md)
	} else if self.BinaryBooleanExpression != nil {
		return self.BinaryBooleanExpression.GetColumns(md)
	}
	return nil, fmt.Errorf("GetColumns: wrong BooleanExpressionNode")
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

func (self *NotBooleanExpressionNode) GetType(md *Util.Metadata) (Util.Type, error) {
	t, err := self.BooleanExpression.GetType(md)
	if err != nil {
		return t, err
	}
	if t != Util.BOOL {
		return t, fmt.Errorf("expression type error")
	}
	return t, nil
}

func (self *NotBooleanExpressionNode) GetColumns(md *Util.Metadata) ([]string, error) {
	return self.BooleanExpression.GetColumns(md)
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
	Operator               *Util.Operator
}

func NewBinaryBooleanExpressionNode(
	left parser.IBooleanExpressionContext,
	right parser.IBooleanExpressionContext,
	op Util.Operator) *BinaryBooleanExpressionNode {

	res := &BinaryBooleanExpressionNode{
		LeftBooleanExpression:  NewBooleanExpressionNode(left),
		RightBooleanExpression: NewBooleanExpressionNode(right),
		Operator:               &op,
	}
	res.Name = res.LeftBooleanExpression.Name + "_" + res.RightBooleanExpression.Name
	return res
}

func (self *BinaryBooleanExpressionNode) GetType(md *Util.Metadata) (Util.Type, error) {
	lt, err1 := self.LeftBooleanExpression.GetType(md)
	if err1 != nil {
		return Util.UNKNOWNTYPE, err1
	}
	if lt != Util.BOOL {
		return lt, fmt.Errorf("expression type error")
	}
	rt, err2 := self.RightBooleanExpression.GetType(md)
	if err2 != nil {
		return Util.UNKNOWNTYPE, err2
	}
	if rt != Util.BOOL {
		return rt, fmt.Errorf("expression type error")
	}

	return Util.BOOL, nil
}

func (self *BinaryBooleanExpressionNode) GetColumns(md *Util.Metadata) ([]string, error) {
	resmp := make(map[string]int)
	res := []string{}
	rl, errl := self.LeftBooleanExpression.GetColumns(md)
	if errl != nil {
		return res, errl
	}
	rr, errr := self.RightBooleanExpression.GetColumns(md)
	if errr != nil {
		return res, errr
	}
	for _, c := range rl {
		resmp[c] = 1
	}
	for _, c := range rr {
		resmp[c] = 1
	}
	for key, _ := range resmp {
		res = append(res, key)
	}
	return res, nil
}

func (self *BinaryBooleanExpressionNode) Result(input *Util.RowsBuffer) (bool, error) {
	if *self.Operator == Util.AND {
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

	} else if *self.Operator == Util.OR {
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
