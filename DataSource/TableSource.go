package DataSource

import ()

type TableSource struct {
	Name            []string
	ColumnNames     []string
	ColumnNameIndex map[string]int
	Vals            [][]interface{}
	Index           int64
}

func NewTableSource(name string, columnNames []string) *TableSource {
	res := &TableSource{
		Name:        []string{name},
		ColumnNames: columnNames,
		Index:       0,
	}
	res.ColumnNameIndex = make(map[string]int)
	for i := 0; i < len(columnNames); i++ {
		res.ColumnNameIndex[columnNames[i]] = i
		if name != "" {
			res.ColumnNameIndex[name+"."+columnNames[i]] = i
		}
	}
	return res
}

func NewTableSourceFromDataSource(ds DataSource) *TableSource {
	res := &TableSource{
		Name:        ds.GetName(),
		ColumnNames: ds.GetColumnNames(),
		Index:       0,
	}
	res.ColumnNameIndex = make(map[string]int)
	for i := 0; i < len(res.ColumnNames); i++ {
		res.ColumnNameIndex[res.ColumnNames[i]] = i
	}
	for i := int64(0); i < ds.Size(); i++ {
		res.Append(ds.ReadRow())
		ds.Next()
	}
	return res
}

func (self *TableSource) GetRow() DataSource {
	res := *self
	res.Index = 0
	res.Vals = res.Vals[self.Index : self.Index+1]
	return &res
}

func (self *TableSource) Reset() {
	self.Index = 0
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
	return self.Vals[self.Index]
}

func (self *TableSource) ReadColumnByName(cols ...string) []interface{} {
	if int64(len(self.Vals)) <= self.Index {
		return []interface{}{}
	}
	res := make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {
		res[i] = self.Vals[self.Index][self.ColumnNameIndex[cols[i]]]
	}
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
	return res
}

func (self *TableSource) Next() error {
	if int64(len(self.Vals)) <= self.Index {
		return nil
	}
	self.Index++
	return nil
}

func (self *TableSource) GetColumnNames() []string {
	return self.ColumnNames
}

func (self *TableSource) GetName() []string {
	return self.Name
}

func (self *TableSource) AddName(name string) {
	if name != "" {
		self.Name = append(self.Name, name)
		for i := 0; i < len(self.ColumnNames); i++ {
			self.ColumnNameIndex[name+"."+self.ColumnNames[i]] = i
		}
	}
}

func (self *TableSource) AddColumnName(colName string, index int) {
	if colName != "" {
		for _, name := range self.Name {
			self.ColumnNameIndex[name+"."+colName] = index
		}
	}
}
