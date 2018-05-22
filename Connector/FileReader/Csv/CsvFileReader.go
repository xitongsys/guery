package csv

import (
	"encoding/csv"
	"fmt"
	"io"

	"github.com/xitongsys/guery/Util"
)

type CsvFileReader struct {
	Metadata *Util.Metadata
	Reader   *csv.Reader
}

func New(reader io.Reader, md *Util.Metadata) *CsvFileReader {
	return &CsvFileReader{
		Metadata: md,
		Reader:   csv.NewReader(reader),
	}
}

func (self *CsvFileReader) Read() (row *util.Row, err error) {
	var record []string
	var objects []interface{}
	record, err = self.Reader.Read()
	if err != nil {
		return
	}
	if len(record) != len(self.Metadata.Columns) {
		return nil, fmt.Errorf("csv file doesn't match metadata")
	}
	row = &Util.Row{}
	for i := 0; i < len(record); i++ {
		valstr := record[i]
		valtype := self.Metadata.Columns[i].ColumnType

		val := Util.ToType(valstr, valtype)
		row.AppendVals(val)
	}
	return row, nil

}

func (self *CsvFileReader) ReadByColumns(indexes []int) (row *util.Row, err error) {
	var record []string
	var objects []interface{}
	record, err = self.Reader.Read()
	if err != nil {
		return
	}
	if len(record) != len(self.Metadata.Columns) {
		return nil, fmt.Errorf("csv file doesn't match metadata")
	}
	row = &Util.Row{}
	for _, index := range indexes {
		valstr := record[index]
		valtype := self.Metadata.Columns[index].ColumnType

		val := Util.ToType(valstr, valtype)
		row.AppendVals(val)
	}
	return row, nil

}
