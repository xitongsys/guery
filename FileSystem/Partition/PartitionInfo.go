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

func (self *PartitionInfo) Split(n int) []*PartitionInfo {
	res := make([]*PartitionInfo, n)
	recMap := make([]map[int]int, n)

	for i := 0; i < n; i++ {
		res[i] = &PartitionInfo{
			Metadata:  self.Metadata,
			Rows:      []*Util.Row{},
			Locations: []string{},
			FileTypes: []FileSystem.FileTypes{},
			FileLists: [][]*FileSystem.FileLocation{},
			FileList:  []*FileSystem.FileLocation{},
		}
		recMap[i] = map[int]int{}
	}

	if self.IsPartition() {
		pn := self.GetPartitionNum()
		j, k := 0, 0
		ok := false
		for i := 0; i < pn; i++ {
			for _, file := range self.FileList[i] {
				if _, ok = recMap[k][i]; !ok {
					pi := res[k]
					recMap[k][i] = len(pi.Rows)
					pi.Rows = append(pi.Rows, self.Rows[i])
					pi.Locations = append(pi.Locations, self.Locations[i])
					pi.FileTypes = append(pi.FileTypes, self.FileTypes[i])
				}
				j := recMap[k][i]
				self.FileLists[j] = append(self.FileLists[j], file)
			}
		}

	} else {
		for i, file := range self.FileList {
			k := i % n
			res[k].FileList = append(res[k].FileList, file)
		}
	}
	return res
}

func (self *PartitionInfo) GetPartitionNum() int {
	return len(self.Rows)
}

func (self *PartitionInfo) GetPartition(i int) *Util.RowsGroup {
	if i >= len(self.Rows) {
		return nil
	}
	rb := Util.NewRowsGroup(self.Metadata)
	for _, row := range self.Rows {
		rb.Write(row)
	}
	return rb
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
