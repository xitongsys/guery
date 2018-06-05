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
	nilNum := 0
	for i := 0; i < ln; i++ {
		if nums[i] == nil {
			nilNum++
			continue
		}
		if nums[i].(bool) {
			res[i/8] = res[i/8] | (1 << uint32(i%8))
		}
	}
	bufWriter.Write(res)
	res2 := bufWriter.Bytes()

	numBufWriter := new(bytes.Buffer)
	binary.Write(numBufWriter, binary.LittleEndian, int32(len(nums)-nilNum))
	numBuf := numBufWriter.Bytes()
	for i := 0; i < len(numBuf); i++ {
		res2[i] = numBuf[i]
	}
	return res2
}

//INT32, INT64, FLOAT32, FLOAT64
func EncodeNumber(nums []interface{}) []byte {
	bufWriter := new(bytes.Buffer)
	binary.Write(bufWriter, binary.LittleEndian, int32(len(nums)))
	nilNum := 0
	for _, num := range nums {
		if num == nil {
			nilNum++
			continue
		}
		binary.Write(bufWriter, binary.LittleEndian, num)
	}
	res := bufWriter.Bytes()

	numBufWriter := new(bytes.Buffer)
	binary.Write(numBufWriter, binary.LittleEndian, int32(len(nums)-nilNum))
	numBuf := numBufWriter.Bytes()

	for i := 0; i < len(numBuf); i++ {
		res[i] = numBuf[i]
	}
	return res
}

//STRING
func EncodeString(ss []interface{}) []byte {
	bufWriter := new(bytes.Buffer)
	binary.Write(bufWriter, binary.LittleEndian, int32(len(ss)))
	nilNum := 0
	for _, si := range ss {
		if si == nil {
			nilNum++
			continue
		}
		s := si.(string)
		ln := int32(len(s))
		binary.Write(bufWriter, binary.LittleEndian, ln)
		bufWriter.Write([]byte(s))
	}
	res := bufWriter.Bytes()

	numBufWriter := new(bytes.Buffer)
	binary.Write(numBufWriter, binary.LittleEndian, int32(len(ss)-nilNum))
	numBuf := numBufWriter.Bytes()

	for i := 0; i < len(numBuf); i++ {
		res[i] = numBuf[i]
	}
	return res

}

//TIMESTAMP
func EncodeTime(ts []interface{}) []byte {
	nums := []interface{}{}
	for _, ti := range ts {
		if ti == nil {
			nums = append(nums, nil)
		} else {
			nums = append(nums, ti.(time.Time).Unix())
		}
	}
	return EncodeNumber(nums)
}
