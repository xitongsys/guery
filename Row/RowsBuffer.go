package Row

import (
	"bytes"
	"io"
	"sync"

	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Type"
	"github.com/xitongsys/guery/Util"
)

const ROWS_BUFFER_SIZE = 100000

type RowsBuffer struct {
	sync.Mutex
	MD         *Metadata.Metadata
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

func NewRowsBuffer(md *Metadata.Metadata, reader io.Reader, writer io.Writer) *RowsBuffer {
	res := &RowsBuffer{
		MD:         md,
		BufferSize: ROWS_BUFFER_SIZE,
		Reader:     reader,
		Writer:     writer,
	}
	res.ClearValues()
	return res
}

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

func (self *RowsBuffer) Flush() error {
	self.writeRows()
	if err := Util.WriteEOFMessage(self.Writer); err != nil {
		return err
	}
	return nil
}

func (self *RowsBuffer) writeRows() error {
	self.Lock()
	defer self.Unlock()
	defer self.ClearValues()
	ln := len(self.ValueBuffers)

	//for 0 cols, just need send the number of rows
	if ln <= 0 {
		buf := Type.EncodeValues([]interface{}{int64(self.RowsNumber)}, Type.INT64)
		return Util.WriteMessage(self.Writer, buf)
	}

	//for several cols
	for i := 0; i < ln; i++ {
		col := self.ValueNilFlags[i]
		buf := Type.EncodeBool(col)
		if err := Util.WriteMessage(self.Writer, buf); err != nil {
			return err
		}

		col = self.ValueBuffers[i]
		t, err := self.MD.GetTypeByIndex(i)
		if err != nil {
			return err
		}
		buf = Type.EncodeValues(col, t)
		if err := Util.WriteMessage(self.Writer, buf); err != nil {
			return err
		}
	}

	ln = len(self.KeyBuffers)
	for i := 0; i < ln; i++ {
		col := self.KeyNilFlags[i]
		buf := Type.EncodeBool(col)
		if err := Util.WriteMessage(self.Writer, buf); err != nil {
			return err
		}

		col = self.KeyBuffers[i]
		t, err := self.MD.GetKeyTypeByIndex(i)
		if err != nil {
			return err
		}
		buf = Type.EncodeValues(col, t)
		if err := Util.WriteMessage(self.Writer, buf); err != nil {
			return err
		}
	}
	return nil
}

func (self *RowsBuffer) readRows() error {
	self.Lock()
	defer self.Unlock()
	defer func() {
		self.Index = 0
	}()

	colNum := self.MD.GetColumnNumber()
	//for 0 cols
	if colNum <= 0 {
		buf, err := Util.ReadMessage(self.Reader)
		if err != nil {
			return err
		}
		vals, err := Type.DecodeINT64(bytes.NewReader(buf))
		if err != nil || len(vals) <= 0 {
			return err
		}
		self.RowsNumber = int(vals[0].(int64))
	}

	//for cols
	for i := 0; i < colNum; i++ {
		buf, err := Util.ReadMessage(self.Reader)
		if err != nil {
			return err
		}

		self.ValueNilFlags[i], err = Type.DecodeBOOL(bytes.NewReader(buf))
		if err != nil {
			return err
		}

		buf, err = Util.ReadMessage(self.Reader)
		t, err := self.MD.GetTypeByIndex(i)
		if err != nil {
			return err
		}
		values, err := Type.DecodeValue(bytes.NewReader(buf), t)
		if err != nil {
			return err
		}

		//log.Println("=======", buf, values, self.ValueNilFlags)

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
		//log.Println("=======", buf, values, self.ValueBuffers)

		self.RowsNumber = len(self.ValueNilFlags[i])

	}

	keyNum := self.MD.GetKeyNumber()
	for i := 0; i < keyNum; i++ {
		buf, err := Util.ReadMessage(self.Reader)
		if err != nil {
			return err
		}
		self.KeyNilFlags[i], err = Type.DecodeBOOL(bytes.NewReader(buf))
		if err != nil {
			return err
		}

		buf, err = Util.ReadMessage(self.Reader)
		t, err := self.MD.GetKeyTypeByIndex(i)
		if err != nil {
			return err
		}
		keys, err := Type.DecodeValue(bytes.NewReader(buf), t)
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
		//log.Println("=======", buf, keys, self.KeyBuffers, keyNum)
	}

	return nil

}

func (self *RowsBuffer) WriteRow(row *Row) error {
	self.Lock()
	defer self.Unlock()

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
	return nil
}

func (self *RowsBuffer) ReadRow() (*Row, error) {
	self.Lock()
	defer self.Unlock()

	for self.Index >= self.RowsNumber {
		self.ClearValues()
		if err := self.readRows(); err != nil {
			return nil, err
		}
	}

	row := NewRow()
	for _, col := range self.ValueBuffers {
		row.AppendVals(col[self.Index])
	}
	for _, col := range self.KeyBuffers {
		row.AppendKeys(col[self.Index])
	}
	self.Index++
	return row, nil
}
