package partition

import (
	"bytes"

	"github.com/xitongsys/guery/gtype"
)

type Partition struct {
	Type   Type.Type
	Vals   []interface{}
	Buffer []byte
}

func NewPartition(t Type.Type) *Partition {
	return &Partition{
		Type:   t,
		Vals:   []interface{}{},
		Buffer: []byte{},
	}
}

func (self *Partition) Encode() {
	self.Buffer = Type.EncodeValues(self.Vals, self.Type)
}

func (self *Partition) Decode() (err error) {
	reader := bytes.NewReader(self.Buffer)
	self.Vals, err = Type.DecodeValue(reader, self.Type)
	return err
}

func (self *Partition) Append(val interface{}) {
	self.Vals = append(self.Vals, val)
}
