package Util

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

const (
	MessageEOF = math.MinInt32
)

func ReadMessage(reader io.Reader) (res []byte, err error) {
	var length int32
	err = binary.Read(reader, binary.LittleEndian, &length)
	if err == io.EOF || length == MessageEOF {
		return nil, io.EOF
	}
	if err != nil {
		return nil, err
	}
	if length == 0 {
		return nil, nil
	}

	res = make([]byte, length)
	var n int
	n, err = io.ReadFull(reader, res)
	if err == io.EOF {
		return nil, fmt.Errorf("Unexpected EOF when reading message, expected read %v, only read %v", length, n)
	}
	if err != nil {
		return nil, err
	}
	return res, nil
}

func WriteMessage(writer io.Writer, msg []byte) (err error) {
	if err = binary.Write(writer, binary.LittleEndian, int32(len(msg))); err != nil {
		return err
	}
	if _, err = writer.Write(msg); err != nil {
		return err
	}
	return nil
}

func WriteEOFMessage(writer io.Writer) (err error) {
	if err = binary.Write(writer, binary.LittleEndian, MessageEOF); err != nil {
		return err
	}
	return nil
}
