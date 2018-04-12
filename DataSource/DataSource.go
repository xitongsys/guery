package DataSource

import (
	"strings"
)

type DataSource struct {
	Names         []string
	ColumnMap     map[string]int
	ColumnBuffers []ColumnBuffer
	Vals          []interface{}
	CurIndex      int
	RowNum        int
}

func NewDataSource(names []string, columnNames []string, columnBuffers []ColumnBuffer) *DataSource {
	res := &DataSource{
		Names:         names,
		ColumnMap:     make(map[string]int),
		ColumnBuffers: columnBuffers,
		Vals:          []interface{}{},
		CurIndex:      -1,
		RowNum:        0,
	}
	for _, name := range names {
		for i := 0; i < len(columnNames); i++ {
			res.ColumnMap[columnNames[i]] = i
			res.ColumnMap[name+"."+columnNames[i]] = i
		}
	}
	return res
}

func (self *DataSource) Merge(ds *DataSource) {
	for !ds.IsEnd() {
		vals := ds.GetRawVals()
		if len(vals) != len(self.ColumnBuffers) {
			return
		}

		for i := 0; i < len(self.ColumnBuffers); i++ {
			self.ColumnBuffers[i].(*MemColumnBuffer).Append(vals[i])
		}
	}
}

func (self *DataSource) Reset() {
	self.Vals, self.CurIndex = []interface{}{}, -1
	for i := 0; i < len(self.ColumnBuffers); i++ {
		self.ColumnBuffers[i] = self.ColumnBuffers[i].Duplicate()
	}
}

func (self *DataSource) SelectRow() *DataSource {
	res := &DataSource{
		Names:         self.Names,
		ColumnMap:     self.ColumnMap,
		ColumnBuffers: []ColumnBuffer{},
		Vals:          []interface{}{},
		CurIndex:      -1,
		RowNum:        1,
	}

	for i := 0; i < len(self.ColumnBuffers); i++ {
		res.ColumnBuffers[i] = NewMemColumnBuffer()
		res.ColumnBuffers[i].(*MemColumnBuffer).Append([]interface{}{self.Vals[i]})
	}
	return res
}

func (self *DataSource) SelectColumns(cols ...string) *DataSource {
	res := &DataSource{
		Names:         self.Names,
		ColumnMap:     self.ColumnMap,
		ColumnBuffers: make([]ColumnBuffer, len(self.Vals)),
		CurIndex:      self.CurIndex,
		RowNum:        self.RowNum,
	}
	for _, col := range cols {
		index := self.ColumnMap[col]
		res.Vals[index] = self.Vals[index]
	}
	return res
}

func (self *DataSource) GetRawVals() []interface{} {
	if self.CurIndex < 0 {
		self.Next()
	}
	return self.Vals
}

func (self *DataSource) Next() {
	if self.CurIndex+1 < self.RowNum {
		for i := 0; i < len(self.ColumnBuffers); i++ {
			self.Vals[i] = self.ColumnBuffers[i].Read()
		}
		self.CurIndex++
	}
}

func (self *DataSource) IsEnd() bool {
	return self.CurIndex >= self.RowNum
}

func (self *DataSource) GetValsByName(cols ...string) []interface{} {
	res := make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {
		res[i] = self.Vals[self.ColumnMap[cols[i]]]
	}
	return res
}

func (self *DataSource) GetValsByIndex(indexes ...int) []interface{} {
	res := make([]interface{}, len(indexes))
	for i := 0; i < len(indexes); i++ {
		res[i] = self.Vals[indexes[i]]
	}
	return res
}

func (self *DataSource) Alias(name string) {
	self.Names = append(self.Names, name)
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

func (self *DataSource) AliasColumn(colName string, index int) {
	for _, name := range self.Names {
		self.ColumnMap[name+"."+colName] = index
	}
}

func (self *DataSource) GetRowNum() int {
	return self.RowNum
}

func (self *DataSource) GetColumnNum() int {
	return len(self.ColumnBuffers)
}
