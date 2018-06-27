package Orc

import (
	"fmt"
	"io"

	"github.com/scritchley/orc"
	"github.com/scritchley/orc/proto"
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Row"
)

const (
	READ_ROWS_NUMBER = 10000
)

type OrcFileReader struct {
	Reader            *orc.Reader
	Cursor            *orc.Cursor
	Metadata          *Metadata.Metadata
	ReadColumnNames   []string
	ReadColumnTypes   []proto.Type_Kind
	ReadColumnIndexes []int
}

func New(reader orc.SizedReaderAt, md *Metadata.Metadata) (*OrcFileReader, error) {
	t, err := orc.NewReader(reader)
	if err != nil {
		return nil, err
	}
	return &OrcFileReader{
		Reader:   t,
		Metadata: md,
	}, nil
}

func (self *OrcFileReader) SetReadColumns(indexes []int) error {
	self.ReadColumnIndexes = append(self.ReadColumnIndexes, indexes...)
	columns := self.Reader.Schema().Columns()
	types := self.Reader.Schema().Types()
	self.ReadColumnNames = []string{}
	for _, index := range indexes {
		if index >= len(columns) {
			return fmt.Errorf("[Orc.SetReadColumns] index out of range")
		}
		self.ReadColumnNames = append(self.ReadColumnNames, columns[index])
		self.ReadColumnTypes = append(self.ReadColumnTypes, *(types[index].Kind))
	}
	return nil
}

func (self *OrcFileReader) Read(indexes []int) ([]*Row.Row, error) {
	var err error
	if self.Cursor == nil {
		if err = self.SetReadColumns(indexes); err != nil {
			return nil, err
		}
		self.Cursor = self.Reader.Select(self.ReadColumnNames...)
	}

	rows := []*Row.Row{}
	for i := 0; i < READ_ROWS_NUMBER; i++ {
		if self.Cursor.Next() || self.Cursor.Stripes() && self.Cursor.Next() {
			if err = self.Cursor.Err(); err != nil {
				return nil, err
			}
			row := Row.NewRow()
			for j, v := range self.Cursor.Row() {
				gv := OrcTypeToGueryType(v, self.ReadColumnTypes[j])
				row.AppendVals(gv)
			}
			rows = append(rows, row)

		} else {
			break
		}
	}
	if len(rows) <= 0 {
		return nil, io.EOF
	}
	return rows, nil

}
