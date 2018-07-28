package Plan

import (
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Row"
	"github.com/xitongsys/guery/Type"
)

type ParserNode interface {
	GetType(md *Metadata.Metadata) (Type.Type, error)
	GetColumns() ([]string, error)
	Result(input *Row.RowsGroup) (interface{}, error)
	IsAggregate() bool
}
