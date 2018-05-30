package Connector

import (
	"fmt"

	"github.com/xitongsys/guery/Connector/FileConnector"
	"github.com/xitongsys/guery/Connector/HiveConnector"
	"github.com/xitongsys/guery/Connector/TestConnector"
	"github.com/xitongsys/guery/Util"
)

type Connector interface {
	GetMetadata() *Util.Metadata
	GetPartitionInfo() *Util.PartitionInfo
	Read() (*Util.Row, error)
	ReadByColumns(colIndexes []int) (*Util.Row, error)
	SetPartitionRead(parIndex int) error
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
