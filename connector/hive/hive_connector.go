package hive

import (
	"database/sql"
	"fmt"
	"io"
	"strings"

	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/filereader"
	"github.com/xitongsys/guery/filesystem"
	"github.com/xitongsys/guery/filesystem/partition"
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/row"
)

type HiveConnector struct {
	Config                 *config.HiveConnectorConfig
	Catalog, Schema, Table string
	Metadata               *metadata.Metadata

	TableLocation string
	FileType      filesystem.FileType
	PartitionInfo *partition.PartitionInfo

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
	config := config.Conf.HiveConnectorConfigs.GetConfig(name)
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

func (self *HiveConnector) GetMetadata() (*metadata.Metadata, error) {
	if self.Metadata == nil {
		if err := self.setMetadata(); err != nil {
			return nil, err
		}
	}
	return self.Metadata, nil
}

func (self *HiveConnector) GetPartitionInfo() (*partition.PartitionInfo, error) {
	if self.PartitionInfo == nil {
		if err := self.setPartitionInfo(); err != nil {
			return nil, err
		}
	}
	return self.PartitionInfo, nil
}

func (self *HiveConnector) GetReader(file *filesystem.FileLocation, md *metadata.Metadata) func(indexes []int) (*row.RowsGroup, error) {
	reader, err := filereader.NewReader(file, md)

	return func(indexes []int) (*row.RowsGroup, error) {
		var rg *row.RowsGroup
		if err != nil {
			return nil, err
		}
		if rg, err = reader.Read(indexes); err != nil {
			return nil, err
		}
		return HiveTypeConvert(rg)
	}
}

func (self *HiveConnector) ShowTables(catalog, schema, table string, like, escape *string) func() (*row.Row, error) {
	sqlStr := fmt.Sprintf(SHOWTABLES_SQL, schema)
	var rs *sql.Rows
	var err error
	if err = self.getConn(); err == nil {
		rs, err = self.db.Query(sqlStr)
	}

	return func() (*row.Row, error) {
		if err != nil {
			return nil, err
		}
		if rs.Next() {
			var table string
			rs.Scan(&table)
			r := row.NewRow()
			r.AppendVals(table)
			return r, nil

		} else {
			if err = rs.Err(); err == nil {
				err = io.EOF
			}

			return nil, err
		}
	}
}

func (self *HiveConnector) ShowSchemas(catalog, schema, table string, like, escape *string) func() (*row.Row, error) {
	var err error
	var rs *sql.Rows
	sqlStr := fmt.Sprintf(SHOWSCHEMAS_SQL)
	if err = self.getConn(); err == nil {
		rs, err = self.db.Query(sqlStr)
	}

	return func() (*row.Row, error) {
		if err != nil {
			return nil, err
		}
		if rs.Next() {
			var table string
			rs.Scan(&table)
			r := row.NewRow()
			r.AppendVals(table)
			return r, nil

		} else {
			if err = rs.Err(); err == nil {
				err = io.EOF
			}

			return nil, err
		}
	}
}

func (self *HiveConnector) ShowColumns(catalog, schema, table string) func() (*row.Row, error) {
	var err error
	var rs *sql.Rows
	sqlStr := fmt.Sprintf(MD_SQL, self.Schema, self.Table, self.Schema, self.Table)
	if err = self.getConn(); err == nil {
		rs, err = self.db.Query(sqlStr)
	}

	return func() (*row.Row, error) {
		if err != nil {
			return nil, err
		}
		if rs.Next() {
			var colName, colType string
			rs.Scan(&colName, &colType)
			colType = HiveTypeToGueryType(colType).String()
			r := row.NewRow()
			r.AppendVals(colName, colType)
			return r, nil

		} else {
			if err = rs.Err(); err == nil {
				err = io.EOF
			}

			return nil, err
		}
	}
}

func (self *HiveConnector) ShowPartitions(catalog, schema, table string) func() (*row.Row, error) {
	var err error
	var parInfo *partition.PartitionInfo
	parInfo, err = self.GetPartitionInfo()
	i := 0

	return func() (*row.Row, error) {
		if err != nil {
			return nil, err
		}
		r := parInfo.GetPartitionRow(i)
		if r == nil {
			err = io.EOF
		}
		i++
		return r, err
	}
}
