package Plan

import (
	"strings"

	"github.com/xitongsys/guery/DataSource"
)

type Context struct {
	DataSourceMap map[string]DataSource.DataSource
	Current       string
}

func (self *Context) GetDataValue(name string) interface{} {
	names := strings.Split(name, ".")
	if len(names) == 1 {
		return self.DataSourceMap[self.Current].ReadColumn(names[0])
	} else {
		return self.DataSourceMap[names[0]].ReadColumn(names[1])
	}
}
