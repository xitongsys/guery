package Util

import (
	"fmt"
	"io"
)

type RowsBuffer struct {
	Metadata *Metadata
	Keys     []interface{}
	Rows     []*Row
	Index    int
}

func NewRowsBuffer(md *Metadata) *RowsBuffer {
	return &RowsBuffer{
		Metadata: md,
		Keys:     []interface{}{},
		Rows:     []*Row{},
		Index:    0,
	}
}

func (self *RowsBuffer) Read() (*Row, error) {
	if self.Index >= len(self.Rows) {
		return nil, io.EOF
	}
	self.Index++
	return self.Rows[self.Index-1], nil
}

func (self *RowsBuffer) Write(row *Row) {
	self.Rows = append(self.Rows, row)
	self.Keys = row.Keys
}

func (self *RowsBuffer) Reset() {
	self.Index = 0
}

func (self *RowsBuffer) GetIndex(name string) int {
	if i, ok := self.Metadata.ColumnMap[name]; ok {
		return i
	}
	return -1
}

func (self *RowsBuffer) GetKeyString() string {
	res := ""
	for _, key := range self.Keys {
		res += fmt.Sprintf("%v", key)
	}
	return res
}
