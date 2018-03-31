package DataSource

import (
	"github.com/xitongsys/guery/Common"
)

type DataSource interface {
	ReadRow() []interface{}
	ReadColumn(cols ...string) []interface{}
	Size() int64
	Names() []string
	Types() []Common.Type
}
