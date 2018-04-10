package DataSource

import ()

type DataSource interface {
	ReadRow() []interface{}
	ReadColumnByName(cols ...string) []interface{}
	ReadColumnByIndex(indexes ...int) []interface{}
	Next() error
	Size() int64
	Names() []string
	IsEnd() bool
}
