package Plan

import (
	"fmt"

	"github.com/xitongsys/guery/Common"
	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/DataSource"
	"github.com/xitongsys/guery/parser"
)

type PrimaryExpressionNode struct {
	//	Null         *NullNode
	Name                    string
	Number                  *NumberNode
	BooleanValue            *BooleanValueNode
	StringValue             *StringValueNode
	Identifier              *IdentifierNode
	FuncCall                *FuncCallNode
	ParenthesizedExpression *ExpressionNode
	Base                    *PrimaryExpressionNode
	FieldName               *IdentifierNode
}

func NewPrimaryExpressionNode(ctx *Context.Context, t parser.IPrimaryExpressionContext) *PrimaryExpressionNode {
	tt := t.(*parser.PrimaryExpressionContext)
	res := &PrimaryExpressionNode{}
	children := tt.GetChildren()
	if tt.NULL() != nil {
		res.Name = "NULL"

	} else if nu := tt.Number(); nu != nil {
		res.Number = NewNumberNode(ctx, nu)
		res.Name = "COL_" + res.Number.Name

	} else if bv := tt.BooleanValue(); bv != nil {
		res.BooleanValue = NewBooleanValueNode(ctx, bv)
		res.Name = "COL_" + res.BooleanValue.Name

	} else if sv := tt.StringValue(); sv != nil {
		res.StringValue = NewStringValueNode(ctx, sv)
		res.Name = "COL_" + res.StringValue.Name

	} else if qn := tt.QualifiedName(); qn != nil {
		res.FuncCall = NewFuncCallNode(ctx, qn.GetText(), tt.AllExpression())
		res.Name = "COL_" + qn.GetText()

	} else if be := tt.GetBase(); be != nil {
		res.Base = NewPrimaryExpressionNode(ctx, be)
		res.FieldName = NewIdentifierNode(ctx, tt.GetFieldName())
		res.Name = res.Base.Name + "." + res.FieldName.GetText()

	} else if id := tt.Identifier(); id != nil {
		res.Identifier = NewIdentifierNode(ctx, id)
		res.Name = res.Identifier.GetText()

	} else {
		res.ParenthesizedExpression = NewExpressionNode(ctx, children[1].(parser.IExpressionContext))
		res.Name = res.ParenthesizedExpression.Name
	}

	return res
}

func (self *PrimaryExpressionNode) Result(input *DataSource.DataSource) interface{} {
	if self.Number != nil {
		return self.Number.Result(input)

	} else if self.BooleanValue != nil {
		return self.BooleanValue.Result(input)

	} else if self.StringValue != nil {
		return self.StringValue.Result(input)

	} else if self.Base != nil {
		return input.GetValsByName(self.Name)[0]

	} else if self.Identifier != nil {
		return self.Identifier.Result(input)

	} else if self.ParenthesizedExpression != nil {
		return self.ParenthesizedExpression.Result(input)

	} else if self.FuncCall != nil {
		return self.FuncCall.Result(input)

	} else if self.Base != nil {
		name := fmt.Sprintf("%v", self.Base.Result(input)) + "." + self.FieldName.GetText()
		return input.GetValsByName(name)[0]
	}
	return nil
}

func (self *PrimaryExpressionNode) IsAggregate() bool {
	if self.Number != nil {
		return false

	} else if self.BooleanValue != nil {
		return false

	} else if self.StringValue != nil {
		return false

	} else if self.Identifier != nil {
		return false

	} else if self.ParenthesizedExpression != nil {
		return self.ParenthesizedExpression.IsAggregate()

	} else if self.FuncCall != nil {
		return self.FuncCall.IsAggregate()
	}
	return false
}

/////////////////////////////
type FuncCallNode struct {
	FuncName    string
	Expressions []*ExpressionNode
}

func NewFuncCallNode(ctx *Context.Context, name string, expressions []parser.IExpressionContext) *FuncCallNode {
	res := &FuncCallNode{
		FuncName:    name,
		Expressions: make([]*ExpressionNode, len(expressions)),
	}
	for i := 0; i < len(expressions); i++ {
		res.Expressions[i] = NewExpressionNode(ctx, expressions[i])
	}
	return res
}

func (self *FuncCallNode) Result(input *DataSource.DataSource) interface{} {
	switch self.FuncName {
	case "SUM":
		return SUM(input, self.Expressions[0])
	case "MIN":
		return MIN(input, self.Expressions[0])
	case "MAX":
		return MAX(input, self.Expressions[0])
	case "ABS":
		return ABS(input)
	}
	return nil
}

func (self *FuncCallNode) IsAggregate() bool {
	switch self.FuncName {
	case "SUM":
		return true
	case "MIN":
		return true
	case "MAX":
		return true
	case "ABS":
		return false
	}
	return false
}

func SUM(input *DataSource.DataSource, t *ExpressionNode) interface{} {
	var res interface{}
	for !input.IsEnd() {
		tmp := t.Result(input)
		if res == nil {
			res = tmp
		} else {
			res = Common.Arithmetic(res, tmp, Common.PLUS)
		}
		input.Next()
	}
	return res
}

func MIN(input *DataSource.DataSource, t *ExpressionNode) interface{} {
	var res interface{}
	for !input.IsEnd() {
		tmp := t.Result(input)
		if res == nil {
			res = tmp
		} else {
			if Common.Cmp(res, tmp) > 0 {
				res = tmp
			}
		}
		input.Next()
	}
	return res
}

func MAX(input *DataSource.DataSource, t *ExpressionNode) interface{} {
	var res interface{}
	for !input.IsEnd() {
		tmp := t.Result(input)
		if res == nil {
			res = tmp
		} else {
			if Common.Cmp(res, tmp) < 0 {
				res = tmp
			}
		}
		input.Next()
	}
	return res
}

func ABS(input *DataSource.DataSource) interface{} {
	return nil
}
