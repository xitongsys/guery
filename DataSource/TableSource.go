package DataSource

import (
	"fmt"
)

type TableSource struct {
	Name            string
	ColumnNames     []string
	ColumnTypes     []string
	ColumnNameIndex map[string]int
	Vals            [][]interface{}
}
