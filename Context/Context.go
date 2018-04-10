package Context

import (
	"github.com/xitongsys/guery/DataSource"
)

type Context struct {
	Tables map[string]DataSource.DataSource
}

func NewContext() *Context {
	res := &Context{
		Tables: make(map[string]DataSource.DataSource),
	}
	return res
}

func (self *Context) AddTable(name string, ds DataSource.DataSource) {
	self.Tables[name] = ds
}
