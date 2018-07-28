package row

import (
	"io"

	"github.com/xitongsys/guery/gtype"
	"github.com/xitongsys/guery/metadata"
)

//RowsGroup to store rows
type RowsGroup struct {
	Metadata   *metadata.Metadata
	RowsNumber int
	Keys       [][]interface{}
	Vals       [][]interface{}
	Index      int
}

func NewRowsGroup(md *metadata.Metadata) *RowsGroup {
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
	r := NewRow()
	for i := 0; i < len(self.Vals); i++ {
		r.AppendVals(self.Vals[i][self.Index])
	}
	for i := 0; i < len(self.Keys); i++ {
		r.AppendKeys(self.Keys[i][self.Index])
	}
	self.Index++
	return r, nil
}

func (self *RowsGroup) Write(r *Row) {
	for i, v := range r.Vals {
		self.Vals[i] = append(self.Vals[i], v)
	}
	for i, v := range r.Keys {
		self.Keys[i] = append(self.Keys[i], v)
	}
	self.RowsNumber++
}

//Append a row of values. ColumnNumber must be same.
func (self *RowsGroup) AppendRowVals(vals ...interface{}) {
	for i := 0; i < len(self.Vals); i++ {
		self.Vals[i] = append(self.Vals[i], vals[i])
	}
	self.RowsNumber += 1
}

//Append a row of values. KeyColumnNumber must be same.
func (self *RowsGroup) AppendRowKeys(keys ...interface{}) {
	for i := 0; i < len(self.Keys); i++ {
		self.Keys[i] = append(self.Keys[i], keys[i])
	}
}

//Append rows from RowsGroup. ColumnNumber muse be same.
func (self *RowsGroup) AppendRowGroupRows(rg *RowsGroup) {
	for i := 0; i < len(self.Vals); i++ {
		self.Vals[i] = append(self.Vals[i], rg.Vals[i]...)
	}
	for i := 0; i < len(self.Keys); i++ {
		self.Keys[i] = append(self.Keys[i], rg.Keys[i]...)
	}
	self.RowsNumber += rg.GetRowsNumber()
}

//Append columns from RowsGroup. RowsNumber must be same.
func (self *RowsGroup) AppendRowGroupColumns(rg *RowsGroup) {
	self.RowsNumber = rg.GetRowsNumber()
	for i := 0; i < rg.GetColumnsNumber(); i++ {
		self.Metadata.AppendColumn(rg.Metadata.Columns[i])
		self.Vals = append(self.Vals, rg.Vals[i])
	}
}

//Get values of ri row
func (self *RowsGroup) GetRowVals(ri int) []interface{} {
	res := make([]interface{}, len(self.Vals))
	for i := 0; i < len(self.Vals); i++ {
		res[i] = self.Vals[i][ri]
	}
	return res
}

//Get keys of ri row
func (self *RowsGroup) GetRowKeys(ri int) []interface{} {
	res := make([]interface{}, len(self.Keys))
	for i := 0; i < len(self.Keys); i++ {
		res[i] = self.Keys[i][ri]
	}
	return res
}

//Get ri row
func (self *RowsGroup) GetRow(ri int) *Row {
	if ri >= self.RowsNumber {
		return nil
	}

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

//Reset index
func (self *RowsGroup) ResetIndex() {
	self.Index = 0
}

//Get column index by column name
func (self *RowsGroup) GetColumnIndex(name string) int {
	if i, ok := self.Metadata.ColumnMap[name]; ok {
		return i
	}
	return -1
}

//Get row key string of index. Key string is key1:key2:key3...
func (self *RowsGroup) GetKeyString(index int) string {
	res := ""
	for _, ks := range self.Keys {
		res += gtype.ToKeyString(ks[index]) + ":"
	}
	return res
}

//Clear rows
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

//Clear columns
func (self *RowsGroup) ClearColumns() {
	self.Index = 0
	self.RowsNumber = 0
	self.Vals = [][]interface{}{}
}

func (self *RowsGroup) GetRowsNumber() int {
	return self.RowsNumber
}

func (self *RowsGroup) GetColumnsNumber() int {
	return len(self.Vals)
}

func (self *RowsGroup) GetKeyColumnsNumber() int {
	return len(self.Keys)
}

//Append column values. RowsNumber must be same.
func (self *RowsGroup) AppendValColumns(cols ...[]interface{}) {
	for _, col := range cols {
		self.Vals = append(self.Vals, col)
		self.RowsNumber = len(col)
	}
}

//Append key columns. RowsNumber must be same.
func (self *RowsGroup) AppendKeyColumns(keys ...[]interface{}) {
	self.Keys = append(self.Keys, keys...)
}

func (self *RowsGroup) SetColumn(index int, col []interface{}) {
	self.Vals[index] = col
}
