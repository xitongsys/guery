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

	TableLocation string
	PartitionInfo *Util.PartitionInfo

	PartitionIndex int
	FileList       []*FileSystem.FileLocation
	FileIndex      int
	FileReader     FileReader.FileReader

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

	if err := self.setTableLocation(); err != nil {
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

func (self *HiveConnector) SetReadPartition(parIndex int) (err error) {
	//with partitions
	if self.PartitionInfo.GetPartitionNum() > 0 {
		if parIndex > self.PartitionInfo.IsPartition() {
			return fmt.Errorf("Index out of partition number")
		}
		if self.FileList, err = FileSystem.List(self.PartitionInfo.GetLocation(parIndex)); err != nil {
			return err
		}

	} else { //no partitions
		if self.FileList, err = FileSystem.List(self.TableLocation); err != nil {
			return err
		}
	}
	self.PartitionIndex = parIndex
	self.FileIndex = 0
	self.FileReader = nil

	return nil
}

func (self *HiveConnector) ReadByColumns(colIndexes []int) (*Util.Row, error) {
	if self.FileReader == nil && self.FileIndex < len(self.FilePathList) {
		vf, err := FileSystem.Open(self.FileList[self.FileIndex].Location)
		if err != nil {
			return nil, err
		}
		self.FileReader, err = FileReader.NewReader(vf, self.FileType, self.Metadata)
		if err != nil {
			return nil, err
		}
		self.FileIndex++

	} else if self.FileReader == nil && self.FileIndex >= len(self.FilePathList) {
		return nil, io.EOF

	}

	row, err := self.FileReader.Read()
	if err == io.EOF {
		self.FileReader = nil
		return self.Read()
	}
	if err != nil {
		return nil, err
	}
	return row, err

}
