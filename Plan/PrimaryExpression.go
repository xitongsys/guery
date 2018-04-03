package Plan

import (
	"github.com/xitongsys/guery/parser"
)

type PrimaryExpressionNode struct {
	//	Null         *NullNode
	Number                  *NumberNode
	BooleanValue            *BooleanValueNode
	StringValue             *StringValueNode
	Identifier              *IdentifierNode
	FuncCall                *FuncCallNode
	ParenthesizedExpression *ExpressionNode
}

func NewPrimaryExpressionNode(t parser.IPrimaryExpressionContext) *PrimaryExpressionNode {
	tt := t.(*parser.PrimaryExpressionContext)
	res := &PrimaryExpressionNode{}
	children := tt.GetChildren()
	if tt.NULL() != nil {
	} else if nu := tt.Number(); nu != nil {
		res.Number = NewNumberNode(nu)

	} else if bv := tt.BooleanValue(); bv != nil {
		res.BooleanValue = NewBooleanValueNode(bv)

	} else if sv := tt.StringValue(); sv != nil {
		res.StringValue = NewStringValueNode(sv)

	} else if id := tt.Identifier(); id != nil {
		res.Identifier = NewIdentifierNode(id)

	} else if tt.QualifiedName() != nil {

	} else {
		res.ParenthesizedExpression = NewExpressionNode(children[1].(parser.IExpressionContext))
	}

	return res
}

func (self *PrimaryExpressionNode) Result(input DataSource.DataSource) interface{} {
	if self.Number != nil {
		return self.Number.Result(input)

	} else if self.BooleanValue != nil {
		return self.BooleanValue.Result(input)

	} else if self.StringValue != nil {
		return self.StringValue.Result(input)

	} else if self.Identifier != nil {
		return self.Identifier.Result(input)

	} else if self.ParenthesizedExpression != nil {
		return self.ParenthesizedExpression.Result(input)

	} else if self.FuncCall != nil {
		return self.FuncCall.Result(input)
	}
	return nil
}

/////////////////////////////
type FuncCallNode struct {
	FuncName    string
	Expressions []*ExpressionNode
}

func NewFuncCallNode(name string, expressions []parser.IExpressionContext) *FuncCallNode {
	res := &FuncCallNode{
		FuncName:    name,
		Expressions: make([]*ExpressionNode, len(expressions)),
	}
	for i := 0; i < len(expressions); i++ {
		res.Expressions[i] = NewExpressionNode(expressions[i])
	}
	return res
}

func (self *FuncCallNode) Result(input DataSource.DataSource) interface{} {
	if Common.GetFuncType(self.Name) == Common.AGGREGATE {
		var res interface{}
		it := input.First()
		if self.Name == "SUM" {
			for it != nil {
				tmp := self.Expressions[0].Result(it)
				it = it.Next()
				if res == nil {
					res = tmp
				} else {
					res = Common.Arithmetic(res, tmp, Common.PLUS)
				}
			}
			return res

		} else if self.Name == "MAX" {
			for it != nil {
				tmp := self.Expressions[0].Result(it)
				it = it.Next()
				if res == nil {
					res = tmp
				} else {
					if Common.Cmp(res, tmp) < 0 {
						res = tmp
					}
				}
			}
			return res

		} else if self.Name == "MIN" {
			for it != nil {
				tmp := self.Expressions[0].Result(it)
				it = it.Next()
				if res == nil {
					res = tmp
				} else {
					if Common.Cmp(res, tmp) > 0 {
						res = tmp
					}
				}
			}
			return res
		}

	} else {

	}
	return nil
}
