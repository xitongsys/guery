package FileReader

import (
	"fmt"
	"strings"

	"github.com/xitongsys/guery/Connector/FileReader/Csv"
	"github.com/xitongsys/guery/FileSystem"
	"github.com/xitongsys/guery/Util"
)

type FileReader interface {
	Read() (row *Util.Row, err error)
	ReadByColumns(indexes []int) (row *Util.Row, err error)
}

func NewReader(vf FileSystem.VirtualFile, fileType string, md *Util.Metadata) (FileReader, error) {
	switch strings.ToUpper(fileType) {
	case "CSV":
		return Csv.New(vf, md), nil
	}
	return nil, fmt.Errorf("File type %s is not defined.", fileType)
}
