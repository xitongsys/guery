package Connector

import (
	"io"
	"io/ioutil"
	"strings"

	"github.com/xitongsys/guery/Config"
	"github.com/xitongsys/guery/Connector/FileReader"
	"github.com/xitongsys/guery/FileSystem"
	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/Util"
)

type FileConnector struct {
	Metadata     *Util.Metadata
	FilePathList []string
	FileReader   FileReader.FileReader
	FileIndex    int
	FileType     string
}

func NewFileConnector(schema, table string) (*FileConnector, error) {
	var err error
	res := &FileConnector{}
	catalog, schema, table := "FILE", strings.ToUpper(schema), strings.ToUpper(table)
	key := strings.Join([]string{catalog, schema, table}, ".")
	conf := Config.Conf.FileConnectorConfigs[key]
	res.FileType = conf.FileType
	res.FilePathList = conf.FilePathList
	if err != nil {
		return res, err
	}

	var data []byte
	if data, err = ioutil.ReadFile(conf.FileMD); err != nil {
		Logger.Errorf("Fail to load the configure file, due to %v ", err)
		return res, err
	}
	res.Metadata, err = Util.NewMetadataFromJson(data)
	return res, err
}

func (self *FileConnector) GetMetadata() *Util.Metadata {
	return self.Metadata
}

func (self *FileConnector) Read() (*Util.Row, error) {
	if self.FileReader == nil && self.FileIndex < len(self.FilePathList) {
		vf, err := FileSystem.Open(self.FilePathList[self.FileIndex])
		if err != nil {
			return nil, err
		}
		self.FileReader, err = FileReader.NewReader(vf, self.FileType, self.Metadata)
		if err != nil {
			return nil, err
		}
		self.FileIndex++

	} else if self.FileReader == nil && self.FileIndex >= len(self.FilePathList) {
		return nil, io.EOF

	}

	row, err := self.FileReader.Read()
	if err == io.EOF {
		self.FileReader = nil
		return self.Read()
	}
	if err != nil {
		return nil, err
	}
	return row, err
}

func (self *FileConnector) ReadByColumns(colIndexes []int) (*Util.Row, error) {
	if self.FileReader == nil && self.FileIndex < len(self.FilePathList) {
		vf, err := FileSystem.Open(self.FilePathList[self.FileIndex])
		if err != nil {
			return nil, err
		}
		self.FileReader, err = FileReader.NewReader(vf, self.FileType, self.Metadata)
		if err != nil {
			return nil, err
		}
		self.FileIndex++

	} else if self.FileReader == nil && self.FileIndex >= len(self.FilePathList) {
		return nil, io.EOF

	}

	row, err := self.FileReader.ReadByColumns(colIndexes)
	if err == io.EOF {
		self.FileReader = nil
		return self.ReadByColumns(colIndexes)
	}
	if err != nil {
		return nil, err
	}
	return row, err
}
