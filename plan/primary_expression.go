package plan

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/gtype"
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/parser"
	"github.com/xitongsys/guery/row"
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
	Index     int //for performance

	//case when
	Case *CaseNode
}

func NewPrimaryExpressionNode(runtime *config.ConfigRuntime, t parser.IPrimaryExpressionContext) *PrimaryExpressionNode {
	tt := t.(*parser.PrimaryExpressionContext)
	res := &PrimaryExpressionNode{}
	children := tt.GetChildren()
	if tt.NULL() != nil {
		res.Name = "NULL"

	} else if tt.Identifier() != nil && tt.StringValue() != nil {
		res.Identifier = NewIdentifierNode(runtime, tt.Identifier())
		res.StringValue = NewStringValueNode(runtime, tt.StringValue())
		res.Name = "COL_" + res.Identifier.GetText()

	} else if nu := tt.Number(); nu != nil {
		res.Number = NewNumberNode(runtime, nu)
		res.Name = "COL_" + res.Number.Name

	} else if bv := tt.BooleanValue(); bv != nil {
		res.BooleanValue = NewBooleanValueNode(runtime, bv)
		res.Name = "COL_" + res.BooleanValue.Name

	} else if sv := tt.StringValue(); sv != nil {
		res.StringValue = NewStringValueNode(runtime, sv)
		res.Name = "COL_" + res.StringValue.Name

	} else if qn := tt.QualifiedName(); qn != nil {
		res.FuncCall = NewFuncCallNode(runtime, qn.GetText(), tt.AllExpression())
		res.Name = "COL_" + qn.GetText()

	} else if be := tt.GetBase(); be != nil {
		res.Base = NewPrimaryExpressionNode(runtime, be)
		res.FieldName = NewIdentifierNode(runtime, tt.GetFieldName())
		res.Name = res.Base.Name + "." + res.FieldName.GetText()

	} else if id := tt.Identifier(); id != nil {
		res.Identifier = NewIdentifierNode(runtime, id)
		res.Name = res.Identifier.GetText()

	} else if tt.CASE() != nil {
		res.Case = NewCaseNode(runtime, tt.AllWhenClause(), tt.GetElseExpression())
		res.Name = "CASE"

	} else {
		res.ParenthesizedExpression = NewExpressionNode(runtime, children[1].(parser.IExpressionContext))
		res.Name = res.ParenthesizedExpression.Name
	}

	return res
}

func (self *PrimaryExpressionNode) ExtractAggFunc(res *[]*FuncCallNode) {
	if self.FuncCall != nil && self.FuncCall.IsAggregate() {
		colName := fmt.Sprintf("AGG_%v_%v", len(*res), rand.Int())
		globalFunc := self.FuncCall
		globalFunc.ResColName = colName
		*res = append(*res, globalFunc)

		self.FuncCall = nil
		self.Identifier = &IdentifierNode{
			Str: &colName,
		}

	} else if self.Case != nil {
		self.Case.ExtractAggFunc(res)

	} else if self.ParenthesizedExpression != nil {
		self.ParenthesizedExpression.ExtractAggFunc(res)
	}

}

func (self *PrimaryExpressionNode) GetType(md *metadata.Metadata) (gtype.Type, error) {
	if self.Number != nil {
		return self.Number.GetType(md)

	} else if self.Identifier != nil && self.StringValue != nil {
		if self.Identifier.NonReserved == nil {
			return gtype.UNKNOWNTYPE, fmt.Errorf("GetType: wrong PrimaryExpressionNode")
		}
		t := strings.ToUpper(*self.Identifier.NonReserved)
		switch t {
		case "TIMESTAMP":
			return gtype.TIMESTAMP, nil
		}

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
	return gtype.UNKNOWNTYPE, fmt.Errorf("GetType: wrong PrimaryExpressionNode")
}

func (self *PrimaryExpressionNode) GetColumns() ([]string, error) {
	res := []string{}
	if self.Number != nil {
		return res, nil

	} else if self.Identifier != nil && self.StringValue != nil {
		return res, nil

	} else if self.BooleanValue != nil {
		return res, nil

	} else if self.StringValue != nil {
		return res, nil

	} else if self.Identifier != nil {
		return self.Identifier.GetColumns()

	} else if self.ParenthesizedExpression != nil {
		return self.ParenthesizedExpression.GetColumns()

	} else if self.FuncCall != nil {
		return self.FuncCall.GetColumns()

	} else if self.Case != nil {
		return self.Case.GetColumns()

	} else if self.Base != nil {
		return []string{self.Name}, nil
	}
	return res, fmt.Errorf("GetColumns: wrong PrimaryExpressionNode")
}

func (self *PrimaryExpressionNode) Init(md *metadata.Metadata) error {
	if self.Number != nil {
		return self.Number.Init(md)

	} else if self.BooleanValue != nil {
		return self.BooleanValue.Init(md)

	} else if self.StringValue != nil {
		return self.StringValue.Init(md)

	} else if self.Identifier != nil {
		return self.Identifier.Init(md)

	} else if self.ParenthesizedExpression != nil {
		return self.ParenthesizedExpression.Init(md)

	} else if self.FuncCall != nil {
		return self.FuncCall.Init(md)

	} else if self.Case != nil {
		return self.Case.Init(md)

	} else if self.Base != nil {
		index, err := md.GetIndexByName(self.Name)
		if err != nil {
			return err
		}
		self.Index = index
		return nil
	}
	return fmt.Errorf("Result: wrong PrimaryExpressionNode")
}

func (self *PrimaryExpressionNode) Result(input *row.RowsGroup) (interface{}, error) {
	if self.Number != nil {
		return self.Number.Result(input)

	} else if self.Identifier != nil && self.StringValue != nil {
		t := *self.Identifier.NonReserved
		switch t {
		case "TIMESTAMP":
			resi, err := self.StringValue.Result(input)
			if err != nil {
				return nil, err
			}
			res := resi.([]interface{})
			for i := 0; i < len(res); i++ {
				res[i] = gtype.ToTimestamp(res[i])
			}
			return res, nil
		}

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
		rn := input.GetRowsNumber()
		res := make([]interface{}, rn)
		index := self.Index
		for i := 0; i < rn; i++ {
			res[i] = input.Vals[index][i]
		}
		return res, nil
	}
	return nil, fmt.Errorf("Result: wrong PrimaryExpressionNode")
}

func (self *PrimaryExpressionNode) IsAggregate() bool {
	if self.Number != nil {
		return false

	} else if self.Identifier != nil && self.StringValue != nil {
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
