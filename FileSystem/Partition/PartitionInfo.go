package Partition

import (
	"github.com/xitongsys/guery/FileSystem"
	"github.com/xitongsys/guery/Util"
)

type PartitionInfo struct {
	Metadata  *Util.Metadata
	Rows      []*Util.Row
	Locations []string
	FileTypes []FileSystem.FileType
	FileLists [][]*FileSystem.FileLocation

	//for no partition
	FileList []*FileSystem.FileLocation
}

func NewPartitionInfo(md *Util.Metadata) *PartitionInfo {
	return &PartitionInfo{
		Metadata:  md,
		Rows:      []*Util.Row{},
		Locations: []string{},
		FileTypes: []FileSystem.FileType{},
		FileLists: [][]*FileSystem.FileLocation{},

		FileList: []*FileSystem.FileLocation{},
	}
}

func (self *PartitionInfo) GetPartitionNum() int {
	return len(self.Rows)
}

func (self *PartitionInfo) GetPartition(i int) *Util.RowsGroup {
	if i >= len(self.Rows) {
		return nil
	}
	rb := Util.NewRowsGroup(self.Metadata)
	rb.Write(self.Rows[i])
	return rb
}

func (self *PartitionInfo) GetPartitionRow(i int) *Util.Row {
	if i >= len(self.Rows) {
		return nil
	}
	return self.Rows[i]
}

func (self *PartitionInfo) GetPartitionFiles(i int) []*FileSystem.FileLocation {
	if i >= len(self.FileLists) {
		return []*FileSystem.FileLocation{}
	}
	return self.FileLists[i]
}

func (self *PartitionInfo) GetNoPartititonFiles() []*FileSystem.FileLocation {
	return self.FileList
}

func (self *PartitionInfo) GetLocation(i int) string {
	if i >= len(self.Locations) {
		return ""
	}
	return self.Locations[i]
}

func (self *PartitionInfo) GetFileType(i int) FileSystem.FileType {
	if i >= len(self.FileTypes) {
		return FileSystem.UNKNOWNFILETYPE
	}
	return self.FileTypes[i]
}

func (self *PartitionInfo) Write(row *Util.Row) {
	self.Rows = append(self.Rows, row)
}

func (self *PartitionInfo) IsPartition() bool {
	if self.Metadata != nil && len(self.Metadata.Columns) > 0 {
		return true
	}
	return false
}
