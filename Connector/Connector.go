package Connector

import (
	"fmt"
	"strings"

	"github.com/xitongsys/guery/Util"
)

type Connector interface {
	GetMetadata() *Util.Metadata
	ReadRow() (*Util.Row, error)
	ReadRowByColumns(colIndexes []int) (*Util.Row, error)
	SkipRows(num int64)
}

func NewConnector(catalog string, schema string, table string) (Connector, error) {
	catalog, schema, table = strings.ToUpper(catalog), strings.ToUpper(schema), strings.ToUpper(table)
	switch catalog {
	case "TEST":
		return NewTestConnector(schema, table)
	case "LOCALFILE":
		return NewLocalFileConnector(schema, table)

	}
	return nil, fmt.Errorf("NewConnector failed")
}
