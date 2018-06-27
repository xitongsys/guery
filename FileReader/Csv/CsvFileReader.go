package Csv

import (
	"encoding/csv"
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

	Indexes []int
}

func New(reader io.Reader, md *Metadata.Metadata) *CsvFileReader {
	return &CsvFileReader{
		Metadata: md,
		Reader:   csv.NewReader(reader),
	}
}

func (self *CsvFileReader) TypeConvert(rows []*Row.Row) ([]*Row.Row, error) {
	jobs := make(chan int)
	done := make(chan bool)
	colNum := self.Metadata.GetColumnNumber()
	colTypes := self.Metadata.GetColumnTypes()

	for i := 0; i < int(Config.Conf.Runtime.ParallelNumber); i++ {
		go func() {
			for {
				i, ok := <-jobs
				if ok {
					for j := 0; j < colNum; i++ {
						v := rows[i].Vals[j]
						cv := Type.ToType(v, colTypes[j])
						rows[i].Vals[j] = cv
					}
				} else {
					done <- true
					break
				}
			}
		}()
	}

	for i := 0; i < len(rows); i++ {
		jobs <- i
	}
	close(jobs)
	for i := 0; i < int(Config.Conf.Runtime.ParallelNumber); i++ {
		<-done
	}
	return rows, nil
}

func (self *CsvFileReader) Read(indexes []int) ([]*Row.Row, error) {
	var (
		err    error
		record []string
	)
	if self.Indexes == nil || len(self.Indexes) <= 0 {
		self.Indexes = make([]int, 0)
		if indexes == nil {
			for i := 0; i < self.Metadata.GetColumnNumber(); i++ {
				self.Indexes = append(self.Indexes, i)
			}
		} else {
			self.Indexes = indexes
		}
	}

	rows := []*Row.Row{}
	for i := 0; i < READ_ROWS_NUMBER; i++ {
		if record, err = self.Reader.Read(); err != nil {
			break
		}
		row := Row.NewRow()
		for _, index := range self.Indexes {
			row.AppendVals(record[index])
		}
		rows = append(rows, row)
	}

	if err != nil {
		if err == io.EOF && len(rows) > 0 {
			err = nil
		} else {
			return nil, err
		}
	}

	return self.TypeConvert(rows)
}
