package Catalog

import (
	"github.com/xitongsys/guery/Util"
)

type Catalog interface {
	GetMetadata() *Metadata
	ReadRow() (*Util.Row, error)
}
