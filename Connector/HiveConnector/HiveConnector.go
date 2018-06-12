package HiveConnector

import (
	"database/sql"
	"fmt"
	"io"
	"strings"

	"github.com/xitongsys/guery/FileReader"
	"github.com/xitongsys/guery/FileSystem"
	"github.com/xitongsys/guery/FileSystem/Partition"
	"github.com/xitongsys/guery/Util"
)

type HiveConnector struct {
	Config                 *HiveConnectorConfig
	Catalog, Schema, Table string
	Metadata               *Util.Metadata

	TableLocation string
	FileType      FileSystem.FileType
	PartitionInfo *Partition.PartitionInfo

	db *sql.DB
}

func NewHiveConnector(schema, table string) (*HiveConnector, error) {
	name := strings.Join([]string{"hive", schema, table}, ".")
	config := Configs.GetConfig(name)
	if config == nil {
		return nil, fmt.Errorf("Table not found")
	}
	res := &HiveConnector{
		Config:  config,
		Catalog: "hive",
		Schema:  schema,
		Table:   table,
	}
	if err := res.Init(); err != nil {
		return res, err
	}
	return res, nil
}

func (self *HiveConnector) GetMetadata() *Util.Metadata {
	return self.Metadata
}

func (self *HiveConnector) GetPartitionInfo() *Partition.PartitionInfo {
	if self.PartitionInfo == nil {
		if err := self.setPartitionInfo(); err != nil {
			return nil
		}
	}
	return self.PartitionInfo
}

func (self *HiveConnector) GetReader(file *FileSystem.FileLocation, md *Util.Metadata) func(indexes []int) (*Util.Row, error) {
	reader, err := FileReader.NewReader(file, md)

	return func(indexes []int) (*Util.Row, error) {
		if err != nil {
			return nil, err
		}
		return reader.Read(indexes)
	}
}
