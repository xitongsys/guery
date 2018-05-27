package HiveConnector

import (
	"fmt"

	"github.com/xitongsys/guery/Connector/FileReader"
	"github.com/xitongsys/guery/Util"
)

func (self *HiveConnector) setMetadata() (err error) {
	if err = self.getConn(); err != nil {
		return err
	}
	sqlStr := fmt.Sprintf(MD_SQL, self.Schema, self.Table, self.Schema, self.Table)
	rows, err := self.db.Query(sqlStr)
	if err != nil {
		return err
	}
	var colName, colType string
	names, types := []string{}, []Util.Type{}
	for rows.Next() {
		rows.Scan(&colName, colType)
		names = append(names, colName)
		types = append(types, HiveTypeToGueryType(colType))
	}
	self.Metadata = Util.NewMetadata()
	for i, name := range names {
		t := types[i]
		column := Util.NewColumnMetadata(t, "HIVE", self.Schema, self.Table, name)
		self.Metadata.AppendColumn(column)
	}
	self.Metadata.Reset()
	return nil
}

func (self *HiveConnector) setPartitionInfo() (err error) {
	if err = self.getConn(); err != nil {
		return err
	}
	sqlStr := fmt.Sprintf(PARTITION_MD_SQL, self.Schema, self.Table)
	rows, err := self.db.Query(sqlStr)
	if err != nil {
		return err
	}
	var colName, colType string
	names, types := []string{}, []Util.Type{}
	for rows.Next() {
		rows.Scan(&colName, colType)
		names = append(names, colName)
		types = append(types, HiveTypeToGueryType(colType))
	}

	md := Util.NewMetadata()
	for i, name := range names {
		t := types[i]
		column := Util.NewColumnMetadata(t, "HIVE", self.Schema, self.Table, name)
		md.AppendColumn(column)
	}
	md.Reset()
	self.PartitionInfo = Util.NewPartitionInfo(md)

	sqlStr = fmt.Sprintf(PARTITION_DATA_SQL, self.Schema, self.Table)
	rows, err = self.db.Query(sqlStr)
	if err != nil {
		return err
	}

	pnum := md.GetColumnNumber()
	partitions := make([]string, pnum)
	self.PartitionReaders = make([]FileReader.FileReader, pnum)

	location, fileType := "", ""
	i := 0
	for rows.Next() {
		rows.Scan(&location, &fileType, &partitions[i])
		if i == pnum-1 {
			row := Util.NewRow()
			for j := 0; j < pnum; j++ {
				t, err := md.GetTypeByIndex(j)
				if err != nil {
					return err
				}
				row.AppendKeys(Util.ToType(partitions[i], t))
			}
			self.PartitionInfo.Write(row)
			self.PartitionInfo.Locations = append(self.PartitionInfo.Locations, location)
			self.PartitionInfo.FileTypes = append(self.PartitionInfo.FileTypes, HiveFileTypeToSimpleType(fileType))

			i = 0
		} else {
			i++
		}
	}
	return nil
}

func (self *HiveConnector) getConn() error {
	if self.db != nil {
		if err := self.db.Ping(); err != nil {
			self.db.Close()
			self.db = nil
		} else {
			return nil
		}
	}

	db, err := Util.OpenDBConn("mysql", self.Config.GetURI())
	if err != nil {
		self.db = nil
		return err
	}
	self.db = db
	return nil
}
