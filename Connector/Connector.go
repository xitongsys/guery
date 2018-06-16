package Connector

import (
	"fmt"

	"github.com/xitongsys/guery/Connector/FileConnector"
	"github.com/xitongsys/guery/Connector/HiveConnector"
	"github.com/xitongsys/guery/Connector/TestConnector"
	"github.com/xitongsys/guery/FileSystem"
	"github.com/xitongsys/guery/FileSystem/Partition"
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Row"
)

type Connector interface {
	GetMetadata() *Metadata.Metadata
	GetPartitionInfo() *Partition.PartitionInfo
	GetReader(file *FileSystem.FileLocation, md *Metadata.Metadata) func(indexes []int) (*Row.Row, error)
	ShowTables(schema string, like, escape *string) func() (*Row.Row, error)
	ShowSchemas(like, escape *string) func() (*Row.Row, error)
	ShowColumns(catalog, schema, table string) func() (*Row.Row, error)
}

func NewConnector(catalog string, schema string, table string) (Connector, error) {
	switch catalog {
	case "test":
		return TestConnector.NewTestConnector(schema, table)
	case "file":
		return FileConnector.NewFileConnector(schema, table)
	case "hive":
		return HiveConnector.NewHiveConnector(schema, table)

	}
	return nil, fmt.Errorf("NewConnector failed: table %s.%s.%s not found", catalog, schema, table)
}
