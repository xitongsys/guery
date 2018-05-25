package Util

import (
	"fmt"
	"io"
)

type PartitionInfo struct {
	Metadata  *Metadata
	Rows      []*Row
	Locations []string
}

func NewPartitionInfo(md *Metadata) *PartitionInfo {
	return &Partition{
		Metadata: md,
		Rows:     []*Row{},
	}
}

func (self *PartitionInfo) GetPartitionNum() int {
	return len(self.Rows)
}

func (self *PartitionInfo) GetPartition(i int) *RowsBuffer {
	if i >= len(self.Rows) {
		return nil
	}
	rowsBuffer := NewRowsBuffer(self.Metadata)
	rowsBuffer.Write(self.Rows)
	return rowsBuffer
}

func (self *PartitionInfo) GetLocation(i int) string {
	if i >= len(self.Rows) {
		return ""
	}
	return self.Locations[i]
}
