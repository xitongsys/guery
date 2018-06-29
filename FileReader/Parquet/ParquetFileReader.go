package Parquet

import (
	"io"

	"github.com/xitongsys/guery/Config"
	"github.com/xitongsys/guery/FileSystem"
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Row"
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
	pqReader *ParquetReader
	Metadata *Metadata.Metadata
	NumRows  int
	Cursor   int

	ReadColumnIndexes        []int
	ReadColumnTypes          []*parquet.Type
	ReadColumnConvertedTypes []*parquet.ConvertedType
}

func New(fileName string, md *Metadata.Metadata) *ParquetFileReader {
	parquetFileReader := new(ParquetFileReader)
	var pqFile ParquetFile = &PqFile{}
	pqFile, _ = pqFile.Open(fileName)
	parquetFileReader.pqReader, _ = NewParquetColumnReader(pqFile, int64(Config.Conf.Runtime.ParallelNumber))
	parquetFileReader.NumRows = int(parquetFileReader.pqReader.GetNumRows())
	parquetFileReader.Metadata = md
	return parquetFileReader
}

func (self *ParquetFileReader) ReadHeader() (fieldNames []string, err error) {
	return self.pqReader.SchemaHandler.ValueColumns, nil
}

func (self *ParquetFileReader) SetReadColumns(indexes []int) {
	for _, index := range indexes {
		fieldName := self.pqReader.SchemaHandler.ValueColumns[index]
		schemaIndex := self.pqReader.SchemaHandler.MapIndex[fieldName]
		t := self.pqReader.SchemaHandler.SchemaElements[schemaIndex].Type
		ct := self.pqReader.SchemaHandler.SchemaElements[schemaIndex].ConvertedType

		self.ReadColumnIndexes = append(self.ReadColumnIndexes, index)
		self.ReadColumnTypes = append(self.ReadColumnTypes, t)
		self.ReadColumnConvertedTypes = append(self.ReadColumnConvertedTypes, ct)
	}
}

//indexes should not change during read process
func (self *ParquetFileReader) Read(indexes []int) ([]*Row.Row, error) {
	if self.Cursor >= self.NumRows {
		return nil, io.EOF
	}
	if (indexes == nil) && len(self.ReadColumnIndexes) <= 0 {
		indexes = make([]int, len(self.pqReader.SchemaHandler.ValueColumns))
		for i := 0; i < len(indexes); i++ {
			indexes[i] = i
		}
		self.SetReadColumns(indexes)
	}

	if len(self.ReadColumnIndexes) <= 0 {
		self.SetReadColumns(indexes)
	}

	//log.Println("=====", indexes, self.pqReader.ColumnBuffers, self.pqReader.SchemaHandler.ValueColumns)
	var err error
	colNum := len(self.ReadColumnIndexes)
	rows := make([]*Row.Row, READ_ROWS_NUMBER)
	for i := 0; i < len(rows); i++ {
		rows[i] = Row.NewRow(make([]interface{}, colNum)...)
	}
	readRowsNumber := 0

	jobs, done := make(chan Pair), make(chan bool)
	for i := 0; i < int(Config.Conf.Runtime.ParallelNumber); i++ {
		go func() {
			defer func() {
				done <- true
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
					readRowsNumber = len(values)

					gt, _ := self.Metadata.GetTypeByIndex(index)

					for j := 0; j < len(values); j++ {
						rows[j].Vals[i] = ParquetTypeToGueryType(values[j],
							self.ReadColumnTypes[i],
							self.ReadColumnConvertedTypes[i],
							gt)
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
	for i := 0; i < int(Config.Conf.Runtime.ParallelNumber); i++ {
		<-done
	}

	self.Cursor += readRowsNumber
	return rows, nil
}
