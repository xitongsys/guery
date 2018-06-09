package HiveConnector

import (
	"database/sql"
	"fmt"
	"strings"

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
	return self.PartitionInfo
}

func (self *HiveConnector) Read() (*Util.Row, error) {
	return nil, nil
}

func (self *HiveConnector) ReadByColumns(colIndexes []int) (*Util.Row, error) {
	return nil, nil
}
