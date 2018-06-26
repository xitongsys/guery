package Split

import (
	"io"

	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Row"
)

const MAX_SPLIT_SIZE = 10000

type Split struct {
	Metadata  *Metadata.Metadata
	SplitKeys []interface{}

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

func (self *Split) GetValues(index int) []interface{} {
	res := make([]interface{}, self.GetColumnNumber())
	for i := 0; i < self.GetColumnNumber(); i++ {
		res[i] = self.Values[i][index]
	}
	return res
}

func (self *Split) GetKeys(index int) []interface{} {
	res := make([]interface{}, self.GetKeyColumnNumber())
	for i := 0; i < self.GetKeyColumnNumber(); i++ {
		res[i] = self.Keys[i][index]
	}
	return res
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

func (self *Split) GetSplitKeyString() string {
	res := ""
	for _, key := range self.SplitKeys {
		res += fmt.Sprintf("%v", key)
	}
	return res
}

func (self *Split) GetKeyString(index int) string {
	res := ""
	for i := 0; i < self.GetKeyColumnNumber(); i++ {
		res += fmt.Sprintf("%v", self.Keys[i][index])
	}
	return res
}

func (self *Split) AppendValues(vals []interface{}) {
	for i := 0; i < len(vals); i++ {
		self.Values[i] = append(self.Values[i], vals[i])
		if vals[i] == nil {
			self.ValueFlags[i] = append(self.ValueFlags[i], false)
		} else {
			self.ValueFlags[i] = append(self.ValueFlags[i], true)
		}
	}
	self.RowsNumber++
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

func (self *Split) AppendColumns(vals [][]interface{}, valFlags [][]bool) {
	self.Values = append(self.Values, vals...)
	self.ValueFlags = append(self.ValueFlags, valFlags...)
}

func (self *Split) AppendKeyColumns(keys [][]interface{}, keyFlags [][]bool) {
	self.Keys = append(self.Keys, keys...)
	self.KeyFlags = append(self.KeyFlags, keyFlags...)
}

func JoinSplit(md *Metadata, spL, spR *Split.Split) *Split.Split {
	leftCN, rightCN := spL.GetColumnNumber(), spR.GetColumnNumber()
	leftRN, rightRN := spL.GetRowsNumber(), spR.GetRowsNumber()
	res := NewSplit(md)
	for i := 0; i < leftRN || i < rightRN; i++ {
		for j := 0; j < leftCN; j++ {
			if i < leftRN {
				res.Values[j] = append(res.Values[j], spL.Values[j][i])
				res.ValueFlags[j] = append(res.ValueFlags[j], spL.ValueFlags[j][i])
			} else {
				res.Values[j] = append(res.Values[j], nil)
				res.ValueFlags[j] = append(res.ValueFlags[j], false)
			}
		}

		for j := leftCN; j < leftCN+rightCN; j++ {
			if i < rightRN {
				res.Values[j] = append(res.Values[j], spL.Values[j-leftCN][i])
				res.ValueFlags[j] = append(res.ValueFlags[j], spL.ValueFlags[j-leftCN][i])
			} else {
				res.Values[j] = append(res.Values[j], nil)
				res.ValueFlags[j] = append(res.ValueFlags[j], false)
			}
		}
	}
	return res
}
