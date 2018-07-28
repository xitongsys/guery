package Type

import (
	"fmt"
	"math"
)

//LittleEndian

func ToKeyStringINT32(num int32) string {
	buf := make([]byte, 4)
	v := uint32(num)
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
	return string(buf)

}

func ToKeyStringINT64(num int64) string {
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
	return string(buf)
}

func ToKeyStringFLOAT32(num float32) string {
	buf := make([]byte, 4)
	v := math.Float32bits(num)
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
	return string(buf)
}

func ToKeyStringFLOAT64(num float64) string {
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
	return string(buf)
}

func ToKeyStringBOOL(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

func ToKeyStringTIMESTAMP(t Timestamp) string {
	return ToKeyStringINT64(t.Sec)
}

func ToKeyStringDATE(t Date) string {
	return ToKeyStringINT64(t.Sec)
}

func ToKeyString(n interface{}) string {
	switch n.(type) {
	case bool:
		return ToKeyStringBOOL(n.(bool))
	case int32:
		return ToKeyStringINT32(n.(int32))
	case int64:
		return ToKeyStringINT64(n.(int64))
	case float32:
		return ToKeyStringFLOAT32(n.(float32))
	case float64:
		return ToKeyStringFLOAT64(n.(float64))
	case Date:
		return ToKeyStringDATE(n.(Date))
	case Timestamp:
		return ToKeyStringTIMESTAMP(n.(Timestamp))
	}
	return fmt.Sprintf("%v", n)
}
