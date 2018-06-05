package Plan

import (
	"github.com/xitongsys/guery/Util"
)

type ParserNode interface {
	GetType(md *Util.Metadata) (Util.Type, error)
	GetColumns() ([]string, error)
	Result(input *Util.RowsGroup) (interface{}, error)
	IsAggregate() bool
}
