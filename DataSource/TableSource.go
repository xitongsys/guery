package DataSource

import (
	"fmt"
)

type TableSchema interface {
	GetTypeByIndex(index int) (string, error)
	GetTypeByName(name string) (string, error)
	GetNameIndex(name string) (int, error)
	GetColumnNumber() int
	GetColumns() []string
}

type TableSchemaBase struct {
	types      []string
	typesIndex map[string]int
}

func (self *TableSchemaBase) GetTypeByIndex(index int) (string, error) {
	if index < 0 || index >= len(self.types) {
		return "", fmt.Errorf("index out of range")
	}
	return self.types[index], nil
}

func (self *TableSchemaBase) GetTypeByName(name string) (string, error) {
	if index, ok := self.typesIndex[name]; ok {
		return self.types[index], nil
	}
	return "", fmt.Errorf("name not found")
}

func (self *TableSchemaBase) GetNameIndex(name string) (int, error) {
	if index, ok := self.typesIndex[name]; ok {
		return index, nil
	}
	return -1, fmt.Errorf("name not found")
}

type TableSource interface {
	TableSchema
	ReadColumnByName(name string) (interface{}, error)
	ReadColumnByIndex(index int) (interface{}, error)
}

type TableSourceBase struct {
	TableSchemaBase
	data     [][]interface{}
	colIndex []int
}

func NewTableSourceBase() *TableSourceBase {
	res := new(TableSourceBase)
	return res
}

func (self *TableSourceBase) ReadColumnByName(name string) (interface{}, error) {
	if index, err := self.GetNameIndex(name); err != nil {
		return nil, err
	} else {
		ln := len(self.data[index])
		if self.colIndex[index] >= ln {
			return nil, fmt.Errorf("End")
		}
		res := self.data[index][self.colIndex[index]]
		self.colIndex[index]++
		return res, nil
	}
}

func (self *TableSourceBase) ReadColumnByIndex(index string) (interface{}, error) {
	lnc := len(self.data)
	if index < 0 || index >= lnc {
		return nil, fmt.Errorf("index out of range")
	}
	if ln := len(self.data[index]); colIndex[index] >= ln {
		return nil, fmt.Errorf("End")
	} else {
		res := self.data[index][self.colIndex[index]]
		self.colIndex[index]++
		return res, nil
	}
}
