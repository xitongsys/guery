package Util

import (
	"sort"
)

type Rows struct {
	Data  []*Row
	Order []OrderType
}

func NewRows() *Rows {
	return &Rows{
		Data:  []*Row{},
		Order: []OrderType{},
	}
}

func (self *Rows) Less(i, j int) bool {
	return false
}

func (self Rows) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func (self Rows) Len() int {
	return len(self)
}

func (self Rows) SortDesc() {
	sort.Reverse(self)
}

func (self Rows) SortASC() {
	sort.Sort(self)
}
