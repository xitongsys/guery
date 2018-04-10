package Plan

import (
	"github.com/xitongsys/guery/Common"
	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/DataSource"
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

func NewPrimaryExpressionNode(ctx *Context.Context, t parser.IPrimaryExpressionContext) *PrimaryExpressionNode {
	tt := t.(*parser.PrimaryExpressionContext)
	res := &PrimaryExpressionNode{}
	children := tt.GetChildren()
	if tt.NULL() != nil {
	} else if nu := tt.Number(); nu != nil {
		res.Number = NewNumberNode(ctx, nu)

	} else if bv := tt.BooleanValue(); bv != nil {
		res.BooleanValue = NewBooleanValueNode(ctx, bv)

	} else if sv := tt.StringValue(); sv != nil {
		res.StringValue = NewStringValueNode(ctx, sv)

	} else if id := tt.Identifier(); id != nil {
		res.Identifier = NewIdentifierNode(ctx, id)

	} else if qn := tt.QualifiedName(); qn != nil {
		res.FuncCall = NewFuncCallNode(ctx, qn.GetText(), tt.AllExpression())

	} else {
		res.ParenthesizedExpression = NewExpressionNode(ctx, children[1].(parser.IExpressionContext))
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

func (self *FuncCallNode) Result(input DataSource.DataSource) interface{} {
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

func SUM(input DataSource.DataSource, t *ExpressionNode) interface{} {
	var res interface{}
	for !input.IsEnd() {
		tmp := t.Result(input)
		if res == nil {
			res = tmp
		} else {
			res = Common.Arithmetic(res, tmp, Common.PLUS)
		}
	}
	return res
}

func MIN(input DataSource.DataSource, t *ExpressionNode) interface{} {
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
	}
	return res
}

func MAX(input DataSource.DataSource, t *ExpressionNode) interface{} {
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
	}
	return res
}

func ABS(input DataSource.DataSource) interface{} {
	return nil
}
