package Split

import (
	"bytes"
	"io"
	"sync"

	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Split"
	"github.com/xitongsys/guery/Type"
	"github.com/xitongsys/guery/Util"
)

type SplitBuffer struct {
	sync.Mutex
	Metadata *Metadata.Metadata
	SP       *Split.Split

	Reader io.Reader
	Writer io.Writer
}

func NewSplitBuffer(md *Metadata.Metadata, reader io.Reader, writer io.Writer) *SplitBuffer {
	res := &RowsBuffer{
		Metadata: md,
		SP:       Split.NewSplit(md),
		Reader:   reader,
		Writer:   writer,
	}
	return res
}

func (self *SplitBuffer) Flush() error {
	err := self.FlushSplit(self.SP)
	self.SP = Split.NewSplit(self.Metadata)
	return err
}

func (self *SplitBuffer) FlushSplit(sp *Split.Split) error {
	colNum := sp.GetColumnNumber()
	//for 0 cols, just need send the number of rows
	if colNum <= 0 {
		buf := Type.EncodeValues([]interface{}{int64(self.RowsNumber)}, Type.INT64)
		return Util.WriteMessage(self.Writer, buf)
	}

	//for several cols
	for i := 0; i < colNum; i++ {
		col := sp.ValueFlags[i]
		buf := Type.EncodeBool(col)
		if err := Util.WriteMessage(self.Writer, buf); err != nil {
			return err
		}

		col = sp.Values[i]
		t, err := self.MD.GetTypeByIndex(i)
		if err != nil {
			return err
		}
		buf = Type.EncodeValues(col, t)
		if err := Util.WriteMessage(self.Writer, buf); err != nil {
			return err
		}
	}

	colNum = sp.GetKeyColumnNumber()
	for i := 0; i < colNum; i++ {
		col := self.KeyFlags[i]
		buf := Type.EncodeBool(col)
		if err := Util.WriteMessage(self.Writer, buf); err != nil {
			return err
		}

		col = self.Keys[i]
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

func (self *SplitBuffer) ReadSplit() (*Split.Split, error) {
	sp := Split.NewSplit(self.Metadata)
	colNum := self.MD.GetColumnNumber()
	//for 0 cols
	if colNum <= 0 {
		buf, err := Util.ReadMessage(self.Reader)
		if err != nil {
			return sp, err
		}
		vals, err := Type.DecodeINT64(bytes.NewReader(buf))
		if err != nil || len(vals) <= 0 {
			return sp, err
		}
		sp.RowsNumber = int(vals[0].(int64))
	}

	//for cols
	for i := 0; i < colNum; i++ {
		buf, err := Util.ReadMessage(self.Reader)
		if err != nil {
			return sp, err
		}

		sp.ValueFlags[i], err = Type.DecodeBOOL(bytes.NewReader(buf))
		if err != nil {
			return sp, err
		}

		buf, err = Util.ReadMessage(self.Reader)
		t, err := self.Metadata.GetTypeByIndex(i)
		if err != nil {
			return sp, err
		}
		values, err := Type.DecodeValue(bytes.NewReader(buf), t)
		if err != nil {
			return sp, err
		}

		//log.Println("=======", buf, values, self.ValueNilFlags)

		sp.Values[i] = make([]interface{}, len(self.ValueNilFlags[i]))
		k := 0
		for j := 0; j < len(sp.ValueFlags[i]) && k < len(values); j++ {
			if sp.ValueFlags[i][j].(bool) {
				sp.Values[i][j] = values[k]
				k++
			} else {
				sp.Values[i][j] = nil
			}
		}
		self.RowsNumber = len(sp.ValueFlags[i])
	}

	keyNum := self.MD.GetKeyNumber()
	for i := 0; i < keyNum; i++ {
		buf, err := Util.ReadMessage(self.Reader)
		if err != nil {
			return sp, err
		}
		self.KeyFlags[i], err = Type.DecodeBOOL(bytes.NewReader(buf))
		if err != nil {
			return sp, err
		}

		buf, err = Util.ReadMessage(self.Reader)
		t, err := self.MD.GetKeyTypeByIndex(i)
		if err != nil {
			return sp, err
		}
		keys, err := Type.DecodeValue(bytes.NewReader(buf), t)
		if err != nil {
			return sp, err
		}

		sp.Keys[i] = make([]interface{}, len(sp.KeyFlags[i]))
		k := 0
		for j := 0; j < len(sp.KeyFlags[i]) && k < len(keys); j++ {
			if sp.KeyFlags[i][j].(bool) {
				sp.Keys[i][j] = keys[k]
				k++
			} else {
				sp.Keys[i][j] = nil
			}
		}
	}
	return sp, nil
}

func (self *SplitBuffer) WriteRow(row *Row) error {
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
