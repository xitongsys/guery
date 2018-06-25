package Split

import (
	"github.com/xitongsys/guery/Metadata"
)

const MAX_SPLIT_SIZE = 10000

type Split struct {
	Metadata *Metadata.Metadata

	RowsNumber           int
	Values, Keys         [][]interface{}
	ValueFlags, KeyFlags [][]bool //false:nil; true:not nil;
}

func NewSplit(md *Metadata.Metadata) *Split {
	colNum := md.GetColumnNumber()
	return &Split{
		Metadata:   md,
		Values:     make([][]interface{}, colNum),
		Keys:       make([][]interface{}, colNum),
		ValueFlags: make([][]bool, colNum),
		KeyFlags:   make([][]bool, colNum),
	}
}

func (self *Split) GetColumnNumber() int {
	return self.Metadata.GetColumnNumber()
}

func (self *Split) GetRowsNumber() int {
	return self.RowsNumber
}

func (self *Split) Merge(sa *Split) *Split {
	self.RowsNumber += sa.RowsNumber

	for i := 0; i < len(self.Values); i++ {
		self.Values[i] = append(self.Values[i], sa.Values[i]...)
		self.ValueFlags[i] = append(self.ValueFlags, sa.ValueFlags[i]...)
	}
	for i := 0; i < len(self.Keys); i++ {
		self.Keys[i] = append(self.Keys[i], sa.Keys[i]...)
		self.KeyFlags[i] = append(self.KeyFlags[i], sa.KeyFlags[i]...)
	}
	return self
}
