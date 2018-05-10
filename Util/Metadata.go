package Util

import (
	"strings"
)

type Metadata struct {
	Catalog     string
	Schema      string
	Table       string
	ColumnNames []string
	ColumnTypes []Type
	ColumnMap   map[string]int
}

func (self *Metadata) Reset() {
	self.ColumnMap = make(map[string]int)
	for i, name := range self.ColumnNames {
		self.ColumnMap[name] = i
		self.ColumnMap[self.Table+"."+name] = i
	}
}

func (self *Metadata) Copy(md *Metadata) {
	self.Catalog, self.Schema, self.Table = md.Catalog, md.Schema, md.Table
	self.ColumnNames = append(self.ColumnNames, md.ColumnNames...)
	self.ColumnTypes = append(self.ColumnTypes, md.ColumnTypes...)
	self.Reset()
}

func NewMetadata(catalog, schema, table string, colNames []string, colTypes []Type) *Metadata {
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
	catalog, schema, table = "TEST", "DEFAULT", "DEFAULT"
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
