package csv

import (
	"encoding/csv"
	"fmt"
	"io"

	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/filesystem"
	"github.com/xitongsys/guery/gtype"
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/row"
)

const (
	READ_ROWS_NUMBER = 10000
)

type CsvFileReader struct {
	Closer   io.Closer
	Metadata *Metadata.Metadata
	Reader   *csv.Reader

	Indexes     []int
	OutMetadata *Metadata.Metadata
}

func New(reader io.Reader, md *metadata.Metadata) *CsvFileReader {
	return &CsvFileReader{
		Metadata: md,
		Reader:   csv.NewReader(reader),
		Closer:   io.Closer(reader.(FileSystem.VirtualFile)),
	}
}

func (self *CsvFileReader) TypeConvert(rg *row.RowsGroup) (*row.RowsGroup, error) {
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
	for i := 0; i < int(config.Conf.Runtime.ParallelNumber); i++ {
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

func (self *CsvFileReader) Read(indexes []int) (*row.RowsGroup, error) {
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
		self.Closer.Close()
		if err == io.EOF && rg.RowsNumber > 0 {
			err = nil
		} else {
			return nil, err
		}
	}

	return self.TypeConvert(rg)
}
