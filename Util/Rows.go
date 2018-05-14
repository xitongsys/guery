package Util

import (
	"sort"
)

type Rows []*Row

func NewRows() *Rows {
	return &[]*Row{}
}

func (self *Rows) Append(row *Row) {
	self = append(self, row)
}

func (self *Rows) Less(i, j int) bool { return self[i].Key < self[i].Key }
func (self *Rows) Swap(i, j int)      { self[i], self[j] = self[j], self[i] }
func (self *Rows) Len() int           { return len(self) }

func (self *Rows) SortDesc() {
	sort.Reverse(self)
}

func (self *Rows) SortASC() {
	sort.Sort(self)
}
