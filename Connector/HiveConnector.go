package Connector

import (
	"io"
	"strings"

	"github.com/xitongsys/guery/Config"
	"github.com/xitongsys/guery/Connector/FileReader"
	"github.com/xitongsys/guery/FileSystem"
	"github.com/xitongsys/guery/Util"
)

type HiveConnector struct {
	Host, User, Password   string
	Catalog, Schema, Table string
}

func NewHiveConnector(schema, table string) (*HiveConnector, error) {

	return nil, nil
}

func (self *HiveConnector) GetMetadata() *Util.Metadata {
	return self.Metadata
}

func (self *HiveConnector) GetPartitionInfo() *Util.PartitionInfo {
	return nil
}

func (self *HiveConnector) Read() (*Util.Row, error) {
}

func (self *HiveConnector) ReadByColumns(colIndexes []int) (*Util.Row, error) {

}
