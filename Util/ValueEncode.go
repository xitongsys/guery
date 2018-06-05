package Util

import (
	"bytes"
	"encoding/binary"
	"time"
)

//BOOLEAN
func EncodeBool(nums []interface{}) []byte {
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
	return res
}

//INT32, INT64, FLOAT32, FLOAT64
func EncodeNumber(nums []interface{}) []byte {
	bufWriter := new(bytes.Buffer)
	for _, num := range nums {
		if num == nil {
			continue
		}
		binary.Write(bufWriter, binary.LittleEndian, nums[i])
	}
	return bufWriter.Bytes()
}

//STRING
func EncodeString(ss []interface{}) []byte {
	bufWriter := new(bytes.Buffer)
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
