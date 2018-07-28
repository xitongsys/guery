package plan

import (
	"github.com/xitongsys/guery/gtype"
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/row"
)

type ParserNode interface {
	GetType(md *metadata.Metadata) (gtype.Type, error)
	GetColumns() ([]string, error)
	Result(input *row.RowsGroup) (interface{}, error)
	IsAggregate() bool
}
