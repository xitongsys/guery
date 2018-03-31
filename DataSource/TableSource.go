package DataSource

import (
	"fmt"
	"github.com/xitongsys/guery/Plan"
)

type TableSource struct {
	Name            string
	ColumnNames     []string
	ColumnTypes     []Plan.Type
	ColumnNameIndex map[string]int
	Vals            [][]interface{}
	Index           int64
}

func NewTableSource(name string, columnNames []string, columnTypes []Plan.Type) {
	res := &TableSource{
		Name:        name,
		ColumnNames: columnNames,
		ColumnTypes: columnTypes,
		Index:       0,
	}
	res.ColumnNameIndex = make(map[string]int)
	for i := 0; i < len(columnNames); i++ {
		res[columnNames[i]] = i
	}
	return res
}

func (self *TableSource) Append(vals []interface{}) {
	self.Vals = append(self.Vals, vals)
}

func (self *TableSource) Size() int64 {
	return len(self.Vals)
}

func (self *TableSource) ReadRow() []interface{} {
	if len(self.Vals) <= self.Index {
		return []interface{}{}
	}
	self.Index++
	return self.Vals[self.Index-1]
}

func (self *TableSource) ReadColumn(cols ...string) []interface{} {
	if len(self.Vals) <= self.Index {
		return []interface{}{}
	}
	res := make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {
		res[i] = self.Vals[self.Index][self.ColumnNameIndex[cols[i]]]
	}
	self.Index++
	return res
}

func (self *TableSource) Names() []string {
	return self.ColumnNames
}

func (self *TableSource) Types() []Plan.Type {
	return self.ColumnTypes
}
