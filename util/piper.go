package util

import (
	"io"
)

const (
	BUFFER_SIZE = 1024 * 512
)

type Piper struct {
	Reader *io.PipeReader
	Writer *io.PipeWriter
}

func NewPiper() *Piper {
	pr, pw := io.Pipe()
	return &Piper{
		Reader: pr,
		Writer: pw,
	}
}

func CopyBuffer(src io.Reader, dst io.Writer) (err error) {
	buf := make([]byte, BUFFER_SIZE)

	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er != nil {
			err = er
			break
		}
	}
	return err
}
