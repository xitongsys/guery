package filereader

import (
	"fmt"

	"github.com/xitongsys/guery/filereader/csv"
	"github.com/xitongsys/guery/filereader/orc"
	"github.com/xitongsys/guery/filereader/parquet"
	"github.com/xitongsys/guery/filesystem"
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/row"
)

type FileReader interface {
	Read(indexes []int) (rg *row.RowsGroup, err error)
}

func NewReader(file *filesystem.FileLocation, md *metadata.Metadata) (FileReader, error) {

	switch file.FileType {
	case filesystem.CSV:
		vf, err := filesystem.Open(file.Location)
		if err != nil {
			return nil, err
		}
		return csv.New(vf, md), nil

	case filesystem.PARQUET:
		return parquet.New(file.Location, md), nil

	case filesystem.ORC:
		vf, err := filesystem.Open(file.Location)
		if err != nil {
			return nil, err
		}
		return orc.New(vf, md)
	}
	return nil, fmt.Errorf("File type %s is not defined.", file.FileType)
}
