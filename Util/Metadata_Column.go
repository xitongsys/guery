package Util

import (
	"fmt"
)

type ColumnMetadata struct {
	Catalog    string
	Schema     string
	Table      string
	ColumnName string
	ColumnType Type
}

func NewColumnMetadata(t Type, metrics ...string) *ColumnMetadata {
	res := &ColumnMetadata{
		Catalog:    "DEFAULT",
		Schema:     "DEFAULT",
		Table:      "DEFAULT",
		ColumnName: "DEFAULT",
		ColumnType: t,
	}
	ln := len(metrics)
	if ln >= 1 {
		res.ColumnName = metrics[ln-1]
	}
	if ln >= 2 {
		res.Table = metrics[ln-2]
	}
	if ln >= 3 {
		res.Schema = metrics[ln-3]
	}
	if ln >= 4 {
		res.Schema = metrics[ln-4]
	}
	return res
}

func (self *ColumnMetadata) Copy() *ColumnMetadata {
	return &ColumnMetadata{
		Catalog:    self.Catalog,
		Schema:     self.Schema,
		Table:      self.Table,
		ColumnName: self.ColumnName,
	}
}

func (self *ColumnMetadata) GetName() string {
	return fmt.Sprintf("%v.%v.%v.%v", self.Catalog, self.Schema, self.Table, self.ColumnName)
}
