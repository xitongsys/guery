package Util

import (
	"fmt"
)

//go:generate msgp

type Row struct {
	Keys []interface{}
	Vals []interface{}
}

func NewRow(vals ...interface{}) *Row {
	res := &Row{}
	res.Vals = append(res.Vals, vals...)
	return res
}

func (self *Row) GetKeyString() string {
	res := ""
	for _, key := range self.Keys {
		res += fmt.Sprintf("%v", key)
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
