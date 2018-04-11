package Context

import (
	"github.com/xitongsys/guery/DataSource"
)

type Context struct {
	Tables       map[string]DataSource.DataSource
	TableRenames map[string]string
}

func NewContext() *Context {
	res := &Context{
		Tables:       make(map[string]DataSource.DataSource),
		TableRenames: make(map[string]string),
	}
	return res
}

func (self *Context) AddTable(name string, ds DataSource.DataSource) {
	self.Tables[name] = ds
	self.AddTableRename(name, name)
}

func (self *Context) AddTableRename(newName string, name string) {
	self.TableRenames[newName] = name
}
