package Context

import (
	"github.com/xitongsys/guery/DataSource"
)

type Context struct {
	Tables map[name]DataSource.DataSource
}
