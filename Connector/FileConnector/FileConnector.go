package FileConnector

import (
	"fmt"
	"strings"

	"github.com/xitongsys/guery/FileSystem"
	"github.com/xitongsys/guery/FileSystem/FileReader"
	"github.com/xitongsys/guery/FileSystem/Partition"
	"github.com/xitongsys/guery/Util"
)

type FileConnector struct {
	Metadata      *Util.Metadata
	FileReader    FileReader.FileReader
	FileType      FileSystem.FileType
	PartitionInfo *Partition.PartitionInfo
}

func NewFileConnector(schema, table string) (*FileConnector, error) {
	var err error
	res := &FileConnector{}
	catalog := "file"
	key := strings.Join([]string{catalog, schema, table}, ".")
	conf := Configs.GetConfig(key)
	if conf == nil {
		return nil, fmt.Errorf("Table not found")
	}
	res.FileType = FileSystem.StringToFileType(conf.FileType)
	if res.Metadata, err = Util.NewMetadataFromJsonMetadata(&conf.FileMD); err != nil {
		return res, err
	}
	res.PartitionInfo = Partition.NewPartitionInfo(nil)
	for _, loc := range conf.PathList {
		fs, err := FileSystem.List(loc)
		if err != nil {
			return res, err
		}
		res.PartitionInfo.FileList = append(res.PartitionInfo.FileList, fs...)
	}
	return res, err
}

func (self *FileConnector) GetMetadata() *Util.Metadata {
	return self.Metadata
}

func (self *FileConnector) GetPartitionInfo() *Partition.PartitionInfo {
	return self.PartitionInfo
}
