package Connector

import (
	"fmt"
	"io"
	"ioutil"
	"path/filepath"
	"strings"

	"github.com/xitongsys/guery/Config"
	"github.com/xitongsys/guery/Util"
)

type LocationFileConnector struct {
	Metadata *Util.Metadata
	FileList []string
	FileType string
}

func GetFiles(filePattern string) ([]string, error) {
	return filepath.Glob(filePattern)
}

func NewLocalFileConnector(schema, table string) (*TestConnector, error) {
	var err error
	res := &LocationFileConnector{}
	conf := Config.Conf["LOCALFILE"][schema][table]
	res.FileType = conf["TYPE"]
	res.FileList, err = GetFiles(conf["DATA"])
	if err != nil {
		return res, err
	}

	return res
}

func (self *TestConnector) GetMetadata() *Util.Metadata {
	return self.Metadata
}

func (self *TestConnector) ReadRow() (*Util.Row, error) {
	if self.Index >= int64(len(self.Rows)) {
		self.Index = 0
		return nil, io.EOF
	}

	self.Index++
	return &self.Rows[self.Index-1], nil
}

func (self *TestConnector) SkipTo(index, total int64) {
	ln := int64(len(self.Rows))
	pn := ln / total
	left := ln % total
	if left > index {
		left = index
	}
	self.Index = pn*index + left
}

func (self *TestConnector) SkipRows(num int64) {
	self.Index += num
}

func (self *TestConnector) ReadRowByColumns(colIndexes []int) (*Util.Row, error) {
	if self.Index >= int64(len(self.Rows)) {
		self.Index = 0
		return nil, io.EOF
	}
	self.Index++
	row := &Util.Row{}
	for _, ci := range colIndexes {
		row.AppendVals(self.Rows[self.Index-1].Vals[ci])
	}
	return row, nil
}
