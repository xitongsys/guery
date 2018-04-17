package DataSource

import (
	"strings"
)

type DataSource struct {
	Name          string
	ColumnNames   []string
	ColumnMap     map[string]int
	ColumnBuffers []ColumnBuffer
	Vals          []interface{}
	CurIndex      int
	RowNum        int
}

func NewDataSource(name string, columnNames []string, columnBuffers []ColumnBuffer) *DataSource {
	res := &DataSource{
		Name:          name,
		ColumnNames:   columnNames,
		ColumnMap:     make(map[string]int),
		ColumnBuffers: columnBuffers,
		Vals:          make([]interface{}, len(columnBuffers)),
		CurIndex:      -1,
		RowNum:        0,
	}

	for i := 0; i < len(columnNames); i++ {
		res.ColumnMap[columnNames[i]] = i
		res.ColumnMap[name+"."+columnNames[i]] = i
	}

	for _, buf := range columnBuffers {
		if res.RowNum < buf.Size() {
			res.RowNum = buf.Size()
		}
	}
	return res
}

func MergeDataSource(leftDs, rightDs *DataSource) *DataSource {
	name := leftDs.Name + rightDs.Name
	columnNames := []string{}
	columnNames = append(columnNames, leftDs.ColumnNames...)
	columnNames = append(columnNames, rightDs.ColumnNames...)
	columnBuffers := []ColumnBuffer{}
	columnBuffers = append(columnBuffers, leftDs.ColumnBuffers...)
	columnBuffers = append(columnBuffers, rightDs.ColumnBuffers...)
	res := NewDataSource(name, columnNames, columnBuffers)

	for k, v := range leftDs.ColumnMap {
		res.ColumnMap[k] = v
	}
	for k, v := range rightDs.ColumnMap {
		res.ColumnMap[k] = v
	}
	return res
}

func (self *DataSource) Append(ds *DataSource) {
	for !ds.IsEnd() {
		vals := ds.GetRowVals()
		if len(vals) != len(self.ColumnBuffers) {
			return
		}

		for i := 0; i < len(self.ColumnBuffers); i++ {
			self.ColumnBuffers[i].(*MemColumnBuffer).Append(vals[i])
		}
		self.RowNum++
		ds.Next()
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
		Name:          self.Name,
		ColumnNames:   self.ColumnNames,
		ColumnMap:     self.ColumnMap,
		ColumnBuffers: make([]ColumnBuffer, len(self.ColumnBuffers)),
		Vals:          []interface{}{},
		CurIndex:      -1,
		RowNum:        1,
	}
	if self.CurIndex < 0 {
		self.Next()
	}

	for i := 0; i < len(self.ColumnBuffers); i++ {
		res.ColumnBuffers[i] = NewMemColumnBuffer()
		res.ColumnBuffers[i].(*MemColumnBuffer).Append(self.Vals[i])
	}
	return res
}

func (self *DataSource) SelectColumns(cols ...string) *DataSource {
	res := &DataSource{
		Name:          self.Name,
		ColumnNames:   self.ColumnNames,
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

func (self *DataSource) GetRowVals() []interface{} {
	if self.CurIndex < 0 {
		self.Next()
	}
	return self.Vals
}

func (self *DataSource) Next() {
	if self.CurIndex < self.RowNum {
		self.Vals = make([]interface{}, len(self.ColumnBuffers))
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
	if self.CurIndex < 0 {
		self.Next()
	}
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
	for i, colName := range self.ColumnNames {
		oldName := self.Name + "." + colName
		if _, ok := self.ColumnMap[oldName]; ok {
			delete(self.ColumnMap, oldName)
		}
		if _, ok := self.ColumnMap[colName]; ok {
			delete(self.ColumnMap, oldName)
		}

		names := strings.Split(colName, ".")
		newColName := names[len(names)-1]
		self.ColumnMap[newColName] = i
		self.ColumnMap[name+"."+newColName] = i
	}
	self.Name = name
}

func (self *DataSource) GetRowNum() int {
	return self.RowNum
}

func (self *DataSource) GetColumnNum() int {
	return len(self.ColumnBuffers)
}
