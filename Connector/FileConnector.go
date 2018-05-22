package Connector

import (
	"fmt"
	"io"
	"ioutil"
	"path/filepath"
	"strings"

	"github.com/xitongsys/guery/Config"
	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/Util"
)

type FileConnector struct {
	Metadata  *Util.Metadata
	FileList  []string
	FileIndex int
	FileType  string
}

func NewFileConnector(schema, table string) (*TestConnector, error) {
	var err error
	res := &FileConnector{}
	conf := Config.Conf["FILE"]["METASTORE"][schema][table]
	res.FileType = conf["TYPE"]
	res.FileList, err = GetFiles(conf["DATA"])
	if err != nil {
		return res, err
	}

	var data []byte
	if data, err = ioutil.ReadFile(conf["SCHEMA"]); err != nil {
		Logger.Errorf("Fail to load the configure file, due to %v ", err)
		return res, err
	}
	res.Metadata = Util.NewMetadataFromJson(data)
	return res
}

func (self *FileConnector) GetMetadata() *Util.Metadata {
	return self.Metadata
}

func (self *FileConnector) ReadRow() (*Util.Row, error) {

}

func (self *FileConnector) SkipRows(num int64) {
	self.Index += num
}

func (self *FileConnector) ReadRowByColumns(colIndexes []int) (*Util.Row, error) {
}
