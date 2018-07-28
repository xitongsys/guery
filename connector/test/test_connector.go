package test

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/xitongsys/guery/filereader"
	"github.com/xitongsys/guery/filesystem"
	"github.com/xitongsys/guery/filesystem/partition"
	"github.com/xitongsys/guery/gtype"
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/row"
)

type TestConnector struct {
	Metadata      *metadata.Metadata
	Rows          []row.Row
	Index         int64
	Table         string
	PartitionInfo *partition.PartitionInfo
}

var columns = []string{"process_date", "var1", "var2", "var3", "data_source", "network_id", "event_date"}

func GenerateTestRows(columns []string) error {
	f1, err := os.Create("/tmp/test01.csv")
	if err != nil {
		return err
	}
	f2, err := os.Create("/tmp/test02.csv")
	if err != nil {
		return err
	}
	defer f1.Close()
	defer f2.Close()

	for i := int64(0); i < int64(100); i++ {
		res := []string{}
		for _, name := range columns {
			switch name {
			case "process_date":
				res = append(res, fmt.Sprintf("%v", time.Now().Format("2006-01-02 15:04:05")))
			case "var1":
				res = append(res, fmt.Sprintf("%v", i))
			case "var2":
				res = append(res, fmt.Sprintf("%v", float64(i)))
			case "var3":
				res = append(res, fmt.Sprintf("%v", "var3"))
			case "network_id":
				res = append(res, fmt.Sprintf("%v", i))
			case "data_source":
				res = append(res, fmt.Sprintf("data_source%v", i))
			case "event_date":
				res = append(res, fmt.Sprintf("%v", time.Now().Format("2006-01-02 15:04:05")))
			}
		}
		s := strings.Join(res, ",") + "\n"
		f1.Write([]byte(s))
		f2.Write([]byte(s))
	}
	return nil
}

func GenerateTestMetadata(columns []string, table string) *metadata.Metadata {
	res := metadata.NewMetadata()
	for _, name := range columns {
		t := gtype.UNKNOWNTYPE
		switch name {
		case "process_date":
			t = gtype.TIMESTAMP
		case "var1":
			t = gtype.INT64
		case "var2":
			t = gtype.FLOAT64
		case "var3":
			t = gtype.STRING
		case "data_source":
			t = gtype.STRING
		case "network_id":
			t = gtype.INT64
		case "event_date":
			t = gtype.DATE
		}
		col := metadata.NewColumnMetadata(t, "test", "test", table, name)
		res.AppendColumn(col)
	}
	return res
}

func NewTestConnector(catalog, schema, table string) (*TestConnector, error) {
	if catalog != "test" || schema != "test" {
		return nil, fmt.Errorf("[NewTestConnector] table not found")
	}
	var res *TestConnector
	switch table {
	case "csv", "parquet", "orc":
		res = &TestConnector{
			Metadata: GenerateTestMetadata(columns, table),
			Index:    0,
			Table:    table,
		}
	}

	return res, nil
}

func (self *TestConnector) GetMetadata() (*metadata.Metadata, error) {
	return self.Metadata, nil
}

func (self *TestConnector) GetPartitionInfo() (*partition.PartitionInfo, error) {
	if self.PartitionInfo == nil {
		self.PartitionInfo = partition.NewPartitionInfo(metadata.NewMetadata())
		if self.Table == "csv" {
			self.PartitionInfo.FileList = []*filesystem.FileLocation{
				&filesystem.FileLocation{
					Location: "/tmp/test01.csv",
					FileType: filesystem.CSV,
				},
				&filesystem.FileLocation{
					Location: "/tmp/test02.csv",
					FileType: filesystem.CSV,
				},
			}
			GenerateTestRows(columns)

		} else if self.Table == "parquet" {
			self.PartitionInfo.FileList = []*filesystem.FileLocation{
				&filesystem.FileLocation{
					Location: "/tmp/test.parquet",
					FileType: filesystem.PARQUET,
				},
			}
		} else if self.Table == "orc" {
			self.PartitionInfo.FileList = []*filesystem.FileLocation{
				&filesystem.FileLocation{
					Location: "/tmp/test.orc",
					FileType: filesystem.ORC,
				},
			}
		}

	}
	return self.PartitionInfo, nil
}

func (self *TestConnector) GetReader(file *filesystem.FileLocation, md *metadata.Metadata) func(indexes []int) (*row.RowsGroup, error) {
	reader, err := filereader.NewReader(file, md)
	return func(indexes []int) (*row.RowsGroup, error) {
		if err != nil {
			return nil, err
		}
		return reader.Read(indexes)
	}
}

func (self *TestConnector) ShowTables(catalog, schema, table string, like, escape *string) func() (*row.Row, error) {
	var err error
	tables := []string{"csv", "parquet", "orc"}
	rs := []*row.Row{}
	for _, table := range tables {
		r := row.NewRow()
		r.AppendVals(table)
		rs = append(rs, r)
	}
	i := 0
	return func() (*row.Row, error) {
		if err != nil {
			return nil, err
		}
		if i >= len(tables) {
			return nil, io.EOF
		}
		i++
		return rs[i-1], nil
	}
}

func (self *TestConnector) ShowSchemas(catalog, schema, table string, like, escape *string) func() (*row.Row, error) {
	var err error

	r := row.NewRow()
	r.AppendVals("test")
	i := 0
	return func() (*row.Row, error) {
		if err != nil {
			return nil, err
		}
		if i > 0 {
			return nil, io.EOF
		}
		i++
		return r, nil
	}
}

func (self *TestConnector) ShowColumns(catalog, schema, table string) func() (*row.Row, error) {
	var err error
	r, rs := row.NewRow(), []*row.Row{}

	r = row.NewRow()
	r.AppendVals("ID", "INT64")
	rs = append(rs, r)

	r = row.NewRow()
	r.AppendVals("INT64", "INT64")
	rs = append(rs, r)

	r = row.NewRow()
	r.AppendVals("FLOAT64", "FLOAT64")
	rs = append(rs, r)

	r = row.NewRow()
	r.AppendVals("STRING", "STRING")
	rs = append(rs, r)

	r = row.NewRow()
	r.AppendVals("TIMESTAMP", "TIMESTAMP")
	rs = append(rs, r)

	i := 0
	return func() (*row.Row, error) {
		if err != nil {
			return nil, err
		}
		if i >= len(rs) {
			return nil, io.EOF
		}
		i++
		return rs[i-1], nil
	}
}

func (self *TestConnector) ShowPartitions(catalog, schema, table string) func() (*row.Row, error) {
	var err error

	return func() (*row.Row, error) {
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	}
}
