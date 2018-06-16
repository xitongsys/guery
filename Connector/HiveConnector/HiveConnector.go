package HiveConnector

import (
	"database/sql"
	"fmt"
	"io"
	"strings"

	"github.com/xitongsys/guery/Config"
	"github.com/xitongsys/guery/FileReader"
	"github.com/xitongsys/guery/FileSystem"
	"github.com/xitongsys/guery/FileSystem/Partition"
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Row"
)

type HiveConnector struct {
	Config                 *Config.HiveConnectorConfig
	Catalog, Schema, Table string
	Metadata               *Metadata.Metadata

	TableLocation string
	FileType      FileSystem.FileType
	PartitionInfo *Partition.PartitionInfo

	db *sql.DB
}

func NewHiveConnector(schema, table string) (*HiveConnector, error) {
	name := strings.Join([]string{"hive", schema, table}, ".")
	config := Config.Conf.HiveConnectorConfigs.GetConfig(name)
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

func (self *HiveConnector) GetMetadata() *Metadata.Metadata {
	return self.Metadata
}

func (self *HiveConnector) GetPartitionInfo() *Partition.PartitionInfo {
	if self.PartitionInfo == nil {
		if err := self.setPartitionInfo(); err != nil {
			return nil
		}
	}
	return self.PartitionInfo
}

func (self *HiveConnector) GetReader(file *FileSystem.FileLocation, md *Metadata.Metadata) func(indexes []int) (*Row.Row, error) {
	reader, err := FileReader.NewReader(file, md)

	return func(indexes []int) (*Row.Row, error) {
		var row *Row.Row
		if err != nil {
			return nil, err
		}
		if row, err = reader.Read(indexes); err != nil {
			return row, err
		}
		return HiveTypeConvert(row, md)
	}
}

func (self *HiveConnector) ShowTables(schema string, like, escape *string) func() (*Row.Row, error) {
	sqlStr := fmt.Sprintf(SHOWTABLES_SQL, schema)
	var rows *sql.Rows
	var err error
	rows, err = self.db.Query(sqlStr)

	return func() (*Row.Row, error) {
		if err != nil {
			return nil, err
		}
		if rows.Next() {
			var table string
			rows.Scan(&table)
			row := Row.NewRow()
			row.AppendVals(table)
			return row, nil

		} else {
			if err = rows.Err(); err == nil {
				err = io.EOF
			}

			return nil, err
		}
	}
}

func (self *HiveConnector) ShowSchemas(like, escape *string) func() (*Row.Row, error) {
	var err error
	var rows *sql.Rows
	sqlStr := fmt.Sprintf(SHOWSCHEMAS_SQL)
	if err = self.getConn(); err == nil {
		rows, err = self.db.Query(sqlStr)
	}

	return func() (*Row.Row, error) {
		if err != nil {
			return nil, err
		}
		if rows.Next() {
			var table string
			rows.Scan(&table)
			row := Row.NewRow()
			row.AppendVals(table)
			return row, nil

		} else {
			if err = rows.Err(); err == nil {
				err = io.EOF
			}

			return nil, err
		}
	}
}

func (self *HiveConnector) ShowColumns(catalog, schema, table string) func() (*Row.Row, error) {
	var err error
	var rows *sql.Rows
	sqlStr := fmt.Sprintf(MD_SQL, self.Schema, self.Table, self.Schema, self.Table)
	if err = self.getConn(); err == nil {
		rows, err = self.db.Query(sqlStr)
	}

	return func() (*Row.Row, error) {
		if err != nil {
			return nil, err
		}
		if rows.Next() {
			var colName, colType string
			rows.Scan(&colName, &colType)
			colType = HiveTypeToGueryType(colType).String()
			row := Row.NewRow()
			row.AppendVals(colName, colType)
			return row, nil

		} else {
			if err = rows.Err(); err == nil {
				err = io.EOF
			}

			return nil, err
		}
	}
}
