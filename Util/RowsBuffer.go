package Util

import (
	"io"
)

type RowsBuffer struct {
	MD         *Metadata
	RowsNumber int

	ValueBuffers  [][]interface{}
	ValueNilFlags [][]bool

	KeyBuffers  [][]interface{}
	KeyNilFlags [][]bool

	Reader io.Reader
	Writer io.Writer
}

func (self *RowsBuffer) WriteRow(row *Row) error {

}
