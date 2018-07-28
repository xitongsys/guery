package row

import (
	"sync"
)

var RowPool *sync.Pool
var RowsGroupPool *sync.Pool

func init() {
	RowPool = &sync.Pool{
		New: func() interface{} {
			return NewRow()
		},
	}
}
