package DataSource

import (
	"github.com/xitongsys/guery/Common"
)

type DataSource interface {
	ReadRow() []interface{}
	ReadColumnByName(cols ...string) []interface{}
	ReadColumnByIndex(indexes ...int) []interface{}
	Size() int64
	Names() []string
	Types() []Common.Type
}
