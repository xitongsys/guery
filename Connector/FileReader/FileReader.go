package FileReader

import (
	"strings"

	"github.com/xitongsys/guery/Connector/FileReader/Csv"
	"github.com/xitongsys/guery/FileSystem"
	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/Util"
)

type FileReader interface {
	Read() (row *Util.Row, err error)
	ReadHeader() (fieldNames []string, err error)
}

func NewReader(vf FileSystem.VirtualFile, fileType string, md *Util.Metadata) (FileReader, error) {
	switch strings.ToUpper(fileType) {
	case "CSV":
		return Csv.New(vf, md), nil
	}
	return nil, Logger.Errorf("File type %s is not defined.", fileType)
}
