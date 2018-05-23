package Connector

import (
	"fmt"
	"strings"

	"github.com/xitongsys/guery/Util"
)

type Connector interface {
	GetMetadata() *Util.Metadata
	Read() (*Util.Row, error)
	ReadByColumns(colIndexes []int) (*Util.Row, error)
}

func NewConnector(catalog string, schema string, table string) (Connector, error) {
	catalog, schema, table = strings.ToUpper(catalog), strings.ToUpper(schema), strings.ToUpper(table)
	switch catalog {
	case "TEST":
		return NewTestConnector(schema, table)
	case "FILE":
		return NewFileConnector(schema, table)

	}
	return nil, fmt.Errorf("NewConnector failed")
}
