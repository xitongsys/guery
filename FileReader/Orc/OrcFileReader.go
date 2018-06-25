package Orc

import (
	"fmt"
	"io"

	"github.com/scritchley/orc"
	"github.com/scritchley/orc/proto"
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Split"
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

func (self *OrcFileReader) Read(indexes []int) (*Split.Split, error) {
	if self.Cursor == nil {
		if err = self.SetReadColumns(indexes); err != nil {
			return nil, err
		}
		self.Cursor = self.Reader.Select(self.ReadColumnNames...)
	}

	sp := Split.NewSplit(self.Metadata)
	for i := 0; i < Split.MAX_SPLIT_SIZE; i++ {
		if self.Cursor.Next() || self.Cursor.Stripes() && self.Cursor.Next() {
			if err = self.Cursor.Err(); err != nil {
				return nil, err
			}
			for i, v := range self.Cursor.Row() {
				gv := OrcTypeToGueryType(v, self.ReadColumnTypes[i])
				sp.Values[i] = append(sp.Values[i], gv)
				if gv == nil {
					sp.ValueFlags = append(sp.ValueFlags, false)
				} else {
					sp.ValueFlags = append(sp.ValueFlags, true)
				}
			}
			sp.NumRows++

		} else {
			break
		}
	}
	if sp.GetRowsNumber() <= 0 {
		return sp, io.EOF
	}
	return sp, nil

}
