package DataSource

type ColumnBuffer interface {
	Read() interface{}
	Size() int
	Duplicate() ColumnBuffer
}

/////////////////////////////////////////////
type MemColumnBuffer struct {
	Vals  []interface{}
	Index int
}

func NewMemColumnBuffer() *MemColumnBuffer {
	return &MemColumnBuffer{}
}

func (self *MemColumnBuffer) Read() interface{} {
	if self.Index >= len(self.Vals) {
		return nil
	}
	res := self.Vals[self.Index]
	self.Index++
	return res
}

func (self *MemColumnBuffer) Size() int {
	return len(self.Vals)
}

func (self *MemColumnBuffer) Duplicate() ColumnBuffer {
	res := &MemColumnBuffer{
		Vals:  self.Vals,
		Index: 0,
	}
	return res
}

func (self *MemColumnBuffer) Append(vals ...interface{}) {
	self.Vals = append(self.Vals, vals)
}
