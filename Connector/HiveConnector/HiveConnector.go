package HiveConnector

import (
	"database/sql"
	"fmt"
	"io"
	"log"
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
	FileType       string
	FileReader     FileReader.FileReader

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

func (self *HiveConnector) GetPartitionInfo() *Util.PartitionInfo {
	return self.PartitionInfo
}

func (self *HiveConnector) Read() (*Util.Row, error) {
	return nil, nil
}

func (self *HiveConnector) SetPartitionRead(parIndex int) (err error) {
	//with partitions
	if self.PartitionInfo.IsPartition() {
		if parIndex >= self.PartitionInfo.GetPartitionNum() {
			return fmt.Errorf("Index out of partition number")
		}
		if self.FileList, err = FileSystem.List(self.PartitionInfo.GetLocation(parIndex)); err != nil {
			return err
		}
		self.FileType = self.PartitionInfo.GetFileType(parIndex)

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

	if self.FileReader == nil && self.FileIndex < len(self.FileList) {
		vf, err := FileSystem.Open(self.FileList[self.FileIndex].Location)
		if err != nil {
			return nil, err
		}
		self.FileReader, err = FileReader.NewReader(vf, self.FileType, self.Metadata)
		if err != nil {
			return nil, err
		}
		self.FileIndex++

	} else if self.FileReader == nil && self.FileIndex >= len(self.FileList) {
		return nil, io.EOF

	}

	row, err := self.FileReader.Read()
	log.Println("=======", colIndexes, self.FileList[0].Location, row, err)
	if err == io.EOF {
		self.FileReader = nil
		return self.Read()
	}
	if err != nil {
		return nil, err
	}
	return row, err

}
