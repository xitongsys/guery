package TestConnector

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/xitongsys/guery/FileSystem"
	"github.com/xitongsys/guery/FileSystem/Partition"
	"github.com/xitongsys/guery/Util"
)

type TestConnector struct {
	Metadata      *Util.Metadata
	Rows          []Util.Row
	Index         int64
	PartitionInfo *Partition.PartitionInfo
}

func GenerateTestRows(columns []string) error {
	f, err := os.Create("/tmp/test.csv")
	if err != nil {
		return err
	}
	defer f.Close()

	for i := int64(0); i < int64(1000); i++ {
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

func GenerateTestMetadata(columns []string) *Util.Metadata {
	res := Util.NewMetadata()
	for _, name := range columns {
		t := Util.UNKNOWNTYPE
		switch name {
		case "ID":
			t = Util.INT64
		case "INT64":
			t = Util.INT64
		case "FLOAT64":
			t = Util.FLOAT64
		case "STRING":
			t = Util.STRING
		case "TIMEVAL":
			t = Util.TIMESTAMP
		}
		col := Util.NewColumnMetadata(t, "test", "test", "test", name)
		res.AppendColumn(col)
	}
	return res
}

func NewTestConnector(schema, table string) (*TestConnector, error) {
	columns := []string{"ID", "INT64", "FLOAT64", "STRING", "TIMEVAL"}
	var res *TestConnector
	switch table {
	case "test":
		res = &TestConnector{
			Metadata: GenerateTestMetadata(columns),
			Index:    0,
		}
		GenerateTestRows(columns)
	}
	res.PartitionInfo = Partition.NewPartitionInfo(Util.NewMetadata())
	res.PartitionInfo.FileList = []*FileSystem.FileLocation{
		&FileSystem.FileLocation{
			Location: "/tmp/test.csv",
			FileType: FileSystem.CSV,
		},
	}
	return res, nil
}

func (self *TestConnector) GetMetadata() *Util.Metadata {
	return self.Metadata
}

func (self *TestConnector) GetPartitionInfo() *Partition.PartitionInfo {
	return self.PartitionInfo
}
