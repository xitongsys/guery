package Plan

import (
	"fmt"

	"github.com/xitongsys/guery/Common"
	"github.com/xitongsys/guery/DataSource"
	"github.com/xitongsys/guery/parser"
)

type QuerySpecificationNode struct {
	Tree          *parser.QuerySpecificationContext
	SetQuantifier *Common.Quantifier
	SelectItems   []*SelectItemNode
	Relations     []*RelationNode
	Where         *BooleanExpressionNode
	GroupBy       *GroupByNode
	Having        *BooleanExpressionNode
}

func NewQuerySpecificationNode(ctx *Context, t *parser.QuerySpecificationContext) *QuerySpecificationNode {
	res := &QuerySpecificationNode{
		Tree:        t,
		SelectItems: make([]*SelectItemNode, 0),
		Relations:   make([]*RelationNode, 0),
	}
	children := t.GetChildren()
	for i := 0; i < len(children); i++ {
		child := children[i]
		switch child.(type) {
		case *parser.SelectItemContext:
			res.SelectItems = append(res.SelectItems,
				NewSelectItemNode(ctx, child.(*parser.SelectItemContext)))
		}
	}
	return res
}

func (self *QuerySpecificationNode) Result(ctx *Context) DataSource.DataSource {
	row := make([]interface{}, len(self.SelectItems))
	for i := 0; i < len(self.SelectItems); i++ {
		row[i] = self.SelectItems[i].Result(ctx)
		fmt.Println("====", row[i])
	}
	return nil
}
