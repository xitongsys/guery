package FileConnector

import (
	"fmt"
	"io"
	"strings"

	"github.com/xitongsys/guery/Config"
	"github.com/xitongsys/guery/FileReader"
	"github.com/xitongsys/guery/FileSystem"
	"github.com/xitongsys/guery/FileSystem/Partition"
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Row"
	"github.com/xitongsys/guery/Type"
)

type FileConnector struct {
	Config        *Config.FileConnectorConfig
	Metadata      *Metadata.Metadata
	FileReader    FileReader.FileReader
	FileType      FileSystem.FileType
	PartitionInfo *Partition.PartitionInfo
}

func NewFileConnectorEmpty() (*FileConnector, error) {
	return &FileConnector{}, nil
}

func NewFileConnector(catalog, schema, table string) (*FileConnector, error) {
	var err error
	res := &FileConnector{}
	key := strings.Join([]string{catalog, schema, table}, ".")
	conf := Config.Conf.FileConnectorConfigs.GetConfig(key)
	if conf == nil {
		return nil, fmt.Errorf("FileConnector: table not found")
	}
	res.Config = conf
	res.FileType = FileSystem.StringToFileType(conf.FileType)
	res.Metadata, err = NewFileMetadata(conf)

	return res, err
}

func NewFileMetadata(conf *Config.FileConnectorConfig) (*Metadata.Metadata, error) {
	res := Metadata.NewMetadata()
	if len(conf.ColumnNames) != len(conf.ColumnTypes) {
		return res, fmt.Errorf("File Config error: ColumnNames and ColumnTypes not match")
	}

	for i := 0; i < len(conf.ColumnNames); i++ {
		col := &Metadata.ColumnMetadata{
			Catalog:    conf.Catalog,
			Schema:     conf.Schema,
			Table:      conf.Table,
			ColumnName: conf.ColumnNames[i],
			ColumnType: Type.TypeNameToType(conf.ColumnTypes[i]),
		}
		res.AppendColumn(col)
	}

	res.Reset()
	return res, nil
}

func (self *FileConnector) GetMetadata() (*Metadata.Metadata, error) {
	return self.Metadata, nil
}

func (self *FileConnector) GetPartitionInfo() (*Partition.PartitionInfo, error) {
	if self.PartitionInfo == nil {
		if err := self.setPartitionInfo(); err != nil {
			return nil, err
		}
	}
	return self.PartitionInfo, nil
}

func (self *FileConnector) setPartitionInfo() error {
	parMD := Metadata.NewMetadata()
	self.PartitionInfo = Partition.NewPartitionInfo(parMD)
	for _, loc := range self.Config.PathList {
		fs, err := FileSystem.List(loc)
		if err != nil {
			return err
		}
		for _, f := range fs {
			f.FileType = self.FileType
		}
		self.PartitionInfo.FileList = append(self.PartitionInfo.FileList, fs...)
	}
	return nil
}

func (self *FileConnector) GetReader(file *FileSystem.FileLocation, md *Metadata.Metadata) func(indexes []int) (*Row.RowsGroup, error) {
	reader, err := FileReader.NewReader(file, md)
	return func(indexes []int) (*Row.RowsGroup, error) {
		if err != nil {
			return nil, err
		}
		return reader.Read(indexes)
	}
}

func (self *FileConnector) ShowTables(catalog, schema, table string, like, escape *string) func() (*Row.Row, error) {
	var err error
	rows := []*Row.Row{}
	for key, _ := range Config.Conf.FileConnectorConfigs {
		ns := strings.Split(key, ".")
		if len(ns) < 3 {
			err = fmt.Errorf("Config name error: key")
			break
		}
		c, s, t := ns[0], ns[1], ns[2]
		if c == catalog && s == schema {
			row := Row.NewRow()
			row.AppendVals(t)
			rows = append(rows, row)
		}
	}
	i := 0

	return func() (*Row.Row, error) {
		if err != nil {
			return nil, err
		}
		if i >= len(rows) {
			return nil, io.EOF
		}
		i++
		return rows[i-1], nil
	}
}

func (self *FileConnector) ShowSchemas(catalog, schema, table string, like, escape *string) func() (*Row.Row, error) {
	var err error
	rows := []*Row.Row{}
	for key, _ := range Config.Conf.FileConnectorConfigs {
		ns := strings.Split(key, ".")
		if len(ns) < 3 {
			err = fmt.Errorf("Config name error: key")
			break
		}
		c, s, _ := ns[0], ns[1], ns[2]
		if c == catalog {
			row := Row.NewRow()
			row.AppendVals(s)
			rows = append(rows, row)
		}
	}
	i := 0

	return func() (*Row.Row, error) {
		if err != nil {
			return nil, err
		}
		if i >= len(rows) {
			return nil, io.EOF
		}
		i++
		return rows[i-1], nil
	}
}

func (self *FileConnector) ShowColumns(catalog, schema, table string) func() (*Row.Row, error) {
	var err error
	rows := []*Row.Row{}
	config := Config.Conf.FileConnectorConfigs.GetConfig(fmt.Sprintf("%s.%s.%s.", catalog, schema, table))
	if config != nil {
		if len(config.ColumnNames) != len(config.ColumnTypes) {
			err = fmt.Errorf("%s.%s.%s: column names doesn't match column types")

		} else {
			for i, name := range config.ColumnNames {
				tname := config.ColumnTypes[i]
				row := Row.NewRow()
				row.AppendVals(name, tname)
			}
		}
	} else {
		err = fmt.Errorf("%s.%s.%s: table not found", catalog, schema, table)
	}
	i := 0

	return func() (*Row.Row, error) {
		if err != nil {
			return nil, err
		}
		if i >= len(rows) {
			return nil, io.EOF
		}

		i++
		return rows[i-1], nil
	}
}

func (self *FileConnector) ShowPartitions(catalog, schema, table string) func() (*Row.Row, error) {
	return func() (*Row.Row, error) {
		return nil, io.EOF
	}
}
