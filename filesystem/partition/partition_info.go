package partition

import (
	"github.com/xitongsys/guery/filesystem"
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/row"
)

type PartitionInfo struct {
	Metadata   *metadata.Metadata
	Partitions []*Partition
	Locations  []string
	FileTypes  []filesystem.FileType
	FileLists  [][]*filesystem.FileLocation

	//for no partition
	FileList []*filesystem.FileLocation
}

func NewPartitionInfo(md *metadata.Metadata) *PartitionInfo {
	res := &PartitionInfo{
		Metadata:  md,
		Locations: []string{},
		FileTypes: []filesystem.FileType{},
		FileLists: [][]*filesystem.FileLocation{},

		FileList: []*filesystem.FileLocation{},
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

func (self *PartitionInfo) GetPartitionRowGroup(i int) *row.RowsGroup {
	r := self.GetPartitionRow(i)
	if r == nil {
		return nil
	}
	rb := row.NewRowsGroup(self.Metadata)
	rb.Write(r)
	return rb
}

func (self *PartitionInfo) GetPartitionRow(i int) *row.Row {
	if i >= self.GetPartitionNum() {
		return nil
	}
	row := new(row.Row)
	for j := 0; j < len(self.Partitions); j++ {
		row.AppendVals(self.Partitions[j].Vals[i])
	}
	return row
}

func (self *PartitionInfo) GetPartitionFiles(i int) []*filesystem.FileLocation {
	if i >= len(self.FileLists) {
		return []*filesystem.FileLocation{}
	}
	return self.FileLists[i]
}

func (self *PartitionInfo) GetNoPartititonFiles() []*filesystem.FileLocation {
	return self.FileList
}

func (self *PartitionInfo) GetLocation(i int) string {
	if i >= len(self.Locations) {
		return ""
	}
	return self.Locations[i]
}

func (self *PartitionInfo) GetFileType(i int) filesystem.FileType {
	if i >= len(self.FileTypes) {
		return filesystem.UNKNOWNFILETYPE
	}
	return self.FileTypes[i]
}

func (self *PartitionInfo) Write(row *row.Row) {
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
