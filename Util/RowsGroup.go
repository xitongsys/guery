package Util

import (
	"fmt"
	"io"
)

type RowsGroup struct {
	Metadata *Metadata
	Keys     []interface{}
	Rows     []*Row
	Index    int
}

func NewRowsGroup(md *Metadata) *RowsGroup {
	return &RowsGroup{
		Metadata: md,
		Keys:     []interface{}{},
		Rows:     []*Row{},
		Index:    0,
	}
}

func (self *RowsGroup) Read() (*Row, error) {
	if self.Index >= len(self.Rows) {
		return nil, io.EOF
	}
	self.Index++
	return self.Rows[self.Index-1], nil
}

func (self *RowsGroup) Write(row *Row) {
	self.Rows = append(self.Rows, row)
	self.Keys = row.Keys
}

func (self *RowsGroup) Reset() {
	self.Index = 0
}

func (self *RowsGroup) GetIndex(name string) int {
	if i, ok := self.Metadata.ColumnMap[name]; ok {
		return i
	}
	return -1
}

func (self *RowsGroup) GetKeyString() string {
	res := ""
	for _, key := range self.Keys {
		res += fmt.Sprintf("%v", key)
	}
	return res
}

func (self *RowsGroup) GetRowsNum() int {
	return len(self.Rows)
}
