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
	for i := 0; i < 100000; i++ {
		RowPool.Put(NewRow())
	}
}
