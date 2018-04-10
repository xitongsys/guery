package DataSource

import ()

type TableSource struct {
	Name            string
	ColumnNames     []string
	ColumnNameIndex map[string]int
	Vals            [][]interface{}
	Index           int64
}

func NewTableSource(name string, columnNames []string) *TableSource {
	res := &TableSource{
		Name:        name,
		ColumnNames: columnNames,
		Index:       0,
	}
	res.ColumnNameIndex = make(map[string]int)
	for i := 0; i < len(columnNames); i++ {
		res.ColumnNameIndex[columnNames[i]] = i
	}
	return res
}

func (self *TableSource) IsEnd() bool {
	return self.Index >= int64(len(self.Vals))
}

func (self *TableSource) Append(vals []interface{}) {
	self.Vals = append(self.Vals, vals)
}

func (self *TableSource) Size() int64 {
	return int64(len(self.Vals))
}

func (self *TableSource) ReadRow() []interface{} {
	if int64(len(self.Vals)) <= self.Index {
		return make([]interface{}, len(self.ColumnNames))
	}
	self.Index++
	return self.Vals[self.Index-1]
}

func (self *TableSource) ReadColumnByName(cols ...string) []interface{} {
	if int64(len(self.Vals)) <= self.Index {
		return []interface{}{}
	}
	res := make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {
		res[i] = self.Vals[self.Index][self.ColumnNameIndex[cols[i]]]
	}
	self.Index++
	return res
}

func (self *TableSource) ReadColumnByIndex(indexes ...int) []interface{} {
	if int64(len(self.Vals)) <= self.Index {
		return []interface{}{}
	}
	res := make([]interface{}, len(indexes))
	for i := 0; i < len(indexes); i++ {
		res[i] = self.Vals[self.Index][indexes[i]]
	}
	self.Index++
	return res
}

func (self *TableSource) Names() []string {
	return self.ColumnNames
}
