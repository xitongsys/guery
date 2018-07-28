package gtype

import (
	"fmt"
	"math"
)

//LittleEndian

func ToBinaryINT32(num int32) []byte {
	buf := make([]byte, 4)
	v := uint32(num)
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
	return buf
}

func ToBinaryINT64(num int64) []byte {
	buf := make([]byte, 8)
	v := uint64(num)
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
	buf[4] = byte(v >> 32)
	buf[5] = byte(v >> 40)
	buf[6] = byte(v >> 48)
	buf[7] = byte(v >> 56)
	return buf
}

func ToBinaryFLOAT32(num float32) []byte {
	buf := make([]byte, 4)
	v := math.Float32bits(num)
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
	return buf
}

func ToBinaryFLOAT64(num float64) []byte {
	buf := make([]byte, 8)
	v := math.Float64bits(num)
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
	buf[4] = byte(v >> 32)
	buf[5] = byte(v >> 40)
	buf[6] = byte(v >> 48)
	buf[7] = byte(v >> 56)
	return buf
}

func ToBinarySTRING(s string) []byte {
	return []byte(s)
}

func ToBinaryTIMESTAMP(t Timestamp) []byte {
	return ToBinaryINT64(t.Sec)
}

func ToBinaryDATE(t Date) []byte {
	return ToBinaryINT64(t.Sec)
}

func ToBinary(n interface{}) []byte {
	switch n.(type) {
	case int32:
		return ToBinaryINT32(n.(int32))
	case int64:
		return ToBinaryINT64(n.(int64))
	case float32:
		return ToBinaryFLOAT32(n.(float32))
	case float64:
		return ToBinaryFLOAT64(n.(float64))
	case Date:
		return ToBinaryDATE(n.(Date))
	case Timestamp:
		return ToBinaryTIMESTAMP(n.(Timestamp))
	}
	return []byte(fmt.Sprintf("%v", n))
}
