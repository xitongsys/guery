package Parquet

import (
	"io"

	"github.com/xitongsys/guery/Config"
	"github.com/xitongsys/guery/FileSystem"
	"github.com/xitongsys/guery/Split"
	. "github.com/xitongsys/parquet-go/ParquetFile"
	. "github.com/xitongsys/parquet-go/ParquetReader"
	"github.com/xitongsys/parquet-go/guery/Metadata"
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
	Metadata *Metadata.Metadata
	pqReader *ParquetReader

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
func (self *ParquetFileReader) Read(indexes []int) (row *Split.Split, err error) {
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

	sq := Split.NewSplit(self.Metadata)
	for i, index := range self.ReadColumnIndexes {
		values, _, _ := self.pqReader.ReadColumnByIndex(index, Split.MAX_SPLIT_SIZE)
		if len(values) <= 0 {
			return nil, io.EOF
		}

		if sq.RowsNumber <= 0 {
			sq.RowsNumber = len(values)
		} else if sq.RowsNumber != len(values) {
			return sq, fmt.Errorf("ParquetFileReader Read error")
		}

		for _, v := range values {
			gv := ParquetTypeToGueryType(v,
				self.ReadColumnTypes[i],
				self.ReadColumnConvertedTypes[i])
			sq.Values[i] = append(sq.Values[i], gv)
			if gv == nil {
				sq.ValueFlags[i] = append(sq.ValueFlags[i], false)
			} else {
				sq.ValueFlags[i] = append(sq.ValueFlags[i], true)
			}
		}
	}
	return sq, nil
}
