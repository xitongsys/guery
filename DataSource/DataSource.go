package DataSource

import ()

type DataSource interface {
	ReadRow() []interface{}
	ReadColumnByName(cols ...string) []interface{}
	ReadColumnByIndex(indexes ...int) []interface{}
	Size() int64
	Names() []string
	IsEnd() bool
}
