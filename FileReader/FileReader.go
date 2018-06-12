package FileReader

import (
	"fmt"

	"github.com/xitongsys/guery/FileReader/Csv"
	"github.com/xitongsys/guery/FileReader/Parquet"
	"github.com/xitongsys/guery/FileSystem"
	"github.com/xitongsys/guery/Util"
)

type FileReader interface {
	Read(indexes []int) (row *Util.Row, err error)
}

//func NewReader(vf FileSystem.VirtualFile, fileType string, md *Util.Metadata) (FileReader, error) {
func NewReader(file *FileSystem.FileLocation, md *Util.Metadata) (FileReader, error) {

	switch file.FileType {
	case FileSystem.CSV:
		vf, err := FileSystem.Open(file.Location)
		if err != nil {
			return nil, err
		}
		return Csv.New(vf, md), nil

	case FileSystem.PARQUET:
		return Parquet.New(file.Location), nil
	}
	return nil, fmt.Errorf("File type %s is not defined.", file.FileType)
}
