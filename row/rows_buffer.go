package row

import (
	"bytes"
	"io"
	"sync"

	"github.com/xitongsys/guery/gtype"
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/util"
)

const ROWS_BUFFER_SIZE = 1000

//Buffer for read/write rows
type RowsBuffer struct {
	sync.Mutex
	MD         *metadata.Metadata
	BufferSize int
	RowsNumber int
	Index      int

	ValueBuffers  [][]interface{}
	ValueNilFlags [][]interface{} //bool

	KeyBuffers  [][]interface{}
	KeyNilFlags [][]interface{} //bool

	Reader io.Reader
	Writer io.Writer
}

//New RowsBuffer for read or write
func NewRowsBuffer(md *metadata.Metadata, reader io.Reader, writer io.Writer) *RowsBuffer {
	res := &RowsBuffer{
		MD:         md,
		BufferSize: ROWS_BUFFER_SIZE,
		Reader:     reader,
		Writer:     writer,
	}
	res.ClearValues()
	return res
}

//Clear buffers
func (self *RowsBuffer) ClearValues() {
	colNum := self.MD.GetColumnNumber()
	self.ValueBuffers = make([][]interface{}, colNum)
	self.ValueNilFlags = make([][]interface{}, colNum)

	keyNum := self.MD.GetKeyNumber()
	self.KeyBuffers = make([][]interface{}, keyNum)
	self.KeyNilFlags = make([][]interface{}, keyNum)
	self.Index = 0
	self.RowsNumber = 0

}

//Flush
func (self *RowsBuffer) Flush() error {
	self.Lock()
	defer self.Unlock()
	if err := self.writeRows(); err != nil {
		return err
	}
	if err := util.WriteEOFMessage(self.Writer); err != nil {
		return err
	}
	return nil
}

//Write rows
func (self *RowsBuffer) writeRows() error {
	defer self.ClearValues()
	ln := len(self.ValueBuffers)

	//for 0 cols, just need send the number of rows
	if ln <= 0 {
		buf := gtype.EncodeValues([]interface{}{int64(self.RowsNumber)}, gtype.INT64)
		return util.WriteMessage(self.Writer, buf)
	}

	//for several cols
	for i := 0; i < ln; i++ {
		col := self.ValueNilFlags[i]
		buf := gtype.EncodeBool(col)
		if err := util.WriteMessage(self.Writer, buf); err != nil {
			return err
		}

		col = self.ValueBuffers[i]
		t, err := self.MD.GetTypeByIndex(i)
		if err != nil {
			return err
		}
		buf = gtype.EncodeValues(col, t)
		if err := util.WriteMessage(self.Writer, buf); err != nil {
			return err
		}
	}

	ln = len(self.KeyBuffers)
	for i := 0; i < ln; i++ {
		col := self.KeyNilFlags[i]
		buf := gtype.EncodeBool(col)
		if err := util.WriteMessage(self.Writer, buf); err != nil {
			return err
		}

		col = self.KeyBuffers[i]
		t, err := self.MD.GetKeyTypeByIndex(i)
		if err != nil {
			return err
		}
		buf = gtype.EncodeValues(col, t)
		if err := util.WriteMessage(self.Writer, buf); err != nil {
			return err
		}
	}
	return nil
}

//read rows
func (self *RowsBuffer) readRows() error {
	defer func() {
		self.Index = 0
	}()

	colNum := self.MD.GetColumnNumber()
	//for 0 cols
	if colNum <= 0 {
		buf, err := util.ReadMessage(self.Reader)
		if err != nil {
			return err
		}
		vals, err := gtype.DecodeINT64(bytes.NewReader(buf))
		if err != nil || len(vals) <= 0 {
			return err
		}
		self.RowsNumber = int(vals[0].(int64))
	}

	//for cols
	for i := 0; i < colNum; i++ {
		buf, err := util.ReadMessage(self.Reader)
		if err != nil {
			return err
		}

		self.ValueNilFlags[i], err = gtype.DecodeBOOL(bytes.NewReader(buf))
		if err != nil {
			return err
		}

		buf, err = util.ReadMessage(self.Reader)
		t, err := self.MD.GetTypeByIndex(i)
		if err != nil {
			return err
		}
		values, err := gtype.DecodeValue(bytes.NewReader(buf), t)
		if err != nil {
			return err
		}

		self.ValueBuffers[i] = make([]interface{}, len(self.ValueNilFlags[i]))
		k := 0
		for j := 0; j < len(self.ValueNilFlags[i]) && k < len(values); j++ {
			if self.ValueNilFlags[i][j].(bool) {
				self.ValueBuffers[i][j] = values[k]
				k++
			} else {
				self.ValueBuffers[i][j] = nil
			}
		}

		self.RowsNumber = len(self.ValueNilFlags[i])

	}

	keyNum := self.MD.GetKeyNumber()
	for i := 0; i < keyNum; i++ {
		buf, err := util.ReadMessage(self.Reader)
		if err != nil {
			return err
		}
		self.KeyNilFlags[i], err = gtype.DecodeBOOL(bytes.NewReader(buf))
		if err != nil {
			return err
		}

		buf, err = util.ReadMessage(self.Reader)
		t, err := self.MD.GetKeyTypeByIndex(i)
		if err != nil {
			return err
		}
		keys, err := gtype.DecodeValue(bytes.NewReader(buf), t)
		if err != nil {
			return err
		}

		self.KeyBuffers[i] = make([]interface{}, len(self.KeyNilFlags[i]))
		k := 0
		for j := 0; j < len(self.KeyNilFlags[i]) && k < len(keys); j++ {
			if self.KeyNilFlags[i][j].(bool) {
				self.KeyBuffers[i][j] = keys[k]
				k++
			} else {
				self.KeyBuffers[i][j] = nil
			}
		}
	}

	return nil

}

//Write rows to RowsBuffer
func (self *RowsBuffer) WriteRow(rows ...*Row) error {
	self.Lock()
	defer self.Unlock()
	for _, row := range rows {
		for i, val := range row.Vals {
			if val != nil {
				self.ValueBuffers[i] = append(self.ValueBuffers[i], val)
				self.ValueNilFlags[i] = append(self.ValueNilFlags[i], true)
			} else {
				self.ValueNilFlags[i] = append(self.ValueNilFlags[i], false)
			}
		}

		for i, key := range row.Keys {
			if key != nil {
				self.KeyBuffers[i] = append(self.KeyBuffers[i], key)
				self.KeyNilFlags[i] = append(self.KeyNilFlags[i], true)
			} else {
				self.KeyNilFlags[i] = append(self.KeyNilFlags[i], false)
			}
		}
		self.RowsNumber++

		if self.RowsNumber >= self.BufferSize {
			if err := self.writeRows(); err != nil {
				return err
			}
		}
	}
	return nil
}

//Read one row from RowsBuffer
func (self *RowsBuffer) ReadRow() (*Row, error) {
	self.Lock()
	defer self.Unlock()

	for self.Index >= self.RowsNumber {
		self.ClearValues()
		if err := self.readRows(); err != nil {
			return nil, err
		}
	}

	row := RowPool.Get().(*Row)
	row.Clear()
	row.Vals = make([]interface{}, len(self.ValueBuffers))
	for i, col := range self.ValueBuffers {
		row.Vals[i] = col[self.Index]
	}

	row.Keys = make([]interface{}, len(self.KeyBuffers))
	for i, col := range self.KeyBuffers {
		row.Keys[i] = col[self.Index]
	}
	self.Index++
	return row, nil
}

//Write one RowsGroup to RowsBuffer
func (self *RowsBuffer) Write(rg *RowsGroup) error {
	self.Lock()
	defer self.Unlock()
	for i, vs := range rg.Vals {
		for _, v := range vs {
			if v != nil {
				self.ValueBuffers[i] = append(self.ValueBuffers[i], v)
				self.ValueNilFlags[i] = append(self.ValueNilFlags[i], true)
			} else {
				self.ValueNilFlags[i] = append(self.ValueNilFlags[i], false)
			}
		}
	}

	for i, ks := range rg.Keys {
		for _, k := range ks {
			if k != nil {
				self.KeyBuffers[i] = append(self.KeyBuffers[i], k)
				self.KeyNilFlags[i] = append(self.KeyNilFlags[i], true)
			} else {
				self.KeyNilFlags[i] = append(self.KeyNilFlags[i], false)
			}
		}
	}
	self.RowsNumber += rg.RowsNumber

	if self.RowsNumber >= self.BufferSize {
		if err := self.writeRows(); err != nil {
			return err
		}
	}
	return nil
}

//Read one RowsGroup from RowsBuffer
func (self *RowsBuffer) Read() (*RowsGroup, error) {
	self.Lock()
	defer self.Unlock()

	for self.Index >= self.RowsNumber {
		self.ClearValues()
		if err := self.readRows(); err != nil {
			return nil, err
		}
	}

	rg := NewRowsGroup(self.MD)
	readSize := self.BufferSize
	if readSize > self.RowsNumber-self.Index {
		readSize = self.RowsNumber - self.Index
	}

	for i := 0; i < len(rg.Vals); i++ {
		rg.Vals[i] = append(rg.Vals[i], self.ValueBuffers[i][self.Index:self.Index+readSize]...)
	}

	for i := 0; i < len(rg.Keys); i++ {
		rg.Keys[i] = append(rg.Keys[i], self.KeyBuffers[i][self.Index:self.Index+readSize]...)
	}

	self.Index += readSize
	rg.RowsNumber = readSize

	return rg, nil
}
