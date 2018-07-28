package Util

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"

	"github.com/golang/snappy"
)

type CompressType int32

const (
	UNCOMPRESSED CompressType = iota
	GZIP
	SNAPPY
)

//Uncompress using Gzip
func UncompressGzip(buf []byte) ([]byte, error) {
	rbuf := bytes.NewReader(buf)
	gzipReader, _ := gzip.NewReader(rbuf)
	res, err := ioutil.ReadAll(gzipReader)
	return res, err
}

//Compress using Gzip
func CompressGzip(buf []byte) []byte {
	var res bytes.Buffer
	gzipWriter := gzip.NewWriter(&res)
	gzipWriter.Write(buf)
	gzipWriter.Close()
	return res.Bytes()
}

//Uncompress using Snappy
func UncompressSnappy(buf []byte) ([]byte, error) {
	return snappy.Decode(nil, buf)
}

//Compress using Snappy
func CompressSnappy(buf []byte) []byte {
	return snappy.Encode(nil, buf)
}

func Uncompress(buf []byte, compressType CompressType) ([]byte, error) {
	switch compressType {
	case GZIP:
		return UncompressGzip(buf)
	case SNAPPY:
		return UncompressSnappy(buf)
	case UNCOMPRESSED:
		return buf, nil
	default:
		return nil, fmt.Errorf("Unsupported compress method")
	}
}

func Compress(buf []byte, compressType CompressType) []byte {
	switch compressType {
	case GZIP:
		return CompressGzip(buf)
	case SNAPPY:
		return CompressSnappy(buf)
	case UNCOMPRESSED:
		return buf
	default:
		return nil
	}
}
