package FileConnector

import (
	"fmt"
	"strings"

	"github.com/xitongsys/guery/FileReader"
	"github.com/xitongsys/guery/FileSystem"
	"github.com/xitongsys/guery/FileSystem/Partition"
	"github.com/xitongsys/guery/Util"
)

type FileConnector struct {
	Config        *FileConnectorConfig
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
	res.Config = conf
	res.FileType = FileSystem.StringToFileType(conf.FileType)
	res.Metadata, err = Util.NewMetadataFromJsonMetadata(&conf.FileMD)
	return res, err
}

func (self *FileConnector) GetMetadata() *Util.Metadata {
	return self.Metadata
}

func (self *FileConnector) GetPartitionInfo() *Partition.PartitionInfo {
	if self.PartitionInfo == nil {
		self.setPartitionInfo()
	}
	return self.PartitionInfo
}

func (self *FileConnector) setPartitionInfo() {
	parMD := Util.NewMetadata()
	self.PartitionInfo = Partition.NewPartitionInfo(parMD)
	for _, loc := range self.Config.PathList {
		fs, err := FileSystem.List(loc)
		if err != nil {
			return
		}
		for _, f := range fs {
			f.FileType = self.FileType
		}
		self.PartitionInfo.FileList = append(self.PartitionInfo.FileList, fs...)
	}
}

func (self *FileConnector) GetReader(file *FileSystem.FileLocation, md *Util.Metadata) func(indexes []int) (*Util.Row, error) {
	reader, err := FileReader.NewReader(file, md)
	return func(indexes []int) (*Util.Row, error) {
		if err != nil {
			return nil, err
		}
		return reader.Read(indexes)
	}
}
