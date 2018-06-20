package TestConnector

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/xitongsys/guery/FileReader"
	"github.com/xitongsys/guery/FileSystem"
	"github.com/xitongsys/guery/FileSystem/Partition"
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Row"
	"github.com/xitongsys/guery/Type"
)

type TestConnector struct {
	Metadata      *Metadata.Metadata
	Rows          []Row.Row
	Index         int64
	PartitionInfo *Partition.PartitionInfo
}

var columns = []string{"ID", "INT64", "FLOAT64", "STRING", "TIMEVAL"}

func GenerateTestRows(columns []string) error {
	f, err := os.Create("/tmp/test.csv")
	if err != nil {
		return err
	}
	defer f.Close()

	for i := int64(0); i < int64(1000000); i++ {
		res := []string{}
		for _, name := range columns {
			switch name {
			case "ID":
				res = append(res, fmt.Sprintf("%v", i))
			case "INT64":
				res = append(res, fmt.Sprintf("%v", int64(-1*i)))
			case "FLOAT64":
				res = append(res, fmt.Sprintf("%v", float64(i)))
			case "STRING":
				res = append(res, fmt.Sprintf("s%v", i))
			case "TIMEVAL":
				res = append(res, fmt.Sprintf("%v", time.Now().Format("2006-01-02 15:04:05")))
			}
		}
		s := strings.Join(res, ",") + "\n"
		f.Write([]byte(s))
	}
	return nil
}

func GenerateTestMetadata(columns []string) *Metadata.Metadata {
	res := Metadata.NewMetadata()
	for _, name := range columns {
		t := Type.UNKNOWNTYPE
		switch name {
		case "ID":
			t = Type.INT64
		case "INT64":
			t = Type.INT64
		case "FLOAT64":
			t = Type.FLOAT64
		case "STRING":
			t = Type.STRING
		case "TIMEVAL":
			t = Type.TIMESTAMP
		}
		col := Metadata.NewColumnMetadata(t, "test", "test", "test", name)
		res.AppendColumn(col)
	}
	return res
}

func NewTestConnector(catalog, schema, table string) (*TestConnector, error) {
	if catalog != "test" || schema != "test" || table != "test" {
		return nil, fmt.Errorf("[NewTestConnector] table not found")
	}
	var res *TestConnector
	switch table {
	case "test":
		res = &TestConnector{
			Metadata: GenerateTestMetadata(columns),
			Index:    0,
		}
	}

	return res, nil
}

func (self *TestConnector) GetMetadata() *Metadata.Metadata {
	return self.Metadata
}

func (self *TestConnector) GetPartitionInfo() *Partition.PartitionInfo {
	if self.PartitionInfo == nil {
		self.PartitionInfo = Partition.NewPartitionInfo(Metadata.NewMetadata())
		self.PartitionInfo.FileList = []*FileSystem.FileLocation{
			&FileSystem.FileLocation{
				Location: "/tmp/test.csv",
				FileType: FileSystem.CSV,
			},
		}
		GenerateTestRows(columns)
	}
	return self.PartitionInfo
}

func (self *TestConnector) GetReader(file *FileSystem.FileLocation, md *Metadata.Metadata) func(indexes []int) (*Row.Row, error) {
	reader, err := FileReader.NewReader(file, md)
	return func(indexes []int) (*Row.Row, error) {
		if err != nil {
			return nil, err
		}
		return reader.Read(indexes)
	}
}

func (self *TestConnector) ShowTables(catalog, schema, table string, like, escape *string) func() (*Row.Row, error) {
	var err error
	row := Row.NewRow()
	row.AppendVals("test")
	i := 0
	return func() (*Row.Row, error) {
		if err != nil {
			return nil, err
		}
		if i > 0 {
			return nil, io.EOF
		}
		i++
		return row, nil
	}
}

func (self *TestConnector) ShowSchemas(catalog, schema, table string, like, escape *string) func() (*Row.Row, error) {
	var err error

	row := Row.NewRow()
	row.AppendVals("test")
	i := 0
	return func() (*Row.Row, error) {
		if err != nil {
			return nil, err
		}
		if i > 0 {
			return nil, io.EOF
		}
		i++
		return row, nil
	}
}

func (self *TestConnector) ShowColumns(catalog, schema, table string) func() (*Row.Row, error) {
	var err error
	row, rows := Row.NewRow(), []*Row.Row{}

	row = Row.NewRow()
	row.AppendVals("ID", "INT64")
	rows = append(rows, row)

	row = Row.NewRow()
	row.AppendVals("INT64", "INT64")
	rows = append(rows, row)

	row = Row.NewRow()
	row.AppendVals("FLOAT64", "FLOAT64")
	rows = append(rows, row)

	row = Row.NewRow()
	row.AppendVals("STRING", "STRING")
	rows = append(rows, row)

	row = Row.NewRow()
	row.AppendVals("TIMESTAMP", "TIMESTAMP")
	rows = append(rows, row)

	i := 0
	return func() (*Row.Row, error) {
		if err != nil {
			return nil, err
		}
		if i >= len(rows) {
			return nil, io.EOF
		}
		i++
		return rows[i-1], nil
	}
}

func (self *TestConnector) ShowPartitions(catalog, schema, table string) func() (*Row.Row, error) {
	var err error

	return func() (*Row.Row, error) {
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	}
}
