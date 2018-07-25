package Row

import (
	"io"

	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Type"
)

type RowsGroup struct {
	Metadata   *Metadata.Metadata
	RowsNumber int
	Keys       [][]interface{}
	Vals       [][]interface{}
	Index      int
}

func NewRowsGroup(md *Metadata.Metadata) *RowsGroup {
	return &RowsGroup{
		Metadata:   md,
		RowsNumber: 0,
		Keys:       make([][]interface{}, md.GetKeyNumber()),
		Vals:       make([][]interface{}, md.GetColumnNumber()),
		Index:      0,
	}
}

func (self *RowsGroup) Read() (*Row, error) {
	if self.Index >= self.RowsNumber {
		return nil, io.EOF
	}
	row := NewRow()
	for i := 0; i < len(self.Vals); i++ {
		row.AppendVals(self.Vals[i][self.Index])
	}
	for i := 0; i < len(self.Keys); i++ {
		row.AppendKeys(self.Keys[i][self.Index])
	}
	self.Index++
	return row, nil
}

func (self *RowsGroup) Write(row *Row) {
	for i, v := range row.Vals {
		self.Vals[i] = append(self.Vals[i], v)
	}
	for i, v := range row.Keys {
		self.Keys[i] = append(self.Keys[i], v)
	}
	self.RowsNumber++
}

func (self *RowsGroup) AppendValRow(vals ...interface{}) {
	for i := 0; i < len(self.Vals); i++ {
		self.Vals[i] = append(self.Vals[i], vals[i])
	}
}

func (self *RowsGroup) AppendKeyRow(keys ...interface{}) {
	for i := 0; i < len(self.Keys); i++ {
		self.Keys[i] = append(self.Keys[i], keys[i])
	}
}

func (self *RowsGroup) AppendRowGroupRows(rg *RowsGroup) {
	for i := 0; i < len(self.Vals); i++ {
		self.Vals[i] = append(self.Vals[i], rg.Vals[i]...)
	}
	for i := 0; i < len(self.Keys); i++ {
		self.Keys[i] = append(self.Keys[i], rg.Keys[i]...)
	}
	self.RowsNumber += rg.GetRowsNumber()
}

func (self *RowsGroup) AppendRowGroupColumns(rg *RowsGroup) {
	for i := 0; i < rg.GetColumnsNumber(); i++ {
		self.Metadata.AppendColumn(rg.Metadata.Columns[i])
		self.Vals = append(self.Vals, rg.Vals[i])
	}
}

func (self *RowsGroup) GetRowVals(ri int) []interface{} {
	res := make([]interface{}, len(self.Vals))
	for i := 0; i < len(self.Vals); i++ {
		res[i] = self.Vals[i][ri]
	}
	return res
}

func (self *RowsGroup) GetRowKeys(ri int) []interface{} {
	res := make([]interface{}, len(self.Keys))
	for i := 0; i < len(self.Keys); i++ {
		res[i] = self.Keys[i][ri]
	}
	return res
}

func (self *RowsGroup) GetRow(ri int) *Row {
	res := RowPool.Get().(*Row)
	res.Clear()
	for i := 0; i < len(self.Vals); i++ {
		res.AppendVals(self.Vals[i][ri])
	}
	for i := 0; i < len(self.Keys); i++ {
		res.AppendKeys(self.Keys[i][ri])
	}
	return res
}

func (self *RowsGroup) Reset() {
	self.Index = 0
}

func (self *RowsGroup) GetIndex(name string) int {
	if i, ok := self.Metadata.ColumnMap[name]; ok {
		return i
	}
	return -1
}

func (self *RowsGroup) GetKeyString(index int) string {
	res := ""
	for _, ks := range self.Keys {
		res += Type.ToKeyString(ks[index]) + ":"
	}
	return res
}

func (self *RowsGroup) ClearRows() {
	self.Index = 0
	self.RowsNumber = 0
	for i := 0; i < len(self.Vals); i++ {
		self.Vals[i] = self.Vals[i][:0]
	}
	for i := 0; i < len(self.Keys); i++ {
		self.Keys[i] = self.Keys[i][:0]
	}
}

func (self *RowsGroup) ClearColumns() {
	self.Vals = [][]interface{}{}
}

func (self *RowsGroup) GetRowsNumber() int {
	return self.RowsNumber
}

func (self *RowsGroup) GetColumnsNumber() int {
	return len(self.Vals)
}

func (self *RowsGroup) GetKeysNumber() int {
	return len(self.Keys)
}

func (self *RowsGroup) AppendValColumns(cols ...[]interface{}) {
	self.Vals = append(self.Vals, cols...)
}

func (self *RowsGroup) AppendKeyColumns(keys ...[]interface{}) {
	self.Keys = append(self.Keys, keys...)
}

func (self *RowsGroup) SetColumn(index int, col []interface{}) {
	self.Vals[index] = col
}
