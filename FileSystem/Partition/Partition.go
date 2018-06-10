package Partition

import (
	"bytes"

	"github.com/xitongsys/guery/Util"
)

type Partition struct {
	Type   Util.Type
	Vals   []interface{}
	Buffer []byte
}

func NewPartition(t Util.Type) *Partition {
	return &Partition{
		Type:   t,
		Vals:   []interface{}{},
		Buffer: []byte{},
	}
}

func (self *Partition) Encode() {
	self.Buffer = Util.EncodeValues(self.Vals, self.Type)
}

func (self *Partition) Decode() (err error) {
	reader := bytes.NewReader(self.Buffer)
	self.Vals, err = Util.DecodeValue(reader, self.Type)
	return err
}

func (self *Partition) Append(val interface{}) {
	self.Vals = append(self.Vals, val)
}
