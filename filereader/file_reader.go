package FileReader

import (
	"fmt"

	"github.com/xitongsys/guery/FileReader/Csv"
	"github.com/xitongsys/guery/FileReader/Orc"
	"github.com/xitongsys/guery/FileReader/Parquet"
	"github.com/xitongsys/guery/FileSystem"
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Row"
)

type FileReader interface {
	Read(indexes []int) (rg *Row.RowsGroup, err error)
}

//func NewReader(vf FileSystem.VirtualFile, fileType string, md *Util.Metadata) (FileReader, error) {
func NewReader(file *FileSystem.FileLocation, md *Metadata.Metadata) (FileReader, error) {

	switch file.FileType {
	case FileSystem.CSV:
		vf, err := FileSystem.Open(file.Location)
		if err != nil {
			return nil, err
		}
		return Csv.New(vf, md), nil

	case FileSystem.PARQUET:
		return Parquet.New(file.Location, md), nil

	case FileSystem.ORC:
		vf, err := FileSystem.Open(file.Location)
		if err != nil {
			return nil, err
		}
		return Orc.New(vf, md)
	}
	return nil, fmt.Errorf("File type %s is not defined.", file.FileType)
}
