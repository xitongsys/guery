package Util

type ColumnType int32

const (
	_ ColumnType = iota
	BOOL
	INT
	DOUBLE
	STRING
)

type Metadata struct {
	Name        string
	ColumnNames []string
	ColumnTypes []ColumnType
	ColumnMap   map[string]int
}

func (self *Metadata) Reset() {
	self.ColumnMap = make(map[string]int)
	for i, name := range self.ColumnNames {
		self.ColumnMap[name] = i
		self.ColumnMap[self.Name+"."+name] = i
	}
}

func NewMetadata(name string, colNames []string, colTypes []ColumnType) *Metadata {
	res := &Metadata{
		Name:        name,
		ColumnNames: colNames,
		ColumnTypes: colTypes,
	}
	res.Reset()
	return res
}
