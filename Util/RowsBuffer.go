package Util

import (
	"io"
)

type RowsBuffer struct {
	Metadata *Metadata
	Key      string
	Rows     []*Row
	Index    int
}

func NewRowsBuffer(md *Metadata) *RowsBuffer {
	return &RowsBuffer{
		Metadata: md,
		Key:      "",
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
	self.Key = row.Key
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
