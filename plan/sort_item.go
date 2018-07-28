package plan

import (
	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/gtype"
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/parser"
	"github.com/xitongsys/guery/row"
)

type SortItemNode struct {
	Expression *ExpressionNode
	OrderType  gtype.OrderType
}

func NewSortItemNode(runtime *config.ConfigRuntime, t parser.ISortItemContext) *SortItemNode {
	tt := t.(*parser.SortItemContext)
	res := &SortItemNode{
		Expression: NewExpressionNode(runtime, tt.Expression()),
		OrderType:  gtype.ASC,
	}

	if ot := tt.GetOrdering(); ot != nil {
		if ot.GetText() != "ASC" {
			res.OrderType = gtype.DESC
		}
	}

	return res
}

func (self *SortItemNode) GetColumns() ([]string, error) {
	return self.Expression.GetColumns()
}

func (self *SortItemNode) Init(md *metadata.Metadata) error {
	return self.Expression.Init(md)
}

func (self *SortItemNode) Result(input *row.RowsGroup) (interface{}, error) {
	return self.Expression.Result(input)
}

func (self *SortItemNode) IsAggregate() bool {
	return self.Expression.IsAggregate()
}

func (self *SortItemNode) GetType(md *metadata.Metadata) (gtype.Type, error) {
	return self.Expression.GetType(md)
}
