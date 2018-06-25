package Orc

import (
	"fmt"
	"io"

	"github.com/scritchley/orc"
	"github.com/scritchley/orc/proto"
	"github.com/xitongsys/guery/Split"
)

type OrcFileReader struct {
	Reader            *orc.Reader
	Cursor            *orc.Cursor
	ReadColumnNames   []string
	ReadColumnTypes   []proto.Type_Kind
	ReadColumnIndexes []int
}

func New(reader orc.SizedReaderAt) (*OrcFileReader, error) {
	t, err := orc.NewReader(reader)
	if err != nil {
		return nil, err
	}
	return &OrcFileReader{
		Reader: t,
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

func (self *OrcFileReader) Read(indexes []int) (row *Split.Split, err error) {
	if self.Cursor == nil {
		if err = self.SetReadColumns(indexes); err != nil {
			return nil, err
		}
		self.Cursor = self.Reader.Select(self.ReadColumnNames...)
	}

	if self.Cursor.Next() || self.Cursor.Stripes() && self.Cursor.Next() {
		if err = self.Cursor.Err(); err != nil {
			return nil, err
		}
		row := Row.NewRow()
		for i, v := range self.Cursor.Row() {
			row.AppendVals(OrcTypeToGueryType(v, self.ReadColumnTypes[i]))
		}
		return row, nil
	}
	return nil, io.EOF
}
