package Catalog

import (
	"fmt"
	"io"
	"strings"

	"github.com/xitongsys/guery/Util"
)

type TestCatalog struct {
	Metadata *Util.Metadata
	Rows     []Util.Row
	Index    int64
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
		col := Util.NewColumnMetadata(t, "TEST", "TEST", "TEST", name)
		res.AppendColumn(col)
	}
	return res
}

func NewTestCatalog(schema, table string) *TestCatalog {
	columns := []string{"ID", "INT64", "FLOAT64", "STRING"}
	schema, table = strings.ToUpper(schema), strings.ToUpper(table)
	var res *TestCatalog
	switch table {
	case "TEST":
		res = &TestCatalog{
			Metadata: GenerateTestMetadata(columns),
			Rows:     GenerateTestRows(columns),
			Index:    0,
		}
	}
	return res
}

func (self *TestCatalog) GetMetadata() *Util.Metadata {
	return self.Metadata
}

func (self *TestCatalog) ReadRow() (*Util.Row, error) {
	if self.Index >= int64(len(self.Rows)) {
		self.Index = 0
		return nil, io.EOF
	}

	self.Index++
	return &self.Rows[self.Index-1], nil
}

func (self *TestCatalog) SkipTo(index, total int64) {
	ln := int64(len(self.Rows))
	pn := ln / total
	left := ln % total
	if left > index {
		left = index
	}
	self.Index = pn*index + left
}

func (self *TestCatalog) SkipRows(num int64) {
	self.Index += num
}

func (self *TestCatalog) ReadRowByColumns(colIndexes []int) (*Util.Row, error) {
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
