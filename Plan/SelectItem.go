package Plan

import (
	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/DataSource"
	"github.com/xitongsys/guery/parser"
)

type SelectItemNode struct {
	Expression    *ExpressionNode
	QualifiedName *QualifiedNameNode
	Identifier    *IdentifierNode
	Star          bool
}

func NewSelectItemNode(ctx *Context.Context, t parser.ISelectItemContext) *SelectItemNode {
	res := &SelectItemNode{
		Star: false,
	}
	tt := t.(*parser.SelectItemContext)
	if id := tt.Identifier(); id != nil {
		res.Identifier = NewIdentifierNode(ctx, id)
	}

	if ep := tt.Expression(); ep != nil {
		res.Expression = NewExpressionNode(ctx, ep)
	} else if qn := tt.QualifiedName(); qn != nil {
		res.QualifiedName = NewQulifiedNameNode(ctx, qn)
		res.Star = true
	} else {
		res.Star = true
	}
	return res
}

func (self *SelectItemNode) Names() []string {
	if self.Star {

	} else {

	}

}

func (self *SelectItemNode) Result(intput DataSource.DataSource) []interface{} {
	res := []interface{}{}
	if self.Expression != nil {
		res = append(res, self.Expression.Result(input))

	} else {
		res = input.ReadRow()
	}

	return res
}
