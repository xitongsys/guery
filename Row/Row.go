package Row

import (
	"fmt"
)

type Row struct {
	Keys []interface{}
	Vals []interface{}
}

func NewRow(vals ...interface{}) *Row {
	colNum := 0
	if vals != nil {
		colNum = len(vals)
	}
	res := &Row{
		Keys: []interface{}{},
		Vals: make([]interface{}, colNum),
	}
	for i := 0; i < colNum; i++ {
		res.Vals[i] = vals[i]
	}
	return res
}

func (self *Row) GetKeyString() string {
	res := ""
	if self.Keys == nil {
		self.Keys = []interface{}{}
	}
	for _, key := range self.Keys {
		res += fmt.Sprintf("%v:", key)
	}
	return res
}

func (self *Row) AppendKeys(keys ...interface{}) *Row {
	self.Keys = append(self.Keys, keys...)
	return self
}

func (self *Row) AppendVals(vals ...interface{}) *Row {
	self.Vals = append(self.Vals, vals...)
	return self
}

func (self *Row) AppendRow(row *Row) *Row {
	self.Vals = append(self.Vals, row.Vals...)
	return self
}

func (self *Row) ClearKeys() {
	self.Keys = []interface{}{}
}
