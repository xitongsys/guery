package DataSource

import ()

type DataSource interface {
	ReadRow() []interface{}
	ReadColumnByName(cols ...string) []interface{}
	ReadColumnByIndex(indexes ...int) []interface{}
	Next() error
	Size() int64
	GetColumnNames() []string
	GetName() string
	IsEnd() bool
	GetRow() DataSource
	Reset()
}
