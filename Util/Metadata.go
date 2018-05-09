package Util

import (
	"strings"
)

type ColumnType int32

const (
	_ ColumnType = iota
	BOOL
	INT
	DOUBLE
	STRING
)

type Metadata struct {
	Catalog     string
	Schema      string
	Table       string
	ColumnNames []string
	ColumnTypes []ColumnType
	ColumnMap   map[string]int
}

func (self *Metadata) Reset() {
	self.ColumnMap = make(map[string]int)
	for i, name := range self.ColumnNames {
		self.ColumnMap[name] = i
		self.ColumnMap[self.Name+"."+name] = i
	}
}

func GetMetadata(catalog, schema, table string) *Metadata {
	return nil
}

func NewMetadata(catalog, schema, table string, colNames []string, colTypes []ColumnType) *Metadata {
	res := &Metadata{
		Catalog:     catalog,
		Schema:      schema,
		Table:       table,
		ColumnNames: colNames,
		ColumnTypes: colTypes,
	}
	res.Reset()
	return res
}

func NewDefaultMetadata() *Metadata {
	res := &Metadata{
		Catalog: "TEST",
		Schema:  "DEFAULT",
		Table:   "DEFAULT",
	}
	return res
}

func SplitName(name string) (catalog, schema, table string) {
	catalog, schema, table := "TEST", "DEFAULT", "DEFAULT"
	names := strings.Split(name, ".")
	ln := len(names)
	if ln >= 1 {
		table = names[ln-1]
	}
	if ln >= 2 {
		schema = names[ln-2]
	}
	if ln >= 3 {
		catalog = names[ln-3]
	}
	return
}
