package Util

import (
	"fmt"
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
		fullName := fmt.Sprintf("%s.%s.%s.%s", self.Catalog, self.Schema, self.Table, name)
		self.ColumnMap[fullName] = i
	}
}

func (self *Metadata) Copy(md *Metadata) {
	self.Catalog, self.Schema, self.Table = md.Catalog, md.Schema, md.Table
	self.ColumnNames = append(self.ColumnNames, md.ColumnNames...)
	self.ColumnTypes = append(self.ColumnTypes, md.ColumnTypes...)
	self.Reset()
}

func (self *Metadata) Rename(rname string) {
	self.Table = rname
	self.Reset()
}

func (self *Metadata) GetTypeByIndex(index int) (Type, error) {
	if index >= len(self.ColumnTypes) {
		return UNKNOWNTYPE, fmt.Errorf("index out of range")
	}
	return self.ColumnTypes[index], nil
}

func (self *Metadata) GetTypeByName(name string) (Type, error) {
	index, ok := self.ColumnMap[name]
	if !ok {
		return UNKNOWNTYPE, fmt.Errorf("unknown column name")
	}
	return self.GetTypeByIndex(index)
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
		Catalog:   "DEFAULT",
		Schema:    "DEFAULT",
		Table:     "DEFAULT",
		ColumnMap: make(map[string]int),
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

func JoinMetadata(mdl, mdr *Metadata) *Metadata {
	res := NewDefaultMetadata()
	res.ColumnNames = append(res.ColumnNames, mdl.ColumnNames...)
	res.ColumnNames = append(res.ColumnNames, mdr.ColumnNames...)
	res.ColumnTypes = append(res.ColumnTypes, mdl.ColumnTypes...)
	res.ColumnTypes = append(res.ColumnTypes, mdr.ColumnTypes...)

	for name, index := range mdl.ColumnMap {
		res.ColumnMap[name] = index
	}
	for name, index := range mdr.ColumnMap {
		res.ColumnMap[name] = index + len(mdl.ColumnNames)
	}
	//bres.Reset()
	return res
}
