package Util

import (
	"bytes"
	"encoding/binary"
	"time"
)

func EncodeValues(vals []interface{}, t Type) []byte {
	switch t {
	case BOOL:
		return EncodeBool(vals)
	case INT32, INT64, FLOAT32, FLOAT64:
		return EncodeNumber(vals)
	case STRING:
		return EncodeString(vals)
	case TIMESTAMP:
		return EncodeTime(vals)
	}
	return []byte{}
}

//BOOLEAN
func EncodeBool(nums []interface{}) []byte {
	bufWriter := new(bytes.Buffer)
	binary.Write(bufWriter, binary.LittleEndian, int32(len(nums)))
	ln := len(nums)
	byteNum := (ln + 7) / 8
	res := make([]byte, byteNum)
	for i := 0; i < ln; i++ {
		if nums[i] == nil {
			continue
		}
		if nums[i].(bool) {
			res[i/8] = res[i/8] | (1 << uint32(i%8))
		}
	}
	bufWriter.Write(res)
	return bufWriter.Bytes()
}

//INT32, INT64, FLOAT32, FLOAT64
func EncodeNumber(nums []interface{}) []byte {
	bufWriter := new(bytes.Buffer)
	binary.Write(bufWriter, binary.LittleEndian, int32(len(nums)))
	for _, num := range nums {
		if num == nil {
			continue
		}
		binary.Write(bufWriter, binary.LittleEndian, num)
	}
	return bufWriter.Bytes()
}

//STRING
func EncodeString(ss []interface{}) []byte {
	bufWriter := new(bytes.Buffer)
	binary.Write(bufWriter, binary.LittleEndian, int32(len(ss)))
	for _, si := range ss {
		if si == nil {
			continue
		}
		s := si.(string)
		ln := int32(len(s))
		binary.Write(bufWriter, binary.LittleEndian, ln)
		bufWriter.Write([]byte(s))
	}
	return bufWriter.Bytes()
}

//TIMESTAMP
func EncodeTime(ts []interface{}) []byte {
	nums := []interface{}{}
	for _, ti := range ts {
		if ti == nil {
			continue
		}
		nums = append(nums, ti.(time.Time).Unix())
	}
	return EncodeNumber(nums)
}
