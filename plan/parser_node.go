package plan

import (
	"github.com/xitongsys/guery/gtype"
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/row"
)

type ParserNode interface {
	GetType(md *Metadata.Metadata) (Type.Type, error)
	GetColumns() ([]string, error)
	Result(input *Row.RowsGroup) (interface{}, error)
	IsAggregate() bool
}
