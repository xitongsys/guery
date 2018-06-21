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
	GetMetadata() (*Metadata.Metadata, error)
	GetPartitionInfo() (*Partition.PartitionInfo, error)
	GetReader(file *FileSystem.FileLocation, md *Metadata.Metadata) func(indexes []int) (*Row.Row, error)

	ShowTables(catalog, schema, table string, like, escape *string) func() (*Row.Row, error)
	ShowSchemas(catalog, schema, table string, like, escape *string) func() (*Row.Row, error)
	ShowColumns(catalog, schema, table string) func() (*Row.Row, error)
	ShowPartitions(catalog, schema, table string) func() (*Row.Row, error)
}

func NewConnector(catalog string, schema string, table string) (Connector, error) {
	if len(catalog) >= 4 {
		switch catalog[:4] {
		case "test":
			return TestConnector.NewTestConnector(catalog, schema, table)
		case "file":
			return FileConnector.NewFileConnector(catalog, schema, table)
		case "hive":
			return HiveConnector.NewHiveConnector(catalog, schema, table)
		}
	}
	return nil, fmt.Errorf("NewConnector failed: table %s.%s.%s not found", catalog, schema, table)
}
