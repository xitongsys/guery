package Plan

import (
	"github.com/xitongsys/guery/DataSource"
	"github.com/xitongsys/guery/parser"
)

type QueryPrimaryNode struct {
	Tree               *parser.QueryPrimaryContext
	QuerySpecification *QuerySpecificationNode
	SubQuery           *QueryNode
}

func NewQueryPrimaryNode(ctx *Context, t *parser.QueryPrimaryContext) *QueryPrimaryNode {
	res := &QueryPrimaryNode{
		Tree: t,
	}
	child := t.GetChildren()[0]
	switch child.(type) {
	case *parser.QuerySpecificationContext:
		res.QuerySpecification = NewQuerySpecificationNode(ctx,
			child.(*parser.QuerySpecificationContext))
	}
	return res
}

func (self *QueryPrimaryNode) Result(ctx *Context) DataSource.DataSource {
	if self.QuerySpecification != nil {
		return self.QuerySpecification.Result(ctx)
	} else {
		return self.SubQuery.Result(ctx)
	}

}
