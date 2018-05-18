package Util

import (
	"fmt"
	"strings"
)

type Metadata struct {
	Columns   []*ColumnMetadata
	ColumnMap map[string]int
}

func (self *Metadata) Reset() {
	self.ColumnMap = map[string]int{}
	for i, col := range self.Columns {
		name := col.ColumnName
		self.ColumnMap[name] = i

		name = col.Table + "." + name
		self.ColumnMap[name] = i

		name = col.Schema + "." + name
		self.ColumnMap[name] = i

		name = col.Catalog + "." + name
		self.ColumnMap[name] = i
	}
}

func (self *Metadata) Copy() *Metadata {
	res := NewMetadata()
	for _, c := range self.Columns {
		res = append(res, c.Copy())
	}
	res.Reset()
	return res
}

func (self *Metadata) Rename(name string) {
	for _, c := range self.Columns {
		c.Table = name
	}
	self.Reset()
}

func (self *Metadata) GetTypeByIndex(index int) (Type, error) {
	if index >= len(self.ColumnTypes) {
		return UNKNOWNTYPE, fmt.Errorf("index out of range")
	}
	return self.Columns[index].ColumnType, nil
}

func (self *Metadata) GetTypeByName(name string) (Type, error) {
	index, ok := self.ColumnMap[name]
	if !ok {
		return UNKNOWNTYPE, fmt.Errorf("unknown column name")
	}
	return self.GetTypeByIndex(index)
}

func (self *Metadata) AppendColumn(column *ColumnMetadata) {
	self.Columns = append(self.Columns, column)
	self.Reset()
}

func (self *Metadata) DeleteColumnByIndex(index int) {
	ln := len(self.Columns)
	if index < 0 || index >= ln {
		return
	}
	self.Columns = append(self.Columns[:index], self.Columns[index+1:]...)
}

func NewMetadata() *Metadata {
	return &Metadata{
		Columns:   []*ColumnMetadata{},
		ColumnMap: map[string]int{},
	}
}

func SplitName(name string) (catalog, schema, table, column string) {
	catalog, schema, table, column = "DEFAULT", "DEFAULT", "DEFAULT", "DEFAULT"
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
	if ln >= 4 {
		column = names[ln-4]
	}
	return
}

func JoinMetadata(mdl, mdr *Metadata) *Metadata {
	res := NewMetadata()
	for _, c := range mdl.Columns {
		res.Columns = append(res.Columns, c.Copy())
	}
	for _, c := range mdr.Columns {
		res.Columns = append(res.Columns, c.Copy())
	}
	res.Reset()
	return res
}
