package Split

import (
	"io"

	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Row"
)

const MAX_SPLIT_SIZE = 10000

type Split struct {
	Metadata *Metadata.Metadata

	RowsNumber           int
	Index                int
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

func (self *Split) GetKeyColumnNumber() int {
	return len(self.Keys)
}

func (self *Split) GetRowsNumber() int {
	return self.RowsNumber
}

func (self *Split) Append(sp *Split, indexes ...int) {
	if len(indexes) <= 0 {
		for i := 0; i < self.GetColumnNumber(); i++ {
			self.Values[i] = append(self.Values[i], sp.Values[i]...)
			self.ValueFlags[i] = append(self.ValueFlags[i], sp.ValueFlags[i]...)
		}
		for i := 0; i < self.GetKeyColumnNumber(); i++ {
			self.Keys[i] = append(self.Keys[i], sp.Keys[i]...)
			self.KeyFlags[i] = append(self.KeyFlags[i], sp.KeyFlags[i]...)
		}
		self.RowsNumber += sp.GetRowsNumber()

	} else {
		for _, j := range indexes {
			for i := 0; i < self.GetColumnNumber(); i++ {
				self.Values[i] = append(self.Values[i], sp.Values[i][j])
				self.ValueFlags[i] = append(self.ValueFlags[i], sp.ValueFlags[i][j])
			}
			for i := 0; i < self.GetKeyColumnNumber(); i++ {
				self.Keys[i] = append(self.Keys[i], sp.Keys[i][j])
				self.KeyFlags[i] = append(self.KeyFlags[i], sp.KeyFlags[i][j])
			}
		}
		self.RowsNumber += len(indexes)
	}
}

func (self *Split) ReadRow() (*Row.Row, error) {
	if self.Index >= self.RowsNumber {
		return nil, io.EOF
	}

	row := Row.NewRow()
	for i := 0; i < self.GetColumnNumber(); i++ {
		row.AppendVals(self.Values[i][self.Index])
	}
	self.Index++
	return row, nil
}
