package Plan

import (
	"fmt"

	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/DataSource"
	"github.com/xitongsys/guery/parser"
)

type PrimaryExpressionNode struct {
	//	Null         *NullNode
	Name string

	Number       *NumberNode
	BooleanValue *BooleanValueNode
	StringValue  *StringValueNode
	Identifier   *IdentifierNode

	//Func
	FuncCall *FuncCallNode
	//
	ParenthesizedExpression *ExpressionNode

	//t.name
	Base      *PrimaryExpressionNode
	FieldName *IdentifierNode

	//case when
	Case *CaseNode
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
	} else if tt.CASE() != nil {
		res.Case = NewCaseNode(ctx, tt.AllWhenClause(), tt.GetElseExpression())
		res.Name = "CASE"

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

	} else if self.Case != nil {
		return self.Case.Result(input)

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

	} else if self.Case != nil {
		return self.Case.IsAggregate()

	} else if self.FuncCall != nil {
		return self.FuncCall.IsAggregate()
	}
	return false
}
