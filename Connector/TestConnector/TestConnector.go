package TestConnector

import (
	"fmt"
	"io"

	"github.com/xitongsys/guery/Util"
)

type TestConnector struct {
	Metadata      *Util.Metadata
	Rows          []Util.Row
	Index         int64
	PartitionInfo *Util.PartitionInfo
}

func GenerateTestRows(columns []string) []Util.Row {
	res := []Util.Row{}
	for i := int64(0); i < int64(1000); i++ {
		row := &Util.Row{}
		for _, name := range columns {
			switch name {
			case "ID":
				row.AppendVals(int64(i))
			case "INT64":
				row.AppendVals(int64(i))
			case "FLOAT64":
				row.AppendVals(float64(i))
			case "STRING":
				row.AppendVals(fmt.Sprintf("%v", i))

			}
		}
		res = append(res, *row)
	}
	return res
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
		}
		col := Util.NewColumnMetadata(t, "test", "test", "test", name)
		res.AppendColumn(col)
	}
	return res
}

func NewTestConnector(schema, table string) (*TestConnector, error) {
	columns := []string{"ID", "INT64", "FLOAT64", "STRING"}
	var res *TestConnector
	switch table {
	case "test":
		res = &TestConnector{
			Metadata: GenerateTestMetadata(columns),
			Rows:     GenerateTestRows(columns),
			Index:    0,
		}
	}
	res.PartitionInfo = Util.NewPartitionInfo(nil)
	return res, nil
}

func (self *TestConnector) GetMetadata() *Util.Metadata {
	return self.Metadata
}

func (self *TestConnector) GetPartitionInfo() *Util.PartitionInfo {
	return self.PartitionInfo
}

func (self *TestConnector) Read() (*Util.Row, error) {
	if self.Index >= int64(len(self.Rows)) {
		self.Index = 0
		return nil, io.EOF
	}

	self.Index++
	return &self.Rows[self.Index-1], nil
}

func (self *TestConnector) SetPartitionRead(parIndex int) error {
	return nil
}

func (self *TestConnector) ReadByColumns(colIndexes []int) (*Util.Row, error) {
	if self.Index >= int64(len(self.Rows)) {
		self.Index = 0
		return nil, io.EOF
	}
	self.Index++
	row := &Util.Row{}
	for _, ci := range colIndexes {
		row.AppendVals(self.Rows[self.Index-1].Vals[ci])
	}
	return row, nil
}
