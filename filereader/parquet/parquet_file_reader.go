package parquet

import (
	"fmt"
	"io"
	"sync"

	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/filesystem"
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/row"
	. "github.com/xitongsys/parquet-go/ParquetFile"
	. "github.com/xitongsys/parquet-go/ParquetReader"
	"github.com/xitongsys/parquet-go/parquet"
)

const (
	READ_ROWS_NUMBER = 10000
)

type Pair struct {
	I, Index int
}

type PqFile struct {
	FileName string
	VF       FileSystem.VirtualFile
}

func (self *PqFile) Create(name string) (ParquetFile, error) {
	return nil, nil
}

func (self *PqFile) Open(name string) (ParquetFile, error) {
	if name == "" {
		name = self.FileName
	}
	vf, err := FileSystem.Open(name)
	if err != nil {
		return nil, err
	}
	res := &PqFile{
		VF:       vf,
		FileName: name,
	}
	return res, nil
}
func (self *PqFile) Seek(offset int64, pos int) (int64, error) {
	return self.VF.Seek(offset, pos)
}
func (self *PqFile) Read(b []byte) (cnt int, err error) {
	var n int
	ln := len(b)
	for cnt < ln {
		n, err = self.VF.Read(b[cnt:])
		cnt += n
		if err != nil {
			break
		}
	}
	return cnt, err
}
func (self *PqFile) Write(b []byte) (n int, err error) {
	return 0, nil
}
func (self *PqFile) Close() error { return nil }

type ParquetFileReader struct {
	ParquetFile ParquetFile
	pqReader    *ParquetReader
	Metadata    *Metadata.Metadata
	NumRows     int
	Cursor      int

	ReadColumnIndexes        []int
	ReadColumnTypes          []*parquet.Type
	ReadColumnConvertedTypes []*parquet.ConvertedType
	OutMetadata              *Metadata.Metadata
}

func New(fileName string, md *metadata.Metadata) *ParquetFileReader {
	parquetFileReader := new(ParquetFileReader)
	var pqFile ParquetFile = &PqFile{}
	pqFile, _ = pqFile.Open(fileName)
	parquetFileReader.pqReader, _ = NewParquetColumnReader(pqFile, int64(Config.Conf.Runtime.ParallelNumber))
	parquetFileReader.NumRows = int(parquetFileReader.pqReader.GetNumRows())
	parquetFileReader.Metadata = md
	parquetFileReader.ParquetFile = pqFile
	return parquetFileReader
}

func (self *ParquetFileReader) Close() error {
	return self.ParquetFile.Close()
}

func (self *ParquetFileReader) ReadHeader() (fieldNames []string, err error) {
	return self.pqReader.SchemaHandler.ValueColumns, nil
}

func (self *ParquetFileReader) SetReadColumns(indexes []int) error {
	if self.ReadColumnIndexes == nil || len(self.ReadColumnIndexes) <= 0 {
		if indexes == nil {
			indexes = make([]int, len(self.pqReader.SchemaHandler.ValueColumns))
			for i := 0; i < len(indexes); i++ {
				indexes[i] = i
			}
		}
		cn := len(self.pqReader.SchemaHandler.ValueColumns)
		for _, index := range indexes {
			if index >= cn {
				return fmt.Errorf("ParquetReader: index out of range")
			}
			fieldName := self.pqReader.SchemaHandler.ValueColumns[index]
			schemaIndex := self.pqReader.SchemaHandler.MapIndex[fieldName]
			t := self.pqReader.SchemaHandler.SchemaElements[schemaIndex].Type
			ct := self.pqReader.SchemaHandler.SchemaElements[schemaIndex].ConvertedType

			self.ReadColumnIndexes = append(self.ReadColumnIndexes, index)
			self.ReadColumnTypes = append(self.ReadColumnTypes, t)
			self.ReadColumnConvertedTypes = append(self.ReadColumnConvertedTypes, ct)
		}

		self.OutMetadata = self.Metadata.SelectColumnsByIndexes(indexes)
	}
	return nil
}

//indexes should not change during read process
func (self *ParquetFileReader) Read(indexes []int) (*row.RowsGroup, error) {
	if self.Cursor >= self.NumRows {
		self.Close()
		return nil, io.EOF
	}

	if err := self.SetReadColumns(indexes); err != nil {
		return nil, err
	}

	var err error
	rg := Row.NewRowsGroup(self.OutMetadata)
	readRowsNumber := 0

	jobs := make(chan Pair)
	var wg sync.WaitGroup

	for i := 0; i < int(Config.Conf.Runtime.ParallelNumber); i++ {
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()

			for {
				pair, ok := <-jobs
				i, index := pair.I, pair.Index
				if ok {
					values, _, _ := self.pqReader.ReadColumnByIndex(index, READ_ROWS_NUMBER)
					if len(values) <= 0 {
						err = io.EOF
						return
					}
					gt, _ := self.Metadata.GetTypeByIndex(index)

					for j := 0; j < len(values); j++ {
						rg.Vals[i] = append(rg.Vals[i],
							ParquetTypeToGueryType(values[j],
								self.ReadColumnTypes[i],
								self.ReadColumnConvertedTypes[i],
								gt),
						)
					}

				} else {
					return
				}
			}
		}()
	}

	for i, index := range self.ReadColumnIndexes {
		jobs <- Pair{i, index}
	}
	close(jobs)
	wg.Wait()

	if len(self.ReadColumnIndexes) <= 0 { //for 0 columns read
		readRowsNumber = self.NumRows - self.Cursor
		if readRowsNumber > READ_ROWS_NUMBER {
			readRowsNumber = READ_ROWS_NUMBER
		}
	} else {
		readRowsNumber = len(rg.Vals[0])
	}

	self.Cursor += readRowsNumber
	rg.RowsNumber = readRowsNumber

	return rg, nil
}
