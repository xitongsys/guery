package Catalog

import (
	"strings"

	"github.com/xitongsys/guery/Util"
)

type Catalog interface {
	GetMetadata() *Util.Metadata
	ReadRow() (*Util.Row, error)
}

func NewCatalog(catalog string, schema string, table string) Catalog {
	catalog, schema, table = strings.ToUpper(catalog), strings.ToUpper(schema), strings.ToUpper(table)
	switch catalog {
	case "TEST":
		return NewTestCatalog(schema, table)

	}
	return nil
}
