package Util

import (
	"sort"
)

type Row struct {
	Keys []interface{}
	Vals []interface{}
}

func NewRow(vals ...interface{}) *Row {
	res := &Row{}
	return res
}

func (self *Row) AppendVals(vals ...interface{}) *Row {
	self.Vals = append(self.Vals, vals...)
	return self
}

func (self *Row) AppendKeys(keys ...interface{}) *Row {
	self.Keys = append(self.Keys, keys...)
	return self
}

func (self *Row) SetKeys(indexes []int) {
	if indexes == nil {
		return
	}
	keys := []interface{}{}
	sort.Ints(indexes)
	for _, index := range indexes {
		if index >= len(self.Vals) {
			break
		}
		keys = append(keys, self.Vals[index])
	}
	self.Keys = keys
}
