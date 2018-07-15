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

func NewHiveConnectorEmpty() (*HiveConnector, error) {
	res := &HiveConnector{}
	return res, nil
}

func NewHiveConnector(ns ...string) (*HiveConnector, error) {
	catalog, schema, table := "", "", ""
	if len(ns) <= 0 {
		return nil, fmt.Errorf("[NewHiveConnector] less para")
	} else if len(ns) == 1 {
		catalog = ns[0]
	} else if len(ns) == 2 {
		catalog, schema = ns[0], ns[1]
	} else {
		catalog, schema, table = ns[0], ns[1], ns[2]
	}

	name := strings.Join([]string{catalog, schema, table}, ".")
	config := Config.Conf.HiveConnectorConfigs.GetConfig(name)
	if config == nil {
		return nil, fmt.Errorf("HiveConnector: table not found")
	}
	res := &HiveConnector{
		Config:  config,
		Catalog: catalog,
		Schema:  schema,
		Table:   table,
	}
	if len(ns) >= 3 && len(table) > 0 {
		if err := res.Init(); err != nil {
			return res, err
		}
	}
	return res, nil
}

func (self *HiveConnector) GetMetadata() (*Metadata.Metadata, error) {
	if self.Metadata == nil {
		if err := self.setMetadata(); err != nil {
			return nil, err
		}
	}
	return self.Metadata, nil
}

func (self *HiveConnector) GetPartitionInfo() (*Partition.PartitionInfo, error) {
	if self.PartitionInfo == nil {
		if err := self.setPartitionInfo(); err != nil {
			return nil, err
		}
	}
	return self.PartitionInfo, nil
}

func (self *HiveConnector) GetReader(file *FileSystem.FileLocation, md *Metadata.Metadata) func(indexes []int) (*Row.RowsGroup, error) {
	reader, err := FileReader.NewReader(file, md)

	return func(indexes []int) (*Row.RowsGroup, error) {
		var rg *Row.RowsGroup
		if err != nil {
			return nil, err
		}
		if rg, err = reader.Read(indexes); err != nil {
			return nil, err
		}
		return HiveTypeConvert(rg, indexes)
	}
}

func (self *HiveConnector) ShowTables(catalog, schema, table string, like, escape *string) func() (*Row.Row, error) {
	sqlStr := fmt.Sprintf(SHOWTABLES_SQL, schema)
	var rows *sql.Rows
	var err error
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

func (self *HiveConnector) ShowSchemas(catalog, schema, table string, like, escape *string) func() (*Row.Row, error) {
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

func (self *HiveConnector) ShowPartitions(catalog, schema, table string) func() (*Row.Row, error) {
	var err error
	var parInfo *Partition.PartitionInfo
	parInfo, err = self.GetPartitionInfo()
	i := 0

	return func() (*Row.Row, error) {
		if err != nil {
			return nil, err
		}
		row := parInfo.GetPartitionRow(i)
		if row == nil {
			err = io.EOF
		}
		i++
		return row, err
	}
}
