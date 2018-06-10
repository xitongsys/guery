package Util

import (
	"time"
	//"github.com/vmihailenco/msgpack"
)

type Date time.Time

/*
var _ msgpack.CustomEncoder = (*Date)(nil)
var _ msgpack.CustomDecoder = (*Date)(nil)

func (d Date) EncodeMsgpack(enc *msgpack.Encoder) error {
	t := time.Time(d)
	return enc.Encode(t)
}

func (d *Date) DecodeMsgpack(dec *msgpack.Decoder) error {
	t := (*time.Time)(d)
	return dec.Decode(t)
}
*/

func (self Date) String() string {
	return time.Time(self).Format("2006-01-02")
}
