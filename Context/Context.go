package Context

import (
	"github.com/xitongsys/guery/DataSource"
)

type Context struct {
	Tables map[string]DataSource.DataSource
}
