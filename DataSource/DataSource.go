package DataSource

import ()

type DataSource interface {
	SelectColumns(cols ...string) DataSource
	First() DataSource
	Next() DataSource
	GetVals() []interface{}
	IsEnd() bool
	GetValsByName(cols ...string) []interface{}
	GetValsByIndex(cols ...int) []interface{}
	Alias(name string)
	AliasColumn(colName string, index int)
	GetRowNum() int
}
