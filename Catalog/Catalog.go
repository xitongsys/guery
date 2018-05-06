package Catalog

import (
	"github.com/xitongsys/guery/Util"
)

type Catalog interface {
	GetMetadata() *Util.Metadata
	ReadRow() (*Util.Row, error)
}
