package Plan

import (
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/parser"
)

type SelectItemNode struct {
	Expression    *ExpressionNode
	QualifiedName *QualifiedNameNode
	Identifier    *IdentifierNode
	Star          bool
	Names         []string
}

func NewSelectItemNode(t parser.ISelectItemContext) *SelectItemNode {
	res := &SelectItemNode{
		Star: false,
	}
	tt := t.(*parser.SelectItemContext)
	if id := tt.Identifier(); id != nil {
		res.Identifier = NewIdentifierNode(id)
	}

	if ep := tt.Expression(); ep != nil {
		res.Expression = NewExpressionNode(ep)
		res.Names = []string{res.Expression.Name}

	} else if qn := tt.QualifiedName(); qn != nil {
		res.QualifiedName = NewQulifiedNameNode(qn)
		res.Star = true
	} else {
		res.Star = true
	}

	if res.Identifier != nil {
		res.Names = []string{tt.Identifier().(*parser.IdentifierContext).GetText()}
	}
	return res
}

func (self *SelectItemNode) GetNames() []string {
	return self.Names
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
		self.Names = input.Metadata.ColumnNames
	}

	return res, nil
}

func (self *SelectItemNode) IsAggregate() bool {
	if self.Expression != nil {
		return self.Expression.IsAggregate()
	}
	return false
}
