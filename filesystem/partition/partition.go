package partition

import (
	"bytes"

	"github.com/xitongsys/guery/gtype"
)

type Partition struct {
	Type   gtype.Type
	Vals   []interface{}
	Buffer []byte
}

func NewPartition(t gtype.Type) *Partition {
	return &Partition{
		Type:   t,
		Vals:   []interface{}{},
		Buffer: []byte{},
	}
}

func (self *Partition) Encode() {
	self.Buffer = gtype.EncodeValues(self.Vals, self.Type)
}

func (self *Partition) Decode() (err error) {
	reader := bytes.NewReader(self.Buffer)
	self.Vals, err = gtype.DecodeValue(reader, self.Type)
	return err
}

func (self *Partition) Append(val interface{}) {
	self.Vals = append(self.Vals, val)
}
