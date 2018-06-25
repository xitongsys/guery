package Csv

import (
	"encoding/csv"
	"io"

	"github.com/xitongsys/guery/Config"
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Split"
	"github.com/xitongsys/guery/Type"
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

func (self *CsvFileReader) TypeConvert(sp *Split.Split) (*Split.Split, error) {
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
						v := sp.Values[j][i]
						cv := Type.ToType(v, colTypes[j])
						sp.Values[j][i] = cv
						if cv == nil {
							sp.ValueFlags[j][i] = false
						}
					}
				} else {
					done <- true
					break
				}
			}
		}()
	}

	for i := 0; i < sp.GetRowsNumber(); i++ {
		jobs <- i
	}
	close(jobs)
	for i := 0; i < int(Config.Conf.Runtime.ParallelNumber); i++ {
		<-done
	}
	return sp, nil
}

func (self *CsvFileReader) Read(indexes []int) (*Split.Split, error) {
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

	sp := Split.NewSplit(self.Metadata)
	for i := 0; i < Split.MAX_SPLIT_SIZE; i++ {
		if record, err = self.Reader.Read(); err != nil {
			break
		}

		for i, index := range self.Indexes {
			sp.Values[i] = append(sp.Values[i], record[index])
			sp.ValueFlags[i] = append(sp.ValueFlags[i], true)
		}
		sp.RowsNumber++
	}

	if err != nil {
		if err == io.EOF && sp.GetRowsNumber() > 0 {
			err = nil
		} else {
			return sp, err
		}
	}

	return self.TypeConvert(sp)
}
