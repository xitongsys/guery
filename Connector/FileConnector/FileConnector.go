package FileConnector

import (
	"fmt"
	"io"
	"strings"

	"github.com/xitongsys/guery/Config"
	"github.com/xitongsys/guery/FileReader"
	"github.com/xitongsys/guery/FileSystem"
	"github.com/xitongsys/guery/FileSystem/Partition"
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Row"
	"github.com/xitongsys/guery/Type"
)

type FileConnector struct {
	Config        *Config.FileConnectorConfig
	Metadata      *Metadata.Metadata
	FileReader    FileReader.FileReader
	FileType      FileSystem.FileType
	PartitionInfo *Partition.PartitionInfo
}

func NewFileConnector(schema, table string) (*FileConnector, error) {
	var err error
	res := &FileConnector{}
	catalog := "file"
	key := strings.Join([]string{catalog, schema, table}, ".")
	conf := Config.Conf.FileConnectorConfigs.GetConfig(key)
	if conf == nil {
		return nil, fmt.Errorf("Table not found")
	}
	res.Config = conf
	res.FileType = FileSystem.StringToFileType(conf.FileType)
	res.Metadata, err = NewFileMetadata(conf)
	return res, err
}

func NewFileMetadata(conf *Config.FileConnectorConfig) (*Metadata.Metadata, error) {
	res := Metadata.NewMetadata()
	if len(conf.ColumnNames) != len(conf.ColumnTypes) {
		return res, fmt.Errorf("File Config error: ColumnNames and ColumnTypes not match")
	}

	for i := 0; i < len(conf.ColumnNames); i++ {
		col := &Metadata.ColumnMetadata{
			Catalog:    conf.Catalog,
			Schema:     conf.Schema,
			Table:      conf.Table,
			ColumnName: conf.ColumnNames[i],
			ColumnType: Type.TypeNameToType(conf.ColumnTypes[i]),
		}
		res.AppendColumn(col)
	}

	res.Reset()
	return res, nil
}

func (self *FileConnector) GetMetadata() *Metadata.Metadata {
	return self.Metadata
}

func (self *FileConnector) GetPartitionInfo() *Partition.PartitionInfo {
	if self.PartitionInfo == nil {
		self.setPartitionInfo()
	}
	return self.PartitionInfo
}

func (self *FileConnector) setPartitionInfo() {
	parMD := Metadata.NewMetadata()
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

func (self *FileConnector) GetReader(file *FileSystem.FileLocation, md *Metadata.Metadata) func(indexes []int) (*Row.Row, error) {
	reader, err := FileReader.NewReader(file, md)
	return func(indexes []int) (*Row.Row, error) {
		if err != nil {
			return nil, err
		}
		return reader.Read(indexes)
	}
}

func (self *FileConnector) ShowTables(schema string, like, escape *string) func() (*Row.Row, error) {
	return func() (*Row.Row, error) {
		return nil, io.EOF
	}
}

func (self *FileConnector) ShowSchemas(like, escape *string) func() (*Row.Row, error) {
	return func() (*Row.Row, error) {
		return nil, io.EOF
	}
}
