package Util

import (
	"io"
)

type RowsBuffer struct {
	MD         *Metadata
	BufferSize int64
	RowsNumber int
	Index      int

	ValueBuffers  [][]interface{}
	ValueNilFlags [][]interface{} //bool

	KeyBuffers  [][]interface{}
	KeyNilFlags [][]interface{} //bool

	Reader io.Reader
	Writer io.Writer
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

func (self *RowsBuffer) writeRows() error {
	ln := len(self.ValueBuffers)
	for i := 0; i < ln; i++ {
		col := self.ValueNilFlags[i]
		buf := CompressGzip(EncodeBool(col))
		if err := WriteMessage(self.Writer, buf); err != nil {
			return err
		}

		col = self.ValueBuffers[i]
		t, err := self.MD.GetTypeByIndex(i)
		if err != nil {
			return err
		}
		buf = CompressGzip(EncodeValues(col, t))
		if err := WriteMessage(self.Writer, buf); err != nil {
			return err
		}
	}

	ln = len(self.KeyBuffers)
	for i := 0; i < ln; i++ {
		col := self.KeyNilFlags[i]
		buf := CompressGzip(EncodeBool(col))
		if err := WriteMessage(self.Writer, buf); err != nil {
			return err
		}

		col = self.KeyBuffers[i]
		t, err := self.MD.GetKeyTypeByIndex(i)
		if err != nil {
			return err
		}
		buf = CompressGzip(EncodeValues(col, t))
		if err := WriteMessage(self.Writer, buf); err != nil {
			return err
		}
	}
	self.ClearValues()
	return nil
}

func (self *RowsBuffer) readRows() error {
	colNum := self.MD.GetColumnNumber()
	for i := 0; i < colNum; i++ {
		msg, err := ReadMessage(self.Reader)
		if err != nil {
			return err
		}
		buf, err := UncompressGzip(msg)
		if err != nil {
			return err
		}
	}
	return nil

}

func (self *RowsBuffer) WriteRow(row *Row) error {
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
	return nil
}

func (self *RowsBuffer) ReadRow() (*Row, error) {
	return nil, nil
}
