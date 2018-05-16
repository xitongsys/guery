package Catalog

import (
	"fmt"
	"io"
	"strings"

	"github.com/xitongsys/guery/Util"
)

type TestCatalog struct {
	Metadata Util.Metadata
	Rows     []Util.Row
	Index    int64
}

var TestMetadata = Util.Metadata{
	Catalog:     "TEST",
	Schema:      "DEFAULT",
	Table:       "TEST",
	ColumnNames: []string{"ID", "INT64", "FLOAT64", "STRING"},
	ColumnTypes: []Util.Type{Util.INT64, Util.INT64, Util.FLOAT64, Util.STRING},
}

var TestRows = []Util.Row{}

func GenerateTestRows() {
	for i := int64(0); i < int64(1000); i++ {
		TestRows = append(TestRows, Util.Row{
			Vals: []interface{}{i, i, float64(i), fmt.Sprintf("%v", i)},
		})
	}
}

var StudentMetadata = Util.Metadata{
	Catalog:     "TEST",
	Schema:      "DEFAULT",
	Table:       "STUDENT",
	ColumnNames: []string{"ID", "NAME"},
	ColumnTypes: []Util.Type{Util.INT64, Util.STRING},
}

var StudentRows = []Util.Row{
	Util.Row{Vals: []interface{}{int64(1), "a"}},
	Util.Row{Vals: []interface{}{int64(2), "b"}},
	Util.Row{Vals: []interface{}{int64(3), "c"}},
}

var ClassMetadata = Util.Metadata{
	Catalog:     "TEST",
	Schema:      "DEFAULT",
	Table:       "CLASS",
	ColumnNames: []string{"ID", "NAME", "TYPEID"},
	ColumnTypes: []Util.Type{Util.INT64, Util.STRING, Util.INT64},
}

var ClassRows = []Util.Row{
	Util.Row{Vals: []interface{}{int64(1), "physics", int64(1)}},
	Util.Row{Vals: []interface{}{int64(2), "math", int64(1)}},
	Util.Row{Vals: []interface{}{int64(4), "english", int64(2)}},
	Util.Row{Vals: []interface{}{int64(3), "algorithm", int64(1)}},
}

func NewTestCatalog(schema, table string) *TestCatalog {
	schema, table = strings.ToUpper(schema), strings.ToUpper(table)
	if len(TestRows) <= 0 {
		GenerateTestRows()
	}
	var res *TestCatalog
	switch table {
	case "STUDENT":
		res = &TestCatalog{
			Metadata: StudentMetadata,
			Rows:     StudentRows,
			Index:    0,
		}

	case "CLASS":
		res = &TestCatalog{
			Metadata: ClassMetadata,
			Rows:     ClassRows,
			Index:    0,
		}
	case "TEST":
		res = &TestCatalog{
			Metadata: TestMetadata,
			Rows:     TestRows,
			Index:    0,
		}
	default:
		res = &TestCatalog{
			Metadata: ClassMetadata,
			Rows:     ClassRows,
			Index:    0,
		}
	}

	res.Metadata.Reset()
	return res
}

func (self *TestCatalog) GetMetadata() *Util.Metadata {
	return &self.Metadata
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
