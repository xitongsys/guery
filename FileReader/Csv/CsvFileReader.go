package Csv

import (
	"encoding/csv"
	"io"

	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Row"
	"github.com/xitongsys/guery/Type"
)

var BUFFER_SIZE = 10000

type CsvFileReader struct {
	Metadata *Metadata.Metadata
	Reader   *csv.Reader

	//buffer
	RawRows [][]string
	Rows    []*Row.Row
	Index   int
	Size    int
}

func New(reader io.Reader, md *Metadata.Metadata) *CsvFileReader {
	return &CsvFileReader{
		Metadata: md,
		Reader:   csv.NewReader(reader),
		RawRows:  make([][]string, BUFFER_SIZE),
		Rows:     make([]*Row.Row, BUFFER_SIZE),
		Index:    0,
		Size:     0,
	}
}

func (self *CsvFileReader) RawRowToRow(indexes []int, i int) {
	row := &Row.Row{}
	record := self.RawRows[i]
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
	self.Rows[i] = row
}

func (self *CsvFileReader) readRows(indexes []int) error {
	self.Size = 0
	self.Index = 0
	for i := 0; i < BUFFER_SIZE; i++ {
		record, err := self.Reader.Read()
		if err != nil {
			return err
		}

		self.RawRows[self.Size] = record
		self.Size++
	}

	jobs := make(chan int)
	done := make(chan bool)

	for i := 0; i < 10; i++ {
		go func() {
			for {
				j, ok := <-jobs
				if ok {
					self.RawRowToRow(indexes, j)
				} else {
					done <- true
					return
				}
			}
		}()
	}

	for i := 0; i < self.Size; i++ {
		jobs <- i
	}
	close(jobs)
	<-done

	return nil
}

func (self *CsvFileReader) Read(indexes []int) (*Row.Row, error) {
	if self.Index >= self.Size {
		err := self.readRows(indexes)
		if err != nil {
			return nil, err
		}
	}

	self.Index++
	return self.Rows[self.Index-1], nil

}
