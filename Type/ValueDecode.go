package Type

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"
)

func DecodeValue(bytesReader *bytes.Reader, t Type) ([]interface{}, error) {
	switch t {
	case BOOL:
		return DecodeBOOL(bytesReader)
	case INT32:
		return DecodeINT32(bytesReader)
	case INT64:
		return DecodeINT64(bytesReader)
	case FLOAT32:
		return DecodeFLOAT32(bytesReader)
	case FLOAT64:
		return DecodeFLOAT64(bytesReader)
	case STRING:
		return DecodeSTRING(bytesReader)
	case TIMESTAMP:
		return DecodeTIMESTAMP(bytesReader)
	case DATE:
		return DecodeDATE(bytesReader)
	}

	return []interface{}{}, fmt.Errorf("unknown type")
}

func DecodeBOOL(bytesReader *bytes.Reader) ([]interface{}, error) {
	var cnt int32
	if err := binary.Read(bytesReader, binary.LittleEndian, &cnt); err != nil {
		return []interface{}{}, err
	}
	res := make([]interface{}, cnt)
	totNum := (cnt + 7) / 8
	k := 0
	for i := 0; i < int(totNum) && k < int(cnt); i++ {
		b, err := bytesReader.ReadByte()
		if err != nil {
			return res, err
		}
		for j := 0; j < 8; j++ {
			if (uint32((1 << uint32(j))) & uint32(b)) > 0 {
				res[k] = true
			} else {
				res[k] = false
			}
			k++
			if k >= int(cnt) {
				break
			}
		}
	}
	return res, nil
}

func DecodeINT32(bytesReader *bytes.Reader) ([]interface{}, error) {
	var err error
	var cnt int32
	if err := binary.Read(bytesReader, binary.LittleEndian, &cnt); err != nil {
		return []interface{}{}, err
	}
	res := make([]interface{}, cnt)
	for i := 0; i < int(cnt); i++ {
		var cur int32
		if err = binary.Read(bytesReader, binary.LittleEndian, &cur); err != nil {
			break
		}
		res[i] = cur
	}
	return res, err
}

func DecodeINT64(bytesReader *bytes.Reader) ([]interface{}, error) {
	var err error
	var cnt int32
	if err := binary.Read(bytesReader, binary.LittleEndian, &cnt); err != nil {
		return []interface{}{}, err
	}
	res := make([]interface{}, cnt)
	for i := 0; i < int(cnt); i++ {
		var cur int64
		if err = binary.Read(bytesReader, binary.LittleEndian, &cur); err != nil {
			break
		}
		res[i] = cur
	}
	return res, err
}

func DecodeFLOAT32(bytesReader *bytes.Reader) ([]interface{}, error) {
	var err error
	var cnt int32
	if err := binary.Read(bytesReader, binary.LittleEndian, &cnt); err != nil {
		return []interface{}{}, err
	}
	res := make([]interface{}, cnt)
	for i := 0; i < int(cnt); i++ {
		var cur float32
		if err = binary.Read(bytesReader, binary.LittleEndian, &cur); err != nil {
			break
		}
		res[i] = cur
	}
	return res, err
}

func DecodeFLOAT64(bytesReader *bytes.Reader) ([]interface{}, error) {
	var err error
	var cnt int32
	if err := binary.Read(bytesReader, binary.LittleEndian, &cnt); err != nil {
		return []interface{}{}, err
	}
	res := make([]interface{}, cnt)
	for i := 0; i < int(cnt); i++ {
		var cur float64
		if err = binary.Read(bytesReader, binary.LittleEndian, &cur); err != nil {
			break
		}
		res[i] = cur
	}
	return res, err
}

func DecodeSTRING(bytesReader *bytes.Reader) ([]interface{}, error) {
	var err error
	var cnt int32
	if err := binary.Read(bytesReader, binary.LittleEndian, &cnt); err != nil {
		return []interface{}{}, err
	}
	res := make([]interface{}, cnt)
	for i := 0; i < int(cnt); i++ {
		buf := make([]byte, 4)
		if _, err = bytesReader.Read(buf); err != nil {
			break
		}
		ln := binary.LittleEndian.Uint32(buf)
		cur := make([]byte, ln)
		if _, err := bytesReader.Read(cur); err != nil {
			return res, err
		}
		res[i] = string(cur)
	}
	return res, err
}

func DecodeTIMESTAMP(bytesReader *bytes.Reader) ([]interface{}, error) {
	nums, err := DecodeINT64(bytesReader)
	if err != nil {
		return nums, err
	}
	res := make([]interface{}, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = time.Unix(nums[i].(int64), 0)
	}
	return res, nil
}

func DecodeDATE(bytesReader *bytes.Reader) ([]interface{}, error) {
	nums, err := DecodeINT64(bytesReader)
	if err != nil {
		return nums, err
	}
	res := make([]interface{}, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = Date(time.Unix(nums[i].(int64), 0))
	}
	return res, nil
}
