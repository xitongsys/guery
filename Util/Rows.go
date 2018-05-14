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
	rowi, rowj := self.Data[i], self.Data[j]
	for k := 0; k < len(self.Order); k++ {
		vi, vj := rowi.Vals[k], rowj.Vals[k]
		if vi == vj {
			continue
		}
		res := Less(vi, vj)
		if self.Order[k] == DESC {
			res = !res
		}
		return res
	}
	return false
}

func (self *Rows) Swap(i, j int) {
	self.Data[i], self.Data[j] = self.Data[j], self.Data[i]
}

func (self *Rows) Len() int {
	return len(self.Data)
}

func (self *Rows) Sort() {
	sort.Sort(self)
}

func (self *Rows) Append(row *Row) {
	self.Data = append(self.Data, row)
}
