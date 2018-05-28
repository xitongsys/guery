package HiveConnector

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/xitongsys/guery/Connector/FileReader"
	"github.com/xitongsys/guery/FileSystem"
	"github.com/xitongsys/guery/Util"
)

type HiveConnector struct {
	Config                 *HiveConnectorConfig
	Catalog, Schema, Table string
	Metadata               *Util.Metadata

	PartitionInfo    *Util.PartitionInfo
	PartitionReaders []FileReader.FileReader

	db *sql.DB
}

func NewHiveConnector(schema, table string) (*HiveConnector, error) {
	name := strings.Join([]string{"HIVE", schema, table}, ".")
	config := Configs.GetConfig(name)
	if config == nil {
		return nil, fmt.Errorf("Table not found")
	}
	res := &HiveConnector{
		Catalog: "HIVE",
		Schema:  schema,
		Table:   table,
	}
	if err := self.setMetadata(); err != nil {
		return res, err
	}

	if err := self.setPartitionInfo(); err != nil {
		return res, err
	}
	self.PartitionReaders = make([]FileReader.FileReader, self.PartitionInfo.GetPartitionNum())
	return res, nil
}

func (self *HiveConnector) GetMetadata() *Util.Metadata {
	return self.Metadata
}

func (self *HiveConnector) GetPartitionInfo() *Util.PartitionInfo {
	return self.PartitionInfo
}

func (self *HiveConnector) Read() (*Util.Row, error) {
	return nil, nil
}

func (self *HiveConnector) ReadByColumns(colIndexes []int) (*Util.Row, error) {
	return nil, nil
}

func (self *HiveConnector) ReadPartitionByColumns(parIndex int, colIndexes []int) (*Util.Row, error) {
	if parIndex >= len(self.PartitionReaders) {
		return nil, fmt.Errorf("partition not found")
	}
	if self.PartitionReaders[parIndex] == nil {
		vf, err := FileSystem.Open(self.PartitionInfo.GetLocation(parIndex))
		if err != nil {
			return nil, err
		}
		self.PartitionReaders[parIndex], err = FileReader.NewReader(vf, self.PartitionInfo.GetFileType(parIndex), self.Metadata)
		if err != nil {
			return nil, err
		}
	}

	row, err := self.PartitionReaders[parIndex].ReadByColumns(colIndexes)
	return row, err
}
