package Parquet

import (
	"io"
	"log"

	"github.com/xitongsys/guery/FileSystem"
	"github.com/xitongsys/guery/Row"
	. "github.com/xitongsys/parquet-go/ParquetFile"
	. "github.com/xitongsys/parquet-go/ParquetReader"
	"github.com/xitongsys/parquet-go/parquet"
)

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
func (self *PqFile) Read(b []byte) (n int, err error) {
	return self.VF.Read(b)
}
func (self *PqFile) Write(b []byte) (n int, err error) {
	return 0, nil
}
func (self *PqFile) Close() error { return nil }

type ParquetFileReader struct {
	pqReader *ParquetReader
	NumRows  int
	Cursor   int

	ReadColumnIndexes        []int
	ReadColumnTypes          []*parquet.Type
	ReadColumnConvertedTypes []*parquet.ConvertedType
}

func New(fileName string) *ParquetFileReader {
	parquetFileReader := new(ParquetFileReader)
	var pqFile ParquetFile = &PqFile{}
	pqFile, _ = pqFile.Open(fileName)
	parquetFileReader.pqReader, _ = NewParquetColumnReader(pqFile, 1)
	parquetFileReader.NumRows = int(parquetFileReader.pqReader.GetNumRows())
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
	log.Println("======", self.ReadColumnIndexes, self.ReadColumnTypes, self.ReadColumnConvertedTypes)
}

//indexes should not change during read process
func (self *ParquetFileReader) Read(indexes []int) (row *Row.Row, err error) {
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

	objects := make([]interface{}, 0)
	for i, index := range self.ReadColumnIndexes {
		values, _, _ := self.pqReader.ReadColumnByIndex(index, 1)
		if len(values) <= 0 {
			return nil, io.EOF
		}
		objects = append(objects, ParquetTypeToGueryType(values[0],
			self.ReadColumnTypes[i],
			self.ReadColumnConvertedTypes[i],
		))
	}
	self.Cursor++
	row = &Row.Row{}
	row.AppendVals(objects...)
	return row, nil
}
