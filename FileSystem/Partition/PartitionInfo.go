package Partition

import (
	"github.com/xitongsys/guery/FileSystem"
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Row"
)

type PartitionInfo struct {
	Metadata   *Metadata.Metadata
	Partitions []*Partition
	Locations  []string
	FileTypes  []FileSystem.FileType
	FileLists  [][]*FileSystem.FileLocation

	//for no partition
	FileList []*FileSystem.FileLocation
}

func NewPartitionInfo(md *Metadata.Metadata) *PartitionInfo {
	res := &PartitionInfo{
		Metadata:  md,
		Locations: []string{},
		FileTypes: []FileSystem.FileType{},
		FileLists: [][]*FileSystem.FileLocation{},

		FileList: []*FileSystem.FileLocation{},
	}
	for i := 0; i < md.GetColumnNumber(); i++ {
		t, _ := md.GetTypeByIndex(i)
		par := NewPartition(t)
		res.Partitions = append(res.Partitions, par)
	}
	return res
}

func (self *PartitionInfo) GetPartitionColumnNum() int {
	return len(self.Partitions)
}

func (self *PartitionInfo) GetPartitionNum() int {
	if len(self.Partitions) <= 0 {
		return 0
	}
	return len(self.Partitions[0].Vals)
}

func (self *PartitionInfo) GetPartitionRowGroup(i int) *Row.RowsGroup {
	row := self.GetPartitionRow(i)
	if row == nil {
		return nil
	}
	rb := Row.NewRowsGroup(self.Metadata)
	rb.Write(row)
	return rb
}

func (self *PartitionInfo) GetPartitionRow(i int) *Row.Row {
	if i >= self.GetPartitionNum() {
		return nil
	}
	row := new(Row.Row)
	for j := 0; j < len(self.Partitions); j++ {
		row.AppendVals(self.Partitions[j].Vals[i])
	}
	return row
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

func (self *PartitionInfo) Write(row *Row.Row) {
	for i, val := range row.Vals {
		self.Partitions[i].Append(val)
	}
}

func (self *PartitionInfo) IsPartition() bool {
	if self.Metadata != nil && len(self.Metadata.Columns) > 0 {
		return true
	}
	return false
}

func (self *PartitionInfo) Encode() {
	for _, par := range self.Partitions {
		par.Encode()
	}
}

func (self *PartitionInfo) Decode() error {
	for _, par := range self.Partitions {
		if err := par.Decode(); err != nil {
			return err
		}
	}
	return nil
}
