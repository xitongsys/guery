package Catalog

import (
	"io"
	"strings"

	"github.com/xitongsys/guery/Util"
)

type TestCatalog struct {
	Metadata Util.Metadata
	Rows     []Util.Row
	Index    int
}

var StudentMetadata = Util.Metadata{
	Catalog:     "TEST",
	Schema:      "DEFAULT",
	Table:       "STUDENT",
	ColumnNames: []string{"ID", "NAME"},
	ColumnTypes: []Util.Type{Util.INT, Util.STRING},
}

var StudentRows = []Util.Row{
	Util.Row{Vals: []interface{}{1, "a"}},
	Util.Row{Vals: []interface{}{2, "b"}},
	Util.Row{Vals: []interface{}{3, "c"}},
}

var ClassMetadata = Util.Metadata{
	Catalog:     "TEST",
	Schema:      "DEFAULT",
	Table:       "CLASS",
	ColumnNames: []string{"ID", "NAME", "TYPEID"},
	ColumnTypes: []Util.Type{Util.INT, Util.STRING, Util.INT},
}

var ClassRows = []Util.Row{
	Util.Row{Vals: []interface{}{1, "physics", 1}},
	Util.Row{Vals: []interface{}{2, "math", 1}},
	Util.Row{Vals: []interface{}{3, "algorithm", 1}},
}

func NewTestCatalog(schema, table string) *TestCatalog {
	schema, table = strings.ToUpper(schema), strings.ToUpper(table)
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
	case "DEFAULT":
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
	if self.Index >= len(self.Rows) {
		self.Index = 0
		return nil, io.EOF
	}

	self.Index++
	return &self.Rows[self.Index-1], nil
}
