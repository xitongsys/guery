package Row

import (
	"fmt"
	"io"

	"github.com/xitongsys/guery/Metadata"
)

type RowsGroup struct {
	Metadata  *Metadata.Metadata
	RowNumber int
	Keys      [][]interface{}
	Vals      [][]interface{}
	Index     int
}

func NewRowsGroup(md *Metadata.Metadata) *RowsGroup {
	return &RowsGroup{
		Metadata:  md,
		RowNumber: 0,
		Keys:      make([][]interface{}, md.GetColumnNumber()),
		Vals:      make([][]interface{}, md.GetKeyNumber()),
		Index:     0,
	}
}

func (self *RowsGroup) Read() (*Row, error) {
	if self.Index >= self.RowNumber {
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
	self.RowNum++
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
		res += fmt.Sprintf("%v:", ks[index])
	}
	return res
}

func (self *RowsGroup) ClearRows() {
	self.Index = 0
	self.RowNumber = 0
	for i := 0; i < len(self.Vals); i++ {
		self.Vals[i] = self.Vals[i][:0]
	}
	for i := 0; i < len(self.Keys); i++ {
		self.Keys[i] = self.Keys[i][:0]
	}
}

func (self *RowsGroup) GetRowsNum() int {
	return self.RowNumber
}
