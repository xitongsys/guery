package Util

import (
	"bytes"
	"encoding/binary"
)

func DecodeBOOLEAN(bytesReader *bytes.Reader, cnt uint64) ([]interface{}, error) {
	var (
		err error
	)

	res := make([]interface{}, cnt)
	bytesBuf := []byte{}
	totNum := (cnt + 7) / 8
	readNum := 0
	for readNum < bytesNum {
		curNum := totNum - readNum
		buf := make([]byte, curNum)
		n, err := bytesReader.Read(buf)
		if err != nil {
			return res, err
		}
		bytesBuf = append(bytesBuf, buf...)
		readNum += n
	}

	for i := 0; i < int(cnt); i++ {
	}
	return res, err
}

func DecodeINT32(bytesReader *bytes.Reader, cnt uint64) ([]interface{}, error) {
	var err error
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

func DecodeINT64(bytesReader *bytes.Reader, cnt uint64) ([]interface{}, error) {
	var err error
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
