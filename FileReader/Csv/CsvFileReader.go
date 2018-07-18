package Csv

import (
	"encoding/csv"
	"fmt"
	"io"

	"github.com/xitongsys/guery/Config"
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Row"
	"github.com/xitongsys/guery/Type"
)

const (
	READ_ROWS_NUMBER = 10000
)

type CsvFileReader struct {
	Metadata *Metadata.Metadata
	Reader   *csv.Reader

	Indexes     []int
	OutMetadata *Metadata.Metadata
}

func New(reader io.Reader, md *Metadata.Metadata) *CsvFileReader {
	return &CsvFileReader{
		Metadata: md,
		Reader:   csv.NewReader(reader),
	}
}

func (self *CsvFileReader) TypeConvert(rg *Row.RowsGroup) (*Row.RowsGroup, error) {
	jobs := make(chan int)
	done := make(chan bool)
	cn := len(self.Indexes)
	colTypes := make([]Type.Type, cn)
	for i := 0; i < cn; i++ {
		colTypes[i], _ = self.Metadata.GetTypeByIndex(self.Indexes[i])
	}

	for i := 0; i < int(Config.Conf.Runtime.ParallelNumber); i++ {
		go func() {
			for {
				j, ok := <-jobs
				if ok {
					for k := 0; k < cn; k++ {
						v := rg.Vals[k][j]
						cv := Type.ToType(v, colTypes[k])
						rg.Vals[k][j] = cv
					}
				} else {
					done <- true
					break
				}
			}
		}()
	}

	for i := 0; i < rg.RowsNumber; i++ {
		jobs <- i
	}
	close(jobs)
	for i := 0; i < int(Config.Conf.Runtime.ParallelNumber); i++ {
		<-done
	}
	return rg, nil
}

func (self *CsvFileReader) SetReadColumns(indexes []int) error {
	cn := self.Metadata.GetColumnNumber()
	if self.Indexes == nil || len(self.Indexes) <= 0 {
		self.Indexes = make([]int, 0)
		if indexes == nil {
			for i := 0; i < cn; i++ {
				self.Indexes = append(self.Indexes, i)
			}
		} else {
			self.Indexes = indexes
		}

		for _, i := range self.Indexes {
			if i >= cn {
				return fmt.Errorf("CsvFileReader: index out of range")
			}
		}
		self.OutMetadata = self.Metadata.SelectColumnsByIndexes(self.Indexes)
	}
	return nil

}

func (self *CsvFileReader) Read(indexes []int) (*Row.RowsGroup, error) {
	var (
		err    error
		record []string
	)

	if err = self.SetReadColumns(indexes); err != nil {
		return nil, err
	}

	rg := Row.NewRowsGroup(self.OutMetadata)
	for i := 0; i < READ_ROWS_NUMBER; i++ {
		if record, err = self.Reader.Read(); err != nil {
			break
		}
		for i, index := range self.Indexes {
			rg.Vals[i] = append(rg.Vals[i], record[index])
		}
		rg.RowsNumber++
	}

	if err != nil {
		if err == io.EOF && rg.RowsNumber > 0 {
			err = nil
		} else {
			return nil, err
		}
	}

	return self.TypeConvert(rg)
}
