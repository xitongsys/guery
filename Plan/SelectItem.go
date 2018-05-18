package Plan

import (
	"io"
	"strings"

	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/parser"
)

type SelectItemNode struct {
	Expression    *ExpressionNode
	QualifiedName *QualifiedNameNode
	Identifier    *IdentifierNode
	Names         []string
}

func NewSelectItemNode(t parser.ISelectItemContext) *SelectItemNode {
	res := &SelectItemNode{}
	tt := t.(*parser.SelectItemContext)
	if id := tt.Identifier(); id != nil {
		res.Identifier = NewIdentifierNode(id)
	}

	if ep := tt.Expression(); ep != nil {
		res.Expression = NewExpressionNode(ep)
		names := strings.Split(res.Expression.Name, ".")
		res.Names = []string{names[len(names)-1]}

	} else if qn := tt.QualifiedName(); qn != nil {
		res.QualifiedName = NewQulifiedNameNode(qn)
	}

	if res.Identifier != nil {
		res.Names = []string{tt.Identifier().(*parser.IdentifierContext).GetText()}
	}
	return res
}

func (self *SelectItemNode) GetNames() []string {
	return self.Names
}

func (self *SelectItemNode) GetNamesAndTypes(md *Util.Metadata) ([]string, []Util.Type, error) {
	types := []Util.Type{}
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
func (self *SelectItemNode) GetColumns(md *Util.Metadata) ([]string, error) {
	if self.Expression != nil {
		return self.Expression.GetColumns()
	} else {
		return md.GetColumnNames(), nil
	}
}

func (self *SelectItemNode) Result(input *Util.RowsBuffer) ([]interface{}, error) {
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
		res = append(res, row.Vals...)
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
