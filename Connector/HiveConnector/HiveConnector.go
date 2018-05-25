package HiveConnector

import (
	"database/sql"
	"fmt"
	"io"
	"strings"

	"github.com/xitongsys/guery/Config"
	"github.com/xitongsys/guery/Connector/FileReader"
	"github.com/xitongsys/guery/FileSystem"
	"github.com/xitongsys/guery/Util"
)

type HiveConnector struct {
	Config                 *HiveConnectorConfig
	Catalog, Schema, Table string
	Metadata               *Util.Metadata

	db *sql.DB
}

func NewHiveConnector(schema, table string) (*HiveConnector, error) {
	name := strings.Join([]string{"HIVE", schema, table}, ".")
	config := Config.Conf.HiveConnectorConfigs.GetConfig(name)
	if config == nil {
		return nil, fmt.Errorf("Table not found")
	}
	res := &HiveConnector{
		Catalog: "HIVE",
		Schema:  schema,
		Table:   table,
	}
	return res, nil
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
