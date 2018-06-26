package Split

import (
	"sort"

	"github.com/xitongsys/guery/Type"
)

func (self *Split) Less(a, b int) bool {
	for k := 0; k < len(self.OrderTypes); k++ {
		va, vb := self.Keys[k][a], self.Keys[k][b]
		if va == vb {
			continue
		}
		res := Type.LTFunc(va, vb).(bool)
		if self.OrderTypes[k] == Type.DESC {
			res = !res
		}
		return res
	}
	return false
}

func (self *Split) Swap(a, b int) {
	for i := 0; i < self.GetColumnNumber(); i++ {
		self.ValueFlags[i][a], self.ValueFlags[i][b] = self.ValueFlags[i][b], self.ValueFlags[i][a]
		self.Values[i][a], self.Values[i][b] = self.Values[i][b], self.Values[i][a]
	}

	for i := 0; i < self.GetKeyColumnNumber(); i++ {
		self.KeyFlags[i][a], self.KeyFlags[i][b] = self.KeyFlags[i][b], self.KeyFlags[i][a]
		self.Keys[i][a], self.Keys[i][b] = self.Keys[i][b], self.Keys[i][a]
	}
}

func (self *Split) Len() int {
	return self.GetRowsNumber()
}

func (self *Split) Sort() {
	sort.Sort(self)
}
