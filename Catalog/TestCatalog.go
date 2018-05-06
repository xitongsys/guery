package Catalog

import (
	"io"
	"strings"

	"github.com/xitongsys/guery/Util"
)

type TestCatalog struct {
	Metadata Metadata
	Rows     []Util.Row
	Index    int
}

var StudentMetadata = Metadata{
	Name:        "STUDENT",
	ColumnNames: []string{"ID", "NAME"},
	ColumnTypes: []ColumnType{INT, STRING},
}

var StudentRows = []Util.Row{
	Util.Row{Vals: []interface{}{1, "a"}},
	Util.Row{Vals: []interface{}{2, "b"}},
	Util.Row{Vals: []interface{}{3, "c"}},
}

var ClassMetadata = Metadata{
	Name:        "CLASS",
	ColumnNames: []string{"ID", "NAME"},
	ColumnTypes: []ColumnType{INT, STRING},
}

var ClassRows = []Util.Row{
	Util.Row{Vals: []interface{}{1, "physics"}},
	Util.Row{Vals: []interface{}{2, "math"}},
	Util.Row{Vals: []interface{}{3, "algorithm"}},
}

func NewTestCatalog(name string) *TestCatalog {
	name = strings.ToUpper(name)
	var res *TestCatalog
	switch name {
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
	}

	res.Metadata.Reset()
	return res
}

func (self *TestCatalog) GetMetadata() *Metadata {
	return &self.Metadata
}

func (self *TestCatalog) ReadRow() (*Util.Row, error) {
	if self.Index >= len(self.Rows) {
		return nil, io.EOF
	}

	self.Index++
	return &self.Rows[self.Index-1], nil
}
