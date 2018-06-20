package Csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"

	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Row"
	"github.com/xitongsys/guery/Type"
)

type CsvFileReader struct {
	Metadata *Metadata.Metadata
	Reader   *csv.Reader
}

func New(reader io.Reader, md *Metadata.Metadata) *CsvFileReader {
	return &CsvFileReader{
		Metadata: md,
		Reader:   csv.NewReader(reader),
	}
}

func (self *CsvFileReader) Read(indexes []int) (row *Row.Row, err error) {
	var record []string
	record, err = self.Reader.Read()
	if err != nil {
		return
	}

	if len(record) != len(self.Metadata.Columns) {
		return nil, fmt.Errorf("csv file doesn't match metadata")
	}

	log.Println("=====", indexes, self.Metadata)

	row = &Row.Row{}
	if indexes != nil {
		for _, index := range indexes {
			valstr := record[index]
			valtype := self.Metadata.Columns[index].ColumnType
			val := Type.ToType(valstr, valtype)
			row.AppendVals(val)
		}
	} else {
		for i := 0; i < len(record); i++ {
			valstr := record[i]
			valtype := self.Metadata.Columns[i].ColumnType
			val := Type.ToType(valstr, valtype)
			row.AppendVals(val)
		}
	}

	return row, nil

}
