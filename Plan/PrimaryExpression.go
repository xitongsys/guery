package Plan

import (
	"fmt"
	"io"

	"github.com/xitongsys/guery/Util"
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

func NewPrimaryExpressionNode(t parser.IPrimaryExpressionContext) *PrimaryExpressionNode {
	tt := t.(*parser.PrimaryExpressionContext)
	res := &PrimaryExpressionNode{}
	children := tt.GetChildren()
	if tt.NULL() != nil {
		res.Name = "NULL"

	} else if nu := tt.Number(); nu != nil {
		res.Number = NewNumberNode(nu)
		res.Name = "COL_" + res.Number.Name

	} else if bv := tt.BooleanValue(); bv != nil {
		res.BooleanValue = NewBooleanValueNode(bv)
		res.Name = "COL_" + res.BooleanValue.Name

	} else if sv := tt.StringValue(); sv != nil {
		res.StringValue = NewStringValueNode(sv)
		res.Name = "COL_" + res.StringValue.Name

	} else if qn := tt.QualifiedName(); qn != nil {
		res.FuncCall = NewFuncCallNode(qn.GetText(), tt.AllExpression())
		res.Name = "COL_" + qn.GetText()

	} else if be := tt.GetBase(); be != nil {
		res.Base = NewPrimaryExpressionNode(be)
		res.FieldName = NewIdentifierNode(tt.GetFieldName())
		res.Name = res.Base.Name + "." + res.FieldName.GetText()

	} else if id := tt.Identifier(); id != nil {
		res.Identifier = NewIdentifierNode(id)
		res.Name = res.Identifier.GetText()

	} else if tt.CASE() != nil {
		res.Case = NewCaseNode(tt.AllWhenClause(), tt.GetElseExpression())
		res.Name = "CASE"

	} else {
		res.ParenthesizedExpression = NewExpressionNode(children[1].(parser.IExpressionContext))
		res.Name = res.ParenthesizedExpression.Name
	}

	return res
}

func (self *PrimaryExpressionNode) GetType(md *Util.Metadata) (Util.Type, error) {
	if self.Number != nil {
		return self.Number.GetType(md)

	} else if self.BooleanValue != nil {
		return self.BooleanValue.GetType(md)

	} else if self.StringValue != nil {
		return self.StringValue.GetType(md)

	} else if self.Identifier != nil {
		return self.Identifier.GetType(md)

	} else if self.ParenthesizedExpression != nil {
		return self.ParenthesizedExpression.GetType(md)

	} else if self.FuncCall != nil {
		return self.FuncCall.GetType(md)

	} else if self.Case != nil {
		return self.Case.GetType(md)

	} else if self.Base != nil {
		return md.GetTypeByName(self.Name)
	}
	return Util.UNKNOWNTYPE, fmt.Errorf("wrong PrimaryExpressionNode")
}

func (self *PrimaryExpressionNode) Result(input *Util.RowsBuffer) (interface{}, error) {
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

	} else if self.Case != nil {
		return self.Case.Result(input)

	} else if self.Base != nil {
		row, err := input.Read()
		if err == io.EOF {
			return nil, nil
		}
		if err != nil {
			return nil, err
		}
		index := input.GetIndex(self.Name)

		if index < 0 || index > len(row.Vals) {
			return nil, fmt.Errorf("index out of range")
		}
		return row.Vals[index], nil
	}
	return nil, fmt.Errorf("wrong PrimaryExpressionNode")
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
