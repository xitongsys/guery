package DataSource

import (
	"github.com/xitongsys/guery/Plan"
)

type DataSource interface {
	ReadRow() []interface{}
	ReadColumn(cols ...string) []interface{}
	Size() int64
	Names() []string
	Types() []Plan.Type
}
