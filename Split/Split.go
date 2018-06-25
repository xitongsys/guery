package Split

import (
	"github.com/xitongsys/guery/Metadata"
)

type Split struct {
	Metadata *Metadata.Metadata

	RowsNumber           int
	Values, Keys         [][]interface{}
	ValueFlags, KeyFlags [][]int //number of nils
}
