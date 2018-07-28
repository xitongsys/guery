package orc

import (
	"fmt"
	"io"

	"github.com/scritchley/orc"
	"github.com/scritchley/orc/proto"
	"github.com/xitongsys/guery/filesystem"
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/row"
)

const (
	READ_ROWS_NUMBER = 10000
)

type OrcFileReader struct {
	Closer            io.Closer
	Reader            *orc.Reader
	Cursor            *orc.Cursor
	Metadata          *metadata.Metadata
	OutMetadata       *metadata.Metadata
	ReadColumnNames   []string
	ReadColumnTypes   []proto.Type_Kind
	ReadColumnIndexes []int
}

func New(reader orc.SizedReaderAt, md *metadata.Metadata) (*OrcFileReader, error) {
	t, err := orc.NewReader(reader)
	if err != nil {
		return nil, err
	}
	return &OrcFileReader{
		Reader:   t,
		Metadata: md,
		Closer:   io.Closer(reader.(filesystem.VirtualFile)),
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
		self.ReadColumnTypes = append(self.ReadColumnTypes, *(types[index+1].Kind))
	}
	self.OutMetadata = self.Metadata.SelectColumnsByIndexes(indexes)
	return nil
}

func (self *OrcFileReader) Read(indexes []int) (*row.RowsGroup, error) {
	var err error
	if self.Cursor == nil {
		if err = self.SetReadColumns(indexes); err != nil {
			return nil, err
		}
		self.Cursor = self.Reader.Select(self.ReadColumnNames...)
	}

	rg := row.NewRowsGroup(self.OutMetadata)
	for i := 0; i < READ_ROWS_NUMBER; i++ {
		if self.Cursor.Next() || self.Cursor.Stripes() && self.Cursor.Next() {
			if err = self.Cursor.Err(); err != nil {
				return nil, err
			}
			for j, v := range self.Cursor.Row() {
				gv := OrcTypeToGueryType(v, self.ReadColumnTypes[j])
				rg.Vals[j] = append(rg.Vals[j], gv)
			}
			rg.RowsNumber++

		} else {
			break
		}
	}
	if rg.RowsNumber <= 0 {
		self.Closer.Close()
		return nil, io.EOF
	}
	return rg, nil
}
