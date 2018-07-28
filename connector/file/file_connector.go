package file

import (
	"fmt"
	"io"
	"strings"

	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/filereader"
	"github.com/xitongsys/guery/filesystem"
	"github.com/xitongsys/guery/filesystem/partition"
	"github.com/xitongsys/guery/gtype"
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/row"
)

type FileConnector struct {
	Config        *config.FileConnectorConfig
	Metadata      *metadata.Metadata
	FileReader    fileReader.FileReader
	FileType      filesystem.FileType
	PartitionInfo *partition.PartitionInfo
}

func NewFileConnectorEmpty() (*FileConnector, error) {
	return &FileConnector{}, nil
}

func NewFileConnector(catalog, schema, table string) (*FileConnector, error) {
	var err error
	res := &FileConnector{}
	key := strings.Join([]string{catalog, schema, table}, ".")
	conf := config.Conf.FileConnectorConfigs.GetConfig(key)
	if conf == nil {
		return nil, fmt.Errorf("FileConnector: table not found")
	}
	res.Config = conf
	res.FileType = filesystem.StringToFileType(conf.FileType)
	res.Metadata, err = NewFileMetadata(conf)

	return res, err
}

func NewFileMetadata(conf *config.FileConnectorConfig) (*metadata.Metadata, error) {
	res := metadata.NewMetadata()
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

func (self *FileConnector) GetMetadata() (*metadata.Metadata, error) {
	return self.Metadata, nil
}

func (self *FileConnector) GetPartitionInfo() (*partition.PartitionInfo, error) {
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
		fs, err := filesystem.List(loc)
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

func (self *FileConnector) GetReader(file *filesystem.FileLocation, md *metadata.Metadata) func(indexes []int) (*Row.RowsGroup, error) {
	reader, err := filereader.NewReader(file, md)
	return func(indexes []int) (*Row.RowsGroup, error) {
		if err != nil {
			return nil, err
		}
		return reader.Read(indexes)
	}
}

func (self *FileConnector) ShowTables(catalog, schema, table string, like, escape *string) func() (*Row.Row, error) {
	var err error
	rows := []*row.Row{}
	for key, _ := range config.Conf.FileConnectorConfigs {
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

func (self *FileConnector) ShowSchemas(catalog, schema, table string, like, escape *string) func() (*row.Row, error) {
	var err error
	rows := []*Row.Row{}
	for key, _ := range config.Conf.FileConnectorConfigs {
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

func (self *FileConnector) ShowColumns(catalog, schema, table string) func() (*row.Row, error) {
	var err error
	rows := []*Row.Row{}
	config := config.Conf.FileConnectorConfigs.GetConfig(fmt.Sprintf("%s.%s.%s.", catalog, schema, table))
	if config != nil {
		if len(config.ColumnNames) != len(config.ColumnTypes) {
			err = fmt.Errorf("%s.%s.%s: column names doesn't match column types")

		} else {
			for i, name := range config.ColumnNames {
				tname := config.ColumnTypes[i]
				row := row.NewRow()
				row.AppendVals(name, tname)
			}
		}
	} else {
		err = fmt.Errorf("%s.%s.%s: table not found", catalog, schema, table)
	}
	i := 0

	return func() (*row.Row, error) {
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

func (self *FileConnector) ShowPartitions(catalog, schema, table string) func() (*row.Row, error) {
	return func() (*row.Row, error) {
		return nil, io.EOF
	}
}
