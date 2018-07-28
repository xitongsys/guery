package gtype

import (
	"bytes"
	"encoding/binary"
)

func EncodeValues(vals []interface{}, t Type) []byte {
	switch t {
	case BOOL:
		return EncodeBool(vals)
	case INT32:
		return EncodeINT32(vals)
	case INT64:
		return EncodeINT64(vals)
	case FLOAT32:
		return EncodeFLOAT32(vals)
	case FLOAT64:
		return EncodeFLOAT64(vals)
	case STRING:
		return EncodeString(vals)
	case TIMESTAMP:
		return EncodeTime(vals)
	case DATE:
		return EncodeDate(vals)
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

//INT32
func EncodeINT32(nums []interface{}) []byte {
	bufWriter := new(bytes.Buffer)
	numBufWriter := new(bytes.Buffer)
	n, _ := BinaryWriteINT32(numBufWriter, nums)
	binary.Write(bufWriter, binary.LittleEndian, int32(n))
	bufWriter.Write(numBufWriter.Bytes())
	return bufWriter.Bytes()
}

//INT64
func EncodeINT64(nums []interface{}) []byte {
	bufWriter := new(bytes.Buffer)
	numBufWriter := new(bytes.Buffer)
	n, _ := BinaryWriteINT64(numBufWriter, nums)
	binary.Write(bufWriter, binary.LittleEndian, int32(n))
	bufWriter.Write(numBufWriter.Bytes())
	return bufWriter.Bytes()
}

//FLOAT32
func EncodeFLOAT32(nums []interface{}) []byte {
	bufWriter := new(bytes.Buffer)
	numBufWriter := new(bytes.Buffer)
	n, _ := BinaryWriteFLOAT32(numBufWriter, nums)
	binary.Write(bufWriter, binary.LittleEndian, int32(n))
	bufWriter.Write(numBufWriter.Bytes())
	return bufWriter.Bytes()
}

//FLOAT64
func EncodeFLOAT64(nums []interface{}) []byte {
	bufWriter := new(bytes.Buffer)
	numBufWriter := new(bytes.Buffer)
	n, _ := BinaryWriteFLOAT64(numBufWriter, nums)
	binary.Write(bufWriter, binary.LittleEndian, int32(n))
	bufWriter.Write(numBufWriter.Bytes())
	return bufWriter.Bytes()
}

//STRING
func EncodeString(ss []interface{}) []byte {
	bufWriter := new(bytes.Buffer)
	bufWriter.Write(ToBinaryINT32(int32(len(ss))))
	nilNum := 0
	for _, si := range ss {
		if si == nil {
			nilNum++
			continue
		}
		s := si.(string)
		ln := int32(len(s))
		bufWriter.Write(ToBinaryINT32(ln))
		bufWriter.Write([]byte(s))
	}
	res := bufWriter.Bytes()

	numBufWriter := new(bytes.Buffer)
	numBufWriter.Write(ToBinaryINT32(int32(len(ss) - nilNum)))
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
			nums = append(nums, ti.(Timestamp).Sec)
		}
	}
	return EncodeINT64(nums)
}

//DATE
func EncodeDate(ts []interface{}) []byte {
	nums := []interface{}{}
	for _, ti := range ts {
		if ti == nil {
			nums = append(nums, nil)
		} else {
			nums = append(nums, ti.(Date).Sec)
		}
	}
	return EncodeINT64(nums)
}
