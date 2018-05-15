package Util

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Row) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Keys":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Keys) >= int(zb0002) {
				z.Keys = (z.Keys)[:zb0002]
			} else {
				z.Keys = make([]interface{}, zb0002)
			}
			for za0001 := range z.Keys {
				z.Keys[za0001], err = dc.ReadIntf()
				if err != nil {
					return
				}
			}
		case "Vals":
			var zb0003 uint32
			zb0003, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vals) >= int(zb0003) {
				z.Vals = (z.Vals)[:zb0003]
			} else {
				z.Vals = make([]interface{}, zb0003)
			}
			for za0002 := range z.Vals {
				z.Vals[za0002], err = dc.ReadIntf()
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Row) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Keys"
	err = en.Append(0x82, 0xa4, 0x4b, 0x65, 0x79, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Keys)))
	if err != nil {
		return
	}
	for za0001 := range z.Keys {
		err = en.WriteIntf(z.Keys[za0001])
		if err != nil {
			return
		}
	}
	// write "Vals"
	err = en.Append(0xa4, 0x56, 0x61, 0x6c, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Vals)))
	if err != nil {
		return
	}
	for za0002 := range z.Vals {
		err = en.WriteIntf(z.Vals[za0002])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Row) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Keys"
	o = append(o, 0x82, 0xa4, 0x4b, 0x65, 0x79, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Keys)))
	for za0001 := range z.Keys {
		o, err = msgp.AppendIntf(o, z.Keys[za0001])
		if err != nil {
			return
		}
	}
	// string "Vals"
	o = append(o, 0xa4, 0x56, 0x61, 0x6c, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Vals)))
	for za0002 := range z.Vals {
		o, err = msgp.AppendIntf(o, z.Vals[za0002])
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Row) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Keys":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Keys) >= int(zb0002) {
				z.Keys = (z.Keys)[:zb0002]
			} else {
				z.Keys = make([]interface{}, zb0002)
			}
			for za0001 := range z.Keys {
				z.Keys[za0001], bts, err = msgp.ReadIntfBytes(bts)
				if err != nil {
					return
				}
			}
		case "Vals":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vals) >= int(zb0003) {
				z.Vals = (z.Vals)[:zb0003]
			} else {
				z.Vals = make([]interface{}, zb0003)
			}
			for za0002 := range z.Vals {
				z.Vals[za0002], bts, err = msgp.ReadIntfBytes(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Row) Msgsize() (s int) {
	s = 1 + 5 + msgp.ArrayHeaderSize
	for za0001 := range z.Keys {
		s += msgp.GuessSize(z.Keys[za0001])
	}
	s += 5 + msgp.ArrayHeaderSize
	for za0002 := range z.Vals {
		s += msgp.GuessSize(z.Vals[za0002])
	}
	return
}
