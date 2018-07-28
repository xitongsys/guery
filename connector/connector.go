package connector

import (
	"fmt"

	"github.com/xitongsys/guery/connector/file"
	"github.com/xitongsys/guery/connector/hive"
	"github.com/xitongsys/guery/connector/test"
	"github.com/xitongsys/guery/filesystem"
	"github.com/xitongsys/guery/filesystem/partition"
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/row"
)

type Connector interface {
	GetMetadata() (*metadata.Metadata, error)
	GetPartitionInfo() (*partition.PartitionInfo, error)
	GetReader(file *filesystem.FileLocation, md *metadata.Metadata) func(indexes []int) (*row.RowsGroup, error)

	ShowTables(catalog, schema, table string, like, escape *string) func() (*row.Row, error)
	ShowSchemas(catalog, schema, table string, like, escape *string) func() (*row.Row, error)
	ShowColumns(catalog, schema, table string) func() (*row.Row, error)
	ShowPartitions(catalog, schema, table string) func() (*row.Row, error)
}

func NewConnector(catalog string, schema string, table string) (Connector, error) {
	if len(catalog) >= 4 {
		switch catalog[:4] {
		case "test":
			return test.NewTestConnector(catalog, schema, table)
		case "file":
			return file.NewFileConnector(catalog, schema, table)
		case "hive":
			return hive.NewHiveConnector(catalog, schema, table)
		}
	}
	return nil, fmt.Errorf("NewConnector failed: table %s.%s.%s not found", catalog, schema, table)
}
