package DataSource

import (
	"strings"
)

type TableSource struct {
	Name      []string
	ColumnMap map[string]int
	Vals      [][]interface{}
	Index     int
	RowNum    int
}

func NewTableSource(name string, columnNames []string) *TableSource {
	res := &TableSource{
		Name:      []string{name},
		ColumnMap: make(map[string]int),
		Vals:      make([][]interface{}, len(columnNames)),
		Index:     0,
		RowNum:    0,
	}
	for i := 0; i < len(columnNames); i++ {
		res.ColumnMap[columnNames[i]] = i
		res.ColumnMap[name+"."+columnNames[i]] = i
	}
	return res
}

func CopyEmptyTableSource() *TableSource {
}

func (self *TableSource) Append(vals []string) {
	if len(vals) == len(self.Vals) {
		for i := 0; i < len(vals); i++ {
			self.Vals[i] = append(self.Vals[i], vals[i])
		}
		self.RowNum++
	}
}

func (self *TableSource) SelectColumns(cols ...string) DataSource {
	res := &TableSource{
		Name:      self.Name,
		ColumnMap: self.ColumnMap,
		Vals:      make([][]interface{}, len(self.Vals)),
	}
	for _, col := range cols {
		index := self.ColumnMap[col]
		res.Vals[index] = self.Vals[index]
	}
	return res
}

func (self *TableSource) First() DataSource {
	res := *self
	res.Index = 0
	return &res
}

func (self *TableSource) Next() DataSource {
	res := *self
	res.Index++
	return &res
}

func (self *TableSource) GetVals() []interface{} {
	res := make([]interface{}, len(self.Vals))
	if self.Index < self.RowNum {
		for i := 0; i < len(self.Vals); i++ {
			res[i] = self.Vals[i][self.Index]
		}
	}
	return res
}

func (self *TableSource) IsEnd() bool {
	return self.Index >= self.RowNum
}

func (self *TableSource) GetValsByName(cols ...string) []interface{} {
	res := make([]interface{}, len(cols))
	if self.Index < self.RowNum {
		for i := 0; i < len(cols); i++ {
			res[i] = self.Vals[self.ColumnMap[cols[i]]][self.Index]
		}
	}
	return res
}

func (self *TableSource) GetValsByIndex(indexes ...int) []interface{} {
	res := make([]interface{}, len(indexes))
	if self.Index < self.RowNum {
		for i := 0; i < len(indexes); i++ {
			res[i] = self.Vals[indexes[i]][self.Index]
		}
	}
	return res
}

func (self *TableSource) Alias(name string) {
	self.Name = append(self.Name, name)
	for key, val := range self.ColumnMap {
		keys := strings.Split(key, ".")
		if len(keys) == 1 {
			self.ColumnMap[name+"."+key] = val
		} else {
			keys[0] = name
			self.ColumnMap[strings.Join(keys, ".")] = val
		}
	}
}

func (self *TableSource) AliasColumn(colName string, index int) {
	for _, name := range self.Name {
		self.ColumnMap[name+"."+colName] = index
	}
}

func (self *TableSource) GetRowNum() int {
	return self.RowNum
}
