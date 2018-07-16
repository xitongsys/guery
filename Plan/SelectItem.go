package Plan

import (
	"io"

	"github.com/xitongsys/guery/Config"
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Row"
	"github.com/xitongsys/guery/Type"
	"github.com/xitongsys/guery/parser"
)

type SelectItemNode struct {
	Expression    *ExpressionNode
	QualifiedName *QualifiedNameNode
	Identifier    *IdentifierNode
	Names         []string
}

func NewSelectItemNode(runtime *Config.ConfigRuntime, t parser.ISelectItemContext) *SelectItemNode {
	res := &SelectItemNode{}
	tt := t.(*parser.SelectItemContext)
	if id := tt.Identifier(); id != nil {
		res.Identifier = NewIdentifierNode(runtime, id)
	}

	if ep := tt.Expression(); ep != nil {
		res.Expression = NewExpressionNode(runtime, ep)
		res.Names = []string{res.Expression.Name}

	} else if qn := tt.QualifiedName(); qn != nil {
		res.QualifiedName = NewQulifiedNameNode(runtime, qn)
	}

	if res.Identifier != nil {
		res.Names = []string{tt.Identifier().(*parser.IdentifierContext).GetText()}
	}
	return res
}

func (self *SelectItemNode) GetNames() []string {
	return self.Names
}

func (self *SelectItemNode) GetNamesAndTypes(md *Metadata.Metadata) ([]string, []Type.Type, error) {
	types := []Type.Type{}
	if self.Expression != nil {
		t, err := self.Expression.GetType(md)
		if err != nil {
			return self.Names, types, err
		}
		types = append(types, t)
		return self.Names, types, nil

	} else {
		return md.GetColumnNames(), md.GetColumnTypes(), nil
	}
}

//get the columns needed in SelectItem
func (self *SelectItemNode) GetColumns(md *Metadata.Metadata) ([]string, error) {
	if self.Expression != nil {
		return self.Expression.GetColumns()
	} else { //*
		return md.GetColumnNames(), nil
	}
}

func (self *SelectItemNode) Init(md *Metadata.Metadata) error {
	if self.Expression != nil { //some items
		if err := self.Expression.Init(md); err != nil {
			return err
		}

	}
	return nil
}

func (self *SelectItemNode) ExtractAggFunc(res *[]*FuncCallNode) {
	if self.Expression != nil { //some items
		self.Expression.ExtractAggFunc(res)
	} else { //*
	}
}

func (self *SelectItemNode) Result(input *Row.RowsGroup) ([]interface{}, error) {
	res := []interface{}{}
	if self.Expression != nil { //some items
		rec, err := self.Expression.Result(input)
		if err != nil {
			return res, err
		}
		res = append(res, rec)

	} else { //*
		row, err := input.Read()
		if err == io.EOF {
			return res, nil
		}
		if err != nil {
			return res, err
		}
		res = append(res, row.Vals)
		self.Names = input.Metadata.GetColumnNames()
	}

	return res, nil
}

func (self *SelectItemNode) IsAggregate() bool {
	if self.Expression != nil {
		return self.Expression.IsAggregate()
	}
	return false
}
